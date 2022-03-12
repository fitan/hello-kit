package service

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/metrics/prometheus"
	http3 "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	endpoint "hello/pkg/endpoint"
	http1 "hello/pkg/http"
	service "hello/pkg/service"
	"net"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"

	http4 "github.com/fitan/gink/transport/http"
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/log"
	group "github.com/oklog/oklog/pkg/group"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var logger log.Logger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("hello", flag.ExitOnError)
var debugAddr = fs.String("debug-addr", ":8081", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8080", "HTTP listen address")
var grpcAddr = fs.String("grpc-addr", ":8082", "gRPC listen address")
var thriftAddr = fs.String("thrift-addr", ":8083", "Thrift listen address")
var thriftProtocol = fs.String("thrift-protocol", "binary", "binary, compact, json, simplejson")
var thriftBuffer = fs.Int("thrift-buffer", 0, "0 for unbuffered")
var thriftFramed = fs.Bool("thrift-framed", false, "true to enable framing")
var jaegerAddr = fs.String("jaeger-addr", "http://localhost:14268/api/traces", "Enable jaeger tracing via an jaeger-addr server http://localhost:14268/api/traces")

const (
	appName     = "trace-demo"
	environment = "production"
	id          = 1
)

func Run() {
	fs.Parse(os.Args[1:])

	// Create a single logger, which we'll use and give to other components.
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stderr))
	//logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	logger = log.With(logger, "caller", log.DefaultCaller)

	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	if *jaegerAddr != "" {
		logger.Log("init tracer", "jaeger", "addr", *jaegerAddr)
		tp := initTracer()
		otel.SetTracerProvider(tp)
	}

	svc := service.New(getServiceMiddleware(logger))
	eps := endpoint.New(svc, getEndpointMiddleware(logger))
	g := createService(eps)
	initGinHandler(eps, g)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Log("exit", g.Run())

}

func initGinHandler(endpoints endpoint.Endpoints, g *group.Group) {
	ginHandler := http1.NewGinHandler(endpoints, map[string][]http4.ServerOption{})
	httpListener, err := net.Listen("tcp", ":8090")
	if err != nil {
		logger.Log("transport", "GinHTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "GinHTTP", "addr", "8090")
		return http2.Serve(httpListener, ginHandler)
	}, func(error) {
		httpListener.Close()
	})
}
func initHttpHandler(endpoints endpoint.Endpoints, g *group.Group) {
	options := defaultHttpOptions(logger)

	pc := http3.ServerBefore(http3.PopulateRequestContext)
	for k, _ := range options {
		options[k] = append(options[k], pc)
	}

	httpHandler := http1.NewHTTPHandler(endpoints, options)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Log("transport", "HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		return http2.Serve(httpListener, httpHandler)
	}, func(error) {
		httpListener.Close()
	})

}

func initTracer() *sdktrace.TracerProvider {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String("test-service"),
		),
	)
	if err != nil {
		logger.Log("failed to create resource", err)
		panic(err)
	}
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(*jaegerAddr)))
	if err != nil {
		logger.Log("jaeger.New err", err)
		panic(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	tp.Tracer(appName)
	return tp
}

func getServiceMiddleware(logger log.Logger) (mw []service.Middleware) {
	mw = []service.Middleware{}
	// Append your middleware here
	mw = append(mw, func(helloService service.HelloService) service.HelloService {
		return service.NewHelloServiceWithPrometheus(helloService, "hello-kit")
	})
	mw = append(mw, func(helloService service.HelloService) service.HelloService {
		return service.NewHelloServiceWithLog(helloService, logger)
	})
	mw = append(mw, func(helloService service.HelloService) service.HelloService {
		return service.NewHelloServiceWithTracing(helloService)
	})

	return
}
func getEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// Add you endpoint middleware here
	addEndpointMiddlewareToAllMethods(mw, otelkit.EndpointMiddleware())
	addEndpointMiddlewareToAllMethods(mw, endpoint.LoggingMiddleware(logger))
	addEndpointMiddlewareToAllMethods(mw, endpoint.InstrumentingMiddleware(prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "hello",
		Subsystem: "api",
		Name:      "request_duration_seconds",
		Help:      "Total time spent serving requests.",
	}, []string{"success"})))

	return
}
func initMetricsEndpoint(g *group.Group) {
	http2.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Log("transport", "debug/HTTP", "during", "Listen", "err", err)
	}
	g.Add(func() error {
		logger.Log("transport", "debug/HTTP", "addr", *debugAddr)
		return http2.Serve(debugListener, http2.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
}
func initCancelInterrupt(g *group.Group) {
	cancelInterrupt := make(chan struct{})
	g.Add(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case sig := <-c:
			return fmt.Errorf("received signal %s", sig)
		case <-cancelInterrupt:
			return nil
		}
	}, func(error) {
		close(cancelInterrupt)
	})
}

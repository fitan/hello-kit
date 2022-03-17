package service

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hello/pkg/services"
	log "hello/utils/log"
	"net"
	http2 "net/http"
	"os"
	"os/signal"
	"syscall"

	group "github.com/oklog/oklog/pkg/group"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

var logger *zap.SugaredLogger

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("hello", flag.ExitOnError)
var debugAddr = fs.String("debug-addr", ":8081", "Debug and metrics listen address")
var httpAddr = fs.String("http-addr", ":8080", "HTTP listen address")
var jaegerAddr = fs.String("jaeger-addr", "http://localhost:14268/api/traces", "Enable jaeger tracing via an jaeger-addr server http://localhost:14268/api/traces")

const (
	appName     = "trace-demo"
	environment = "production"
	id          = 1
)

func Run() {
	fs.Parse(os.Args[1:])

	core := log.DefaultZapCore("hello-kit", "./logs", zapcore.InfoLevel)
	l := zap.New(core, zap.AddCaller())
	logger = l.Sugar()

	//  Determine which tracer to use. We'll pass the tracer to all the
	// components that use it, as a dependency
	if *jaegerAddr != "" {
		logger.Infow("init jaeger", "addr", *jaegerAddr)
		tp := initTracer()
		otel.SetTracerProvider(tp)
	}

	g := &group.Group{}
	initHttpHandler(g, logger)
	initMetricsEndpoint(g)
	initCancelInterrupt(g)
	logger.Infow("exit", g.Run())

}

func initHttpHandler(g *group.Group, log *zap.SugaredLogger) {
	r := gin.Default()

	services.InitHttp(r, log, appName, nil, nil)
	httpListener, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		logger.Errorw(err.Error(), "transport", "HTTP", "during", "Listen")
	}
	g.Add(func() error {
		logger.Infow("http listener", "transport", "HTTP", "addr", *httpAddr)
		return http2.Serve(httpListener, r)
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
		logger.Errorw("resource.New", "err", err.Error())
		panic(err)
	}
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(*jaegerAddr)))
	if err != nil {
		logger.Errorw("jaeger.New", "err", err.Error())
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

func initMetricsEndpoint(g *group.Group) {
	http2.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", *debugAddr)
	if err != nil {
		logger.Errorw("net.Listen", "transport", "debug/HTTP", "during", "Listen", "err", err.Error())
	}
	g.Add(func() error {
		logger.Infow("init metrics", "transport", "debug/HTTP", "addr", *debugAddr)
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

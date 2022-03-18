package app

import (
	"context"
	"flag"
	"fmt"
	ginkHttp "github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/oklog/oklog/pkg/group"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"hello/pkg/repository"
	"hello/pkg/services"
	"hello/utils/conf"
	"hello/utils/log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Define our flags. Your service probably won't need to bind listeners for
// all* supported transports, but we do it here for demonstration purposes.
var fs = flag.NewFlagSet("hello", flag.ExitOnError)
var confName = fs.String("conf", "dev.conf", "open config")
var logger *zap.SugaredLogger

type App struct {
	r          *gin.Engine
	repository *repository.Repository
	services   *services.Services
	g          *group.Group
	conf       *conf.MyConf
	log        *zap.SugaredLogger
	tp         *sdktrace.TracerProvider
	InitCancelInterrupt
	InitMetricsEndpoint
	InitHttpHandler
}

func (a *App) Run() error {
	return a.g.Run()
}

func RunApp() {
	fs.Parse(os.Args[1:])

	r := gin.Default()
	g := &group.Group{}
	app := InitApp(r, g, ConfName(*confName))
	logger.Errorw("exit", app.Run().Error())
}

type ConfName string

func initLog(conf *conf.MyConf) *zap.SugaredLogger {
	core := log.DefaultZapCore(conf.Log.FileName, conf.Log.Dir, zapcore.Level(conf.Log.Lervel))
	logger = zap.New(core, zap.AddCaller()).Sugar()
	return logger
}

func initConf(confName ConfName) *conf.MyConf {
	myConf := conf.MyConf{}
	_, err := conf.WatchFile(string(confName), []string{"./conf"}, &myConf, 5*time.Second)
	if err != nil {
		panic("conf.WatchFile" + err.Error())
	}
	return &myConf
}

type InitHttpHandler struct {
}

func initHttpHandler(g *group.Group, log *zap.SugaredLogger, conf *conf.MyConf) InitHttpHandler {
	r := gin.Default()

	httpListener, err := net.Listen("tcp", conf.App.Addr)
	if err != nil {
		log.Errorw(err.Error(), "transport", "HTTP", "during", "Listen")
	}
	g.Add(func() error {
		log.Infow("http listener", "transport", "HTTP", "addr", conf.App.Name)
		return http.Serve(httpListener, r)
	}, func(error) {
		httpListener.Close()
	})

	return InitHttpHandler{}

}

func initTracer(conf *conf.MyConf) *sdktrace.TracerProvider {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(conf.App.Name),
		),
	)
	if err != nil {
		logger.Fatalw("resource.New", "err", err.Error())
	}
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(conf.Trace.TracerProviderAddr)))
	if err != nil {
		logger.Fatalw("jaeger.New", "err", err.Error())
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tp)
	return tp
}

type InitMetricsEndpoint struct {
}

func initMetricsEndpoint(g *group.Group, conf *conf.MyConf) InitMetricsEndpoint {
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())
	debugListener, err := net.Listen("tcp", conf.Debug.Addr)
	if err != nil {
		logger.Errorw("net.Listen", "transport", "debug/HTTP", "during", "Listen", "err", err.Error())
	}
	g.Add(func() error {
		logger.Infow("init metrics", "transport", "debug/HTTP", "addr", conf.Debug.Addr)
		return http.Serve(debugListener, http.DefaultServeMux)
	}, func(error) {
		debugListener.Close()
	})
	return InitMetricsEndpoint{}
}

type InitCancelInterrupt struct {
}

func initCancelInterrupt(g *group.Group) InitCancelInterrupt {
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
	return InitCancelInterrupt{}
}

func initEndpointMiddleware() []endpoint.Middleware {
	mw := make([]endpoint.Middleware, 0)
	return mw
}

func initHttpServerOption() []ginkHttp.ServerOption {
	so := make([]ginkHttp.ServerOption, 0)
	return so
}

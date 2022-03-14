package hello

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"hello/pkg/mid"
)

func serviceMiddleware(logger log.Logger, appName string) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(helloService HelloService) HelloService {
		return NewHelloServiceWithPrometheus(helloService, appName)
	})
	mw = append(mw, func(helloService HelloService) HelloService {
		return NewHelloServiceWithLog(helloService, logger)
	})
	mw = append(mw, func(helloService HelloService) HelloService {
		return NewHelloServiceWithTracing(helloService)
	})

	return
}

func endpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// Add you endpoint middleware here
	otelkitEMW := func(n string) endpoint1.Middleware {
		return otelkit.EndpointMiddleware(otelkit.WithOperation(n))
	}
	logEMW := func(n string) endpoint1.Middleware {
		return mid.LoggingMiddleware(logger)
	}
	instrumentingEMW := func(n string) endpoint1.Middleware {
		return mid.InstrumentingMiddleware(prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "hello",
			Subsystem: n,
			Name:      "request_duration_seconds",
			Help:      "Total time spent serving requests.",
		}, []string{"success"}))
	}
	AddEndpointMiddlewareToAllMethods(mw, otelkitEMW)
	AddEndpointMiddlewareToAllMethods(mw, logEMW)
	AddEndpointMiddlewareToAllMethods(mw, instrumentingEMW)

	return
}

func InitHttpHandler(r *gin.Engine, log log.Logger, appName string) {

	helloSvc := New(serviceMiddleware(log, appName))
	helloEps := NewEndpoints(helloSvc, endpointMiddleware(log))
	pc := http.ServerBefore(http.PopulateRequestContext)
	helloOptions := make(map[string][]http.ServerOption)
	AddHttpOptionToAllMethods(helloOptions, pc)
	NewHTTPHandler(r, helloEps, helloOptions)
}

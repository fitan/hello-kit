package hello

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
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

func endpointMiddleware(logger log.Logger, ep []endpoint.Middleware) (mw map[string][]endpoint.Middleware) {
	mw = map[string][]endpoint.Middleware{}
	// Add you endpoint middleware here
	otelkitEMW := func(n string) endpoint.Middleware {
		return otelkit.EndpointMiddleware(otelkit.WithOperation(n))
	}
	logEMW := func(n string) endpoint.Middleware {
		return mid.LoggingMiddleware(logger)
	}
	instrumentingEMW := func(n string) endpoint.Middleware {
		return mid.InstrumentingMiddleware(prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "hello",
			Subsystem: n,
			Name:      "request_duration_seconds",
			Help:      "Total time spent serving requests.",
		}, []string{"success"}))
	}
	AddEndpointMiddlewareToAllMethodsWithMethodName(mw, otelkitEMW)
	AddEndpointMiddlewareToAllMethodsWithMethodName(mw, logEMW)
	AddEndpointMiddlewareToAllMethodsWithMethodName(mw, instrumentingEMW)
	AddEndpointMiddlewareToAllMethods(mw, ep)

	return
}

func serviceOption(op []http.ServerOption) map[string][]http.ServerOption {
	ops := make(map[string][]http.ServerOption)
	for _, o := range op {
		AddHttpOptionToAllMethods(ops, o)
	}
	return ops
}

func InitHttpHandler(r *gin.Engine, log log.Logger, appName string, eps []endpoint.Middleware, ops []http.ServerOption) {
	Svc := NewService(serviceMiddleware(log, appName))
	Eps := NewEndpoints(Svc, endpointMiddleware(log, eps))
	Ops := serviceOption(ops)
	NewHTTPHandler(r, Eps, Ops)
}
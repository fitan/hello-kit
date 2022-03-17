package hello

import (
	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.uber.org/zap"
	"hello/pkg/mid"
)

func NewServiceMiddleware(logger *zap.SugaredLogger, appName string) (mw []Middleware) {
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

func NewEndpointMiddleware(logger *zap.SugaredLogger, ep []endpoint.Middleware) (mw map[string][]endpoint.Middleware) {
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
	for _, e := range ep {
		AddEndpointMiddlewareToAllMethods(mw, e)
	}

	return
}

func NewServiceOption(op []http.ServerOption) map[string][]http.ServerOption {
	ops := make(map[string][]http.ServerOption)
	for _, o := range op {
		AddHttpOptionToAllMethods(ops, o)
	}
	return ops
}

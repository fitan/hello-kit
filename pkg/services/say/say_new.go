package say

import (
	"hello/pkg/mid"

	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.uber.org/zap"
)

type Middleware func(SayService) SayService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(SayService SayService) SayService {
		return NewSayServiceWithPrometheus(SayService)
	})
	mw = append(mw, func(SayService SayService) SayService {
		return NewSayServiceWithLog(SayService, logger)
	})
	mw = append(mw, func(SayService SayService) SayService {
		return NewSayServiceWithTracing(SayService)
	})
	return
}

// New returns a SayService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) SayService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

func NewEndpointMiddleware(logger *zap.SugaredLogger, ep []endpoint.Middleware) Mws {
	mws := make(Mws)
	// Add you endpoint middleware here
	otelkitEMW := func(n string) endpoint.Middleware {
		return otelkit.EndpointMiddleware(otelkit.WithOperation(n))
	}
	//logEMW := func(n string) endpoint.Middleware {
	//	return mid.LoggingMiddleware(logger)
	//}
	instrumentingEMW := func(n string) endpoint.Middleware {
		return mid.InstrumentingMiddleware(prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "pkg_services",
			Subsystem: n,
			Name:      "request_duration_seconds",
			Help:      "Total time spent serving requests.",
		}, []string{"success"}))
	}
	AddEndpointMiddlewareToAllMethodsWithMethodName(mws, otelkitEMW)
	//AddEndpointMiddlewareToAllMethodsWithMethodName(mw, logEMW)
	AddEndpointMiddlewareToAllMethodsWithMethodName(mws, instrumentingEMW)
	for _, e := range ep {
		AddEndpointMiddlewareToAllMethods(mws, e)
	}

	return mws
}

func NewServiceOption(op []http.ServerOption) Ops {
	ops := make(Ops)
	for _, o := range op {
		AddHttpOptionToAllMethods(ops, o)
	}
	return ops
}

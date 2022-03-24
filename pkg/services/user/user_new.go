package user

import (
	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.uber.org/zap"
	"hello/pkg/mid"
)

type Middleware func(UserService) UserService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(UserService UserService) UserService {
		return NewUserServiceWithPrometheus(UserService, "pkg.services.User")
	})
	mw = append(mw, func(UserService UserService) UserService {
		return NewUserServiceWithLog(UserService, logger)
	})
	mw = append(mw, func(UserService UserService) UserService {
		return NewUserServiceWithTracing(UserService)
	})
	return
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
			Namespace: "user_endpoint",
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

// New returns a UserService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) UserService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

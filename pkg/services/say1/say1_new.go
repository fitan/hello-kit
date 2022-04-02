package say1

import (
	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.uber.org/zap"
)

type Middleware func(Say1Service) Say1Service

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(Say1Service Say1Service) Say1Service {
		return NewSay1ServiceWithPrometheus(Say1Service)
	})
	mw = append(mw, func(Say1Service Say1Service) Say1Service {
		return NewSay1ServiceWithLog(Say1Service, logger)
	})
	mw = append(mw, func(Say1Service Say1Service) Say1Service {
		return NewSay1ServiceWithTracing(Say1Service)
	})
	return
}

// New returns a Say1Service with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) Say1Service {

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

	AddEndpointMiddlewareToAllMethodsWithMethodName(mws, otelkitEMW)
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

var Say1KitSet = wire.NewSet(NewBasicService, NewService, NewEndpointMiddleware, NewServiceMiddleware, NewEndpoints, NewServiceOption, NewHTTPHandler)

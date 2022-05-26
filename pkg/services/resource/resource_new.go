package resource

import (
	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.uber.org/zap"
)

type Middleware func(ResourceService) ResourceService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(ResourceService ResourceService) ResourceService {
		return NewResourceServiceWithPrometheus(ResourceService)
	})
	mw = append(mw, func(ResourceService ResourceService) ResourceService {
		return NewResourceServiceWithLog(ResourceService, logger)
	})
	mw = append(mw, func(ResourceService ResourceService) ResourceService {
		return NewResourceServiceWithTracing(ResourceService)
	})
	return
}

// New returns a ResourceService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) ResourceService {

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

var ResourceKitSet = wire.NewSet(NewBasicService, NewService, NewEndpointMiddleware, NewServiceMiddleware, NewEndpoints, NewServiceOption, NewHTTPHandler)

package hello

import (
	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"go.uber.org/zap"
)

type Middleware func(HelloService) HelloService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(HelloService HelloService) HelloService {
		return NewHelloServiceWithPrometheus(HelloService)
	})
	mw = append(mw, func(HelloService HelloService) HelloService {
		return NewHelloServiceWithLog(HelloService, logger)
	})
	mw = append(mw, func(HelloService HelloService) HelloService {
		return NewHelloServiceWithTracing(HelloService)
	})
	return
}

// New returns a HelloService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) HelloService {

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
	AddEndpointMiddlewareToAllMethodsWithMethodName(mws, otelkitEMW)
	//AddEndpointMiddlewareToAllMethodsWithMethodName(mw, logEMW)
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

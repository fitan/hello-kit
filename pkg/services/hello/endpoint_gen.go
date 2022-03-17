package hello

// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap/templates/endpoint.tmpl
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"hello/pkg/services/hello/types"

	"go.opentelemetry.io/otel/trace"

	endpoint "github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FooEndpoint endpoint.Endpoint

	SayEndpoint endpoint.Endpoint

	Say1Endpoint endpoint.Endpoint

	SayHelloEndpoint endpoint.Endpoint

	SayHello1Endpoint endpoint.Endpoint
}

func AddEndpointMiddlewareToAllMethods(mw map[string][]endpoint.Middleware, m endpoint.Middleware) {
	methods := []string{

		"Foo",

		"Say",

		"Say1",

		"SayHello",

		"SayHello1",
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func AddEndpointMiddlewareToAllMethodsWithMethodName(mw map[string][]endpoint.Middleware, m func(n string) endpoint.Middleware) {
	methods := []string{

		"Foo",

		"Say",

		"Say1",

		"SayHello",

		"SayHello1",
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m(v))
	}
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpoints(s HelloService, mdw map[string][]endpoint.Middleware) Endpoints {
	eps := Endpoints{

		FooEndpoint: MakeFooEndpoint(s),

		SayEndpoint: MakeSayEndpoint(s),

		Say1Endpoint: MakeSay1Endpoint(s),

		SayHelloEndpoint: MakeSayHelloEndpoint(s),

		SayHello1Endpoint: MakeSayHello1Endpoint(s),
	}

	for _, m := range mdw["Foo"] {
		eps.FooEndpoint = m(eps.FooEndpoint)
	}

	for _, m := range mdw["Say"] {
		eps.SayEndpoint = m(eps.SayEndpoint)
	}

	for _, m := range mdw["Say1"] {
		eps.Say1Endpoint = m(eps.Say1Endpoint)
	}

	for _, m := range mdw["SayHello"] {
		eps.SayHelloEndpoint = m(eps.SayHelloEndpoint)
	}

	for _, m := range mdw["SayHello1"] {
		eps.SayHello1Endpoint = m(eps.SayHello1Endpoint)
	}

	return eps
}

func MakeFooEndpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.SayReq)
		rs, err := s.Foo(ctx, req)
		result := make(map[string]interface{}, 0)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeSayEndpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.SayReq)
		rs, err := s.Say(ctx, req)
		result := make(map[string]interface{}, 0)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeSay1Endpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.SayReq)
		rs, err := s.Say1(ctx, req)
		result := make(map[string]interface{}, 0)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeSayHelloEndpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.SayReq)
		rs, err := s.SayHello(ctx, req)
		result := make(map[string]interface{}, 0)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeSayHello1Endpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(types.SayReq)
		rs, err := s.SayHello1(ctx, req)
		result := make(map[string]interface{}, 0)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

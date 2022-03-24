package user

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	endpoint "github.com/go-kit/kit/endpoint"
)

type Mws map[string][]endpoint.Middleware

type Endpoints struct {
	GetByIdEndpoint endpoint.Endpoint

	GetListEndpoint endpoint.Endpoint
}

func AddEndpointMiddlewareToAllMethods(mw map[string][]endpoint.Middleware, m endpoint.Middleware) {
	methods := []string{

		"GetById",

		"GetList",
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func AddEndpointMiddlewareToAllMethodsWithMethodName(mw map[string][]endpoint.Middleware, m func(n string) endpoint.Middleware) {
	methods := []string{

		"GetById",

		"GetList",
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m(v))
	}
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpoints(s UserService, mdw Mws) Endpoints {
	eps := Endpoints{

		GetByIdEndpoint: MakeGetByIdEndpoint(s),

		GetListEndpoint: MakeGetListEndpoint(s),
	}

	for _, m := range mdw["GetById"] {
		eps.GetByIdEndpoint = m(eps.GetByIdEndpoint)
	}

	for _, m := range mdw["GetList"] {
		eps.GetListEndpoint = m(eps.GetListEndpoint)
	}

	return eps
}

func MakeGetByIdEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIdReq)
		rs, err := s.GetById(ctx, req)
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

func MakeGetListEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetListReq)
		rs, err := s.GetList(ctx, req)
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
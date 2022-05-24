package user

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"

	"go.opentelemetry.io/otel/trace"

	endpoint "github.com/go-kit/kit/endpoint"

	"hello/pkg/ent"
)

const (
	UserRestByQueriesAllMethodName = "UserRestByQueriesAll"

	UserRestCreateMethodName = "UserRestCreate"

	UserRestCreateManyMethodName = "UserRestCreateMany"

	UserRestDeleteByIdMethodName = "UserRestDeleteById"

	UserRestDeleteManyMethodName = "UserRestDeleteMany"

	UserRestGetByIdMethodName = "UserRestGetById"

	UserRestUpdateByIdMethodName = "UserRestUpdateById"

	UserRestUpdateManyMethodName = "UserRestUpdateMany"
)

type Mws map[string][]endpoint.Middleware

type Endpoints struct {
	UserRestByQueriesAllEndpoint endpoint.Endpoint

	UserRestCreateEndpoint endpoint.Endpoint

	UserRestCreateManyEndpoint endpoint.Endpoint

	UserRestDeleteByIdEndpoint endpoint.Endpoint

	UserRestDeleteManyEndpoint endpoint.Endpoint

	UserRestGetByIdEndpoint endpoint.Endpoint

	UserRestUpdateByIdEndpoint endpoint.Endpoint

	UserRestUpdateManyEndpoint endpoint.Endpoint
}

func AddEndpointMiddlewareToAllMethods(mw map[string][]endpoint.Middleware, m endpoint.Middleware) {
	methods := []string{

		UserRestByQueriesAllMethodName,

		UserRestCreateMethodName,

		UserRestCreateManyMethodName,

		UserRestDeleteByIdMethodName,

		UserRestDeleteManyMethodName,

		UserRestGetByIdMethodName,

		UserRestUpdateByIdMethodName,

		UserRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func AddEndpointMiddlewareToAllMethodsWithMethodName(mw map[string][]endpoint.Middleware, m func(n string) endpoint.Middleware) {
	methods := []string{

		UserRestByQueriesAllMethodName,

		UserRestCreateMethodName,

		UserRestCreateManyMethodName,

		UserRestDeleteByIdMethodName,

		UserRestDeleteManyMethodName,

		UserRestGetByIdMethodName,

		UserRestUpdateByIdMethodName,

		UserRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m(v))
	}
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpoints(s UserService, mdw Mws) Endpoints {
	eps := Endpoints{

		UserRestByQueriesAllEndpoint: MakeUserRestByQueriesAllEndpoint(s),

		UserRestCreateEndpoint: MakeUserRestCreateEndpoint(s),

		UserRestCreateManyEndpoint: MakeUserRestCreateManyEndpoint(s),

		UserRestDeleteByIdEndpoint: MakeUserRestDeleteByIdEndpoint(s),

		UserRestDeleteManyEndpoint: MakeUserRestDeleteManyEndpoint(s),

		UserRestGetByIdEndpoint: MakeUserRestGetByIdEndpoint(s),

		UserRestUpdateByIdEndpoint: MakeUserRestUpdateByIdEndpoint(s),

		UserRestUpdateManyEndpoint: MakeUserRestUpdateManyEndpoint(s),
	}

	for _, m := range mdw[UserRestByQueriesAllMethodName] {
		eps.UserRestByQueriesAllEndpoint = m(eps.UserRestByQueriesAllEndpoint)
	}

	for _, m := range mdw[UserRestCreateMethodName] {
		eps.UserRestCreateEndpoint = m(eps.UserRestCreateEndpoint)
	}

	for _, m := range mdw[UserRestCreateManyMethodName] {
		eps.UserRestCreateManyEndpoint = m(eps.UserRestCreateManyEndpoint)
	}

	for _, m := range mdw[UserRestDeleteByIdMethodName] {
		eps.UserRestDeleteByIdEndpoint = m(eps.UserRestDeleteByIdEndpoint)
	}

	for _, m := range mdw[UserRestDeleteManyMethodName] {
		eps.UserRestDeleteManyEndpoint = m(eps.UserRestDeleteManyEndpoint)
	}

	for _, m := range mdw[UserRestGetByIdMethodName] {
		eps.UserRestGetByIdEndpoint = m(eps.UserRestGetByIdEndpoint)
	}

	for _, m := range mdw[UserRestUpdateByIdMethodName] {
		eps.UserRestUpdateByIdEndpoint = m(eps.UserRestUpdateByIdEndpoint)
	}

	for _, m := range mdw[UserRestUpdateManyMethodName] {
		eps.UserRestUpdateManyEndpoint = m(eps.UserRestUpdateManyEndpoint)
	}

	return eps
}

func MakeUserRestByQueriesAllEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestByQueriesAllReq)
		rs, err := s.UserRestByQueriesAll(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestCreateEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestCreateReq)
		rs, err := s.UserRestCreate(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestCreateManyEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestCreateManyReq)
		rs, err := s.UserRestCreateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestDeleteByIdEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestDeleteByIdReq)
		rs, err := s.UserRestDeleteById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestDeleteManyEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestDeleteManyReq)
		rs, err := s.UserRestDeleteMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestGetByIdEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestGetByIdReq)
		rs, err := s.UserRestGetById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestUpdateByIdEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestUpdateByIdReq)
		rs, err := s.UserRestUpdateById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeUserRestUpdateManyEndpoint(s UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.UserRestUpdateManyReq)
		rs, err := s.UserRestUpdateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

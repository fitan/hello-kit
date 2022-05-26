package resource

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
	ResourceRestByQueriesAllMethodName = "ResourceRestByQueriesAll"

	ResourceRestCreateMethodName = "ResourceRestCreate"

	ResourceRestCreateManyMethodName = "ResourceRestCreateMany"

	ResourceRestDeleteByIdMethodName = "ResourceRestDeleteById"

	ResourceRestDeleteManyMethodName = "ResourceRestDeleteMany"

	ResourceRestGetByIdMethodName = "ResourceRestGetById"

	ResourceRestUpdateByIdMethodName = "ResourceRestUpdateById"

	ResourceRestUpdateManyMethodName = "ResourceRestUpdateMany"
)

type Mws map[string][]endpoint.Middleware

type Endpoints struct {
	ResourceRestByQueriesAllEndpoint endpoint.Endpoint

	ResourceRestCreateEndpoint endpoint.Endpoint

	ResourceRestCreateManyEndpoint endpoint.Endpoint

	ResourceRestDeleteByIdEndpoint endpoint.Endpoint

	ResourceRestDeleteManyEndpoint endpoint.Endpoint

	ResourceRestGetByIdEndpoint endpoint.Endpoint

	ResourceRestUpdateByIdEndpoint endpoint.Endpoint

	ResourceRestUpdateManyEndpoint endpoint.Endpoint
}

func AddEndpointMiddlewareToAllMethods(mw map[string][]endpoint.Middleware, m endpoint.Middleware) {
	methods := []string{

		ResourceRestByQueriesAllMethodName,

		ResourceRestCreateMethodName,

		ResourceRestCreateManyMethodName,

		ResourceRestDeleteByIdMethodName,

		ResourceRestDeleteManyMethodName,

		ResourceRestGetByIdMethodName,

		ResourceRestUpdateByIdMethodName,

		ResourceRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func AddEndpointMiddlewareToAllMethodsWithMethodName(mw map[string][]endpoint.Middleware, m func(n string) endpoint.Middleware) {
	methods := []string{

		ResourceRestByQueriesAllMethodName,

		ResourceRestCreateMethodName,

		ResourceRestCreateManyMethodName,

		ResourceRestDeleteByIdMethodName,

		ResourceRestDeleteManyMethodName,

		ResourceRestGetByIdMethodName,

		ResourceRestUpdateByIdMethodName,

		ResourceRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m(v))
	}
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpoints(s ResourceService, mdw Mws) Endpoints {
	eps := Endpoints{

		ResourceRestByQueriesAllEndpoint: MakeResourceRestByQueriesAllEndpoint(s),

		ResourceRestCreateEndpoint: MakeResourceRestCreateEndpoint(s),

		ResourceRestCreateManyEndpoint: MakeResourceRestCreateManyEndpoint(s),

		ResourceRestDeleteByIdEndpoint: MakeResourceRestDeleteByIdEndpoint(s),

		ResourceRestDeleteManyEndpoint: MakeResourceRestDeleteManyEndpoint(s),

		ResourceRestGetByIdEndpoint: MakeResourceRestGetByIdEndpoint(s),

		ResourceRestUpdateByIdEndpoint: MakeResourceRestUpdateByIdEndpoint(s),

		ResourceRestUpdateManyEndpoint: MakeResourceRestUpdateManyEndpoint(s),
	}

	for _, m := range mdw[ResourceRestByQueriesAllMethodName] {
		eps.ResourceRestByQueriesAllEndpoint = m(eps.ResourceRestByQueriesAllEndpoint)
	}

	for _, m := range mdw[ResourceRestCreateMethodName] {
		eps.ResourceRestCreateEndpoint = m(eps.ResourceRestCreateEndpoint)
	}

	for _, m := range mdw[ResourceRestCreateManyMethodName] {
		eps.ResourceRestCreateManyEndpoint = m(eps.ResourceRestCreateManyEndpoint)
	}

	for _, m := range mdw[ResourceRestDeleteByIdMethodName] {
		eps.ResourceRestDeleteByIdEndpoint = m(eps.ResourceRestDeleteByIdEndpoint)
	}

	for _, m := range mdw[ResourceRestDeleteManyMethodName] {
		eps.ResourceRestDeleteManyEndpoint = m(eps.ResourceRestDeleteManyEndpoint)
	}

	for _, m := range mdw[ResourceRestGetByIdMethodName] {
		eps.ResourceRestGetByIdEndpoint = m(eps.ResourceRestGetByIdEndpoint)
	}

	for _, m := range mdw[ResourceRestUpdateByIdMethodName] {
		eps.ResourceRestUpdateByIdEndpoint = m(eps.ResourceRestUpdateByIdEndpoint)
	}

	for _, m := range mdw[ResourceRestUpdateManyMethodName] {
		eps.ResourceRestUpdateManyEndpoint = m(eps.ResourceRestUpdateManyEndpoint)
	}

	return eps
}

func MakeResourceRestByQueriesAllEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestByQueriesAllReq)
		rs, err := s.ResourceRestByQueriesAll(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestCreateEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestCreateReq)
		rs, err := s.ResourceRestCreate(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestCreateManyEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestCreateManyReq)
		rs, err := s.ResourceRestCreateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestDeleteByIdEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestDeleteByIdReq)
		rs, err := s.ResourceRestDeleteById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestDeleteManyEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestDeleteManyReq)
		rs, err := s.ResourceRestDeleteMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestGetByIdEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestGetByIdReq)
		rs, err := s.ResourceRestGetById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestUpdateByIdEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestUpdateByIdReq)
		rs, err := s.ResourceRestUpdateById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeResourceRestUpdateManyEndpoint(s ResourceService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ResourceRestUpdateManyReq)
		rs, err := s.ResourceRestUpdateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

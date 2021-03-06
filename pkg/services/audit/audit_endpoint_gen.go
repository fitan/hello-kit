package audit

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
	AuditRestByQueriesAllMethodName = "AuditRestByQueriesAll"

	AuditRestCreateMethodName = "AuditRestCreate"

	AuditRestCreateManyMethodName = "AuditRestCreateMany"

	AuditRestDeleteByIdMethodName = "AuditRestDeleteById"

	AuditRestDeleteManyMethodName = "AuditRestDeleteMany"

	AuditRestGetByIdMethodName = "AuditRestGetById"

	AuditRestUpdateByIdMethodName = "AuditRestUpdateById"

	AuditRestUpdateManyMethodName = "AuditRestUpdateMany"
)

type Mws map[string][]endpoint.Middleware

type Endpoints struct {
	AuditRestByQueriesAllEndpoint endpoint.Endpoint

	AuditRestCreateEndpoint endpoint.Endpoint

	AuditRestCreateManyEndpoint endpoint.Endpoint

	AuditRestDeleteByIdEndpoint endpoint.Endpoint

	AuditRestDeleteManyEndpoint endpoint.Endpoint

	AuditRestGetByIdEndpoint endpoint.Endpoint

	AuditRestUpdateByIdEndpoint endpoint.Endpoint

	AuditRestUpdateManyEndpoint endpoint.Endpoint
}

func AddEndpointMiddlewareToAllMethods(mw map[string][]endpoint.Middleware, m endpoint.Middleware) {
	methods := []string{

		AuditRestByQueriesAllMethodName,

		AuditRestCreateMethodName,

		AuditRestCreateManyMethodName,

		AuditRestDeleteByIdMethodName,

		AuditRestDeleteManyMethodName,

		AuditRestGetByIdMethodName,

		AuditRestUpdateByIdMethodName,

		AuditRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func AddEndpointMiddlewareToAllMethodsWithMethodName(mw map[string][]endpoint.Middleware, m func(n string) endpoint.Middleware) {
	methods := []string{

		AuditRestByQueriesAllMethodName,

		AuditRestCreateMethodName,

		AuditRestCreateManyMethodName,

		AuditRestDeleteByIdMethodName,

		AuditRestDeleteManyMethodName,

		AuditRestGetByIdMethodName,

		AuditRestUpdateByIdMethodName,

		AuditRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m(v))
	}
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpoints(s AuditService, mdw Mws) Endpoints {
	eps := Endpoints{

		AuditRestByQueriesAllEndpoint: MakeAuditRestByQueriesAllEndpoint(s),

		AuditRestCreateEndpoint: MakeAuditRestCreateEndpoint(s),

		AuditRestCreateManyEndpoint: MakeAuditRestCreateManyEndpoint(s),

		AuditRestDeleteByIdEndpoint: MakeAuditRestDeleteByIdEndpoint(s),

		AuditRestDeleteManyEndpoint: MakeAuditRestDeleteManyEndpoint(s),

		AuditRestGetByIdEndpoint: MakeAuditRestGetByIdEndpoint(s),

		AuditRestUpdateByIdEndpoint: MakeAuditRestUpdateByIdEndpoint(s),

		AuditRestUpdateManyEndpoint: MakeAuditRestUpdateManyEndpoint(s),
	}

	for _, m := range mdw[AuditRestByQueriesAllMethodName] {
		eps.AuditRestByQueriesAllEndpoint = m(eps.AuditRestByQueriesAllEndpoint)
	}

	for _, m := range mdw[AuditRestCreateMethodName] {
		eps.AuditRestCreateEndpoint = m(eps.AuditRestCreateEndpoint)
	}

	for _, m := range mdw[AuditRestCreateManyMethodName] {
		eps.AuditRestCreateManyEndpoint = m(eps.AuditRestCreateManyEndpoint)
	}

	for _, m := range mdw[AuditRestDeleteByIdMethodName] {
		eps.AuditRestDeleteByIdEndpoint = m(eps.AuditRestDeleteByIdEndpoint)
	}

	for _, m := range mdw[AuditRestDeleteManyMethodName] {
		eps.AuditRestDeleteManyEndpoint = m(eps.AuditRestDeleteManyEndpoint)
	}

	for _, m := range mdw[AuditRestGetByIdMethodName] {
		eps.AuditRestGetByIdEndpoint = m(eps.AuditRestGetByIdEndpoint)
	}

	for _, m := range mdw[AuditRestUpdateByIdMethodName] {
		eps.AuditRestUpdateByIdEndpoint = m(eps.AuditRestUpdateByIdEndpoint)
	}

	for _, m := range mdw[AuditRestUpdateManyMethodName] {
		eps.AuditRestUpdateManyEndpoint = m(eps.AuditRestUpdateManyEndpoint)
	}

	return eps
}

func MakeAuditRestByQueriesAllEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestByQueriesAllReq)
		rs, err := s.AuditRestByQueriesAll(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestCreateEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestCreateReq)
		rs, err := s.AuditRestCreate(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestCreateManyEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestCreateManyReq)
		rs, err := s.AuditRestCreateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestDeleteByIdEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestDeleteByIdReq)
		rs, err := s.AuditRestDeleteById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestDeleteManyEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestDeleteManyReq)
		rs, err := s.AuditRestDeleteMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestGetByIdEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestGetByIdReq)
		rs, err := s.AuditRestGetById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestUpdateByIdEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestUpdateByIdReq)
		rs, err := s.AuditRestUpdateById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeAuditRestUpdateManyEndpoint(s AuditService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.AuditRestUpdateManyReq)
		rs, err := s.AuditRestUpdateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

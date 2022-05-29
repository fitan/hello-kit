package project

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
	ProjectRestAddBindServicesByProjectIdMethodName = "ProjectRestAddBindServicesByProjectId"

	ProjectRestByQueriesAllMethodName = "ProjectRestByQueriesAll"

	ProjectRestCreateMethodName = "ProjectRestCreate"

	ProjectRestCreateManyMethodName = "ProjectRestCreateMany"

	ProjectRestCreateServicesByProjectIdMethodName = "ProjectRestCreateServicesByProjectId"

	ProjectRestDeleteByIdMethodName = "ProjectRestDeleteById"

	ProjectRestDeleteManyMethodName = "ProjectRestDeleteMany"

	ProjectRestDeleteServicesByProjectIdMethodName = "ProjectRestDeleteServicesByProjectId"

	ProjectRestGetByIdMethodName = "ProjectRestGetById"

	ProjectRestGetServicesByProjectIdMethodName = "ProjectRestGetServicesByProjectId"

	ProjectRestRemoveBindServicesByProjectIdMethodName = "ProjectRestRemoveBindServicesByProjectId"

	ProjectRestUpdateBindServicesByProjectIdMethodName = "ProjectRestUpdateBindServicesByProjectId"

	ProjectRestUpdateByIdMethodName = "ProjectRestUpdateById"

	ProjectRestUpdateManyMethodName = "ProjectRestUpdateMany"
)

type Mws map[string][]endpoint.Middleware

type Endpoints struct {
	ProjectRestAddBindServicesByProjectIdEndpoint endpoint.Endpoint

	ProjectRestByQueriesAllEndpoint endpoint.Endpoint

	ProjectRestCreateEndpoint endpoint.Endpoint

	ProjectRestCreateManyEndpoint endpoint.Endpoint

	ProjectRestCreateServicesByProjectIdEndpoint endpoint.Endpoint

	ProjectRestDeleteByIdEndpoint endpoint.Endpoint

	ProjectRestDeleteManyEndpoint endpoint.Endpoint

	ProjectRestDeleteServicesByProjectIdEndpoint endpoint.Endpoint

	ProjectRestGetByIdEndpoint endpoint.Endpoint

	ProjectRestGetServicesByProjectIdEndpoint endpoint.Endpoint

	ProjectRestRemoveBindServicesByProjectIdEndpoint endpoint.Endpoint

	ProjectRestUpdateBindServicesByProjectIdEndpoint endpoint.Endpoint

	ProjectRestUpdateByIdEndpoint endpoint.Endpoint

	ProjectRestUpdateManyEndpoint endpoint.Endpoint
}

func AddEndpointMiddlewareToAllMethods(mw map[string][]endpoint.Middleware, m endpoint.Middleware) {
	methods := []string{

		ProjectRestAddBindServicesByProjectIdMethodName,

		ProjectRestByQueriesAllMethodName,

		ProjectRestCreateMethodName,

		ProjectRestCreateManyMethodName,

		ProjectRestCreateServicesByProjectIdMethodName,

		ProjectRestDeleteByIdMethodName,

		ProjectRestDeleteManyMethodName,

		ProjectRestDeleteServicesByProjectIdMethodName,

		ProjectRestGetByIdMethodName,

		ProjectRestGetServicesByProjectIdMethodName,

		ProjectRestRemoveBindServicesByProjectIdMethodName,

		ProjectRestUpdateBindServicesByProjectIdMethodName,

		ProjectRestUpdateByIdMethodName,

		ProjectRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

func AddEndpointMiddlewareToAllMethodsWithMethodName(mw map[string][]endpoint.Middleware, m func(n string) endpoint.Middleware) {
	methods := []string{

		ProjectRestAddBindServicesByProjectIdMethodName,

		ProjectRestByQueriesAllMethodName,

		ProjectRestCreateMethodName,

		ProjectRestCreateManyMethodName,

		ProjectRestCreateServicesByProjectIdMethodName,

		ProjectRestDeleteByIdMethodName,

		ProjectRestDeleteManyMethodName,

		ProjectRestDeleteServicesByProjectIdMethodName,

		ProjectRestGetByIdMethodName,

		ProjectRestGetServicesByProjectIdMethodName,

		ProjectRestRemoveBindServicesByProjectIdMethodName,

		ProjectRestUpdateBindServicesByProjectIdMethodName,

		ProjectRestUpdateByIdMethodName,

		ProjectRestUpdateManyMethodName,
	}
	for _, v := range methods {
		mw[v] = append(mw[v], m(v))
	}
}

// New returns a Endpoints struct that wraps the provided service, and wires in all of the
// expected endpoint middlewares
func NewEndpoints(s ProjectService, mdw Mws) Endpoints {
	eps := Endpoints{

		ProjectRestAddBindServicesByProjectIdEndpoint: MakeProjectRestAddBindServicesByProjectIdEndpoint(s),

		ProjectRestByQueriesAllEndpoint: MakeProjectRestByQueriesAllEndpoint(s),

		ProjectRestCreateEndpoint: MakeProjectRestCreateEndpoint(s),

		ProjectRestCreateManyEndpoint: MakeProjectRestCreateManyEndpoint(s),

		ProjectRestCreateServicesByProjectIdEndpoint: MakeProjectRestCreateServicesByProjectIdEndpoint(s),

		ProjectRestDeleteByIdEndpoint: MakeProjectRestDeleteByIdEndpoint(s),

		ProjectRestDeleteManyEndpoint: MakeProjectRestDeleteManyEndpoint(s),

		ProjectRestDeleteServicesByProjectIdEndpoint: MakeProjectRestDeleteServicesByProjectIdEndpoint(s),

		ProjectRestGetByIdEndpoint: MakeProjectRestGetByIdEndpoint(s),

		ProjectRestGetServicesByProjectIdEndpoint: MakeProjectRestGetServicesByProjectIdEndpoint(s),

		ProjectRestRemoveBindServicesByProjectIdEndpoint: MakeProjectRestRemoveBindServicesByProjectIdEndpoint(s),

		ProjectRestUpdateBindServicesByProjectIdEndpoint: MakeProjectRestUpdateBindServicesByProjectIdEndpoint(s),

		ProjectRestUpdateByIdEndpoint: MakeProjectRestUpdateByIdEndpoint(s),

		ProjectRestUpdateManyEndpoint: MakeProjectRestUpdateManyEndpoint(s),
	}

	for _, m := range mdw[ProjectRestAddBindServicesByProjectIdMethodName] {
		eps.ProjectRestAddBindServicesByProjectIdEndpoint = m(eps.ProjectRestAddBindServicesByProjectIdEndpoint)
	}

	for _, m := range mdw[ProjectRestByQueriesAllMethodName] {
		eps.ProjectRestByQueriesAllEndpoint = m(eps.ProjectRestByQueriesAllEndpoint)
	}

	for _, m := range mdw[ProjectRestCreateMethodName] {
		eps.ProjectRestCreateEndpoint = m(eps.ProjectRestCreateEndpoint)
	}

	for _, m := range mdw[ProjectRestCreateManyMethodName] {
		eps.ProjectRestCreateManyEndpoint = m(eps.ProjectRestCreateManyEndpoint)
	}

	for _, m := range mdw[ProjectRestCreateServicesByProjectIdMethodName] {
		eps.ProjectRestCreateServicesByProjectIdEndpoint = m(eps.ProjectRestCreateServicesByProjectIdEndpoint)
	}

	for _, m := range mdw[ProjectRestDeleteByIdMethodName] {
		eps.ProjectRestDeleteByIdEndpoint = m(eps.ProjectRestDeleteByIdEndpoint)
	}

	for _, m := range mdw[ProjectRestDeleteManyMethodName] {
		eps.ProjectRestDeleteManyEndpoint = m(eps.ProjectRestDeleteManyEndpoint)
	}

	for _, m := range mdw[ProjectRestDeleteServicesByProjectIdMethodName] {
		eps.ProjectRestDeleteServicesByProjectIdEndpoint = m(eps.ProjectRestDeleteServicesByProjectIdEndpoint)
	}

	for _, m := range mdw[ProjectRestGetByIdMethodName] {
		eps.ProjectRestGetByIdEndpoint = m(eps.ProjectRestGetByIdEndpoint)
	}

	for _, m := range mdw[ProjectRestGetServicesByProjectIdMethodName] {
		eps.ProjectRestGetServicesByProjectIdEndpoint = m(eps.ProjectRestGetServicesByProjectIdEndpoint)
	}

	for _, m := range mdw[ProjectRestRemoveBindServicesByProjectIdMethodName] {
		eps.ProjectRestRemoveBindServicesByProjectIdEndpoint = m(eps.ProjectRestRemoveBindServicesByProjectIdEndpoint)
	}

	for _, m := range mdw[ProjectRestUpdateBindServicesByProjectIdMethodName] {
		eps.ProjectRestUpdateBindServicesByProjectIdEndpoint = m(eps.ProjectRestUpdateBindServicesByProjectIdEndpoint)
	}

	for _, m := range mdw[ProjectRestUpdateByIdMethodName] {
		eps.ProjectRestUpdateByIdEndpoint = m(eps.ProjectRestUpdateByIdEndpoint)
	}

	for _, m := range mdw[ProjectRestUpdateManyMethodName] {
		eps.ProjectRestUpdateManyEndpoint = m(eps.ProjectRestUpdateManyEndpoint)
	}

	return eps
}

func MakeProjectRestAddBindServicesByProjectIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestAddBindServicesByProjectIdReq)
		rs, err := s.ProjectRestAddBindServicesByProjectId(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestByQueriesAllEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestByQueriesAllReq)
		rs, err := s.ProjectRestByQueriesAll(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestCreateEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestCreateReq)
		rs, err := s.ProjectRestCreate(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestCreateManyEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestCreateManyReq)
		rs, err := s.ProjectRestCreateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestCreateServicesByProjectIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestCreateServicesByProjectIdReq)
		rs, err := s.ProjectRestCreateServicesByProjectId(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestDeleteByIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestDeleteByIdReq)
		rs, err := s.ProjectRestDeleteById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestDeleteManyEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestDeleteManyReq)
		rs, err := s.ProjectRestDeleteMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestDeleteServicesByProjectIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestDeleteServicesByProjectIdReq)
		rs, err := s.ProjectRestDeleteServicesByProjectId(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestGetByIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestGetByIdReq)
		rs, err := s.ProjectRestGetById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestGetServicesByProjectIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestGetServicesByProjectIdReq)
		rs, err := s.ProjectRestGetServicesByProjectId(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestRemoveBindServicesByProjectIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestRemoveBindServicesByProjectIdReq)
		rs, err := s.ProjectRestRemoveBindServicesByProjectId(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestUpdateBindServicesByProjectIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestUpdateBindServicesByProjectIdReq)
		rs, err := s.ProjectRestUpdateBindServicesByProjectId(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestUpdateByIdEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestUpdateByIdReq)
		rs, err := s.ProjectRestUpdateById(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

func MakeProjectRestUpdateManyEndpoint(s ProjectService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		result := make(map[string]interface{}, 0)
		defer func() {
			if err != nil {
				result["status"] = 500
			}
			result["status"] = 0
		}()
		req := request.(ent.ProjectRestUpdateManyReq)
		rs, err := s.ProjectRestUpdateMany(ctx, req)
		if err != nil {
			result["err"] = err.Error()
			return result, nil
		}
		result["data"] = rs
		result["traceId"] = trace.SpanFromContext(ctx).SpanContext().TraceID().String()
		return result, nil
	}
}

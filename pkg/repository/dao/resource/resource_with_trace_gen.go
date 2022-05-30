package resource

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"encoding/json"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	"context"

	"hello/pkg/ent"
)

// ResourceServiceWithTracing implements ResourceService interface instrumented with opentracing spans
type ResourceServiceWithTracing struct {
	ResourceService
}

// NewResourceServiceWithTracing returns ResourceServiceWithTracing
func NewResourceServiceWithTracing(base ResourceService) ResourceService {
	d := ResourceServiceWithTracing{
		ResourceService: base,
	}

	return d
}

// ByQueriesAll implements ResourceService
func (_d ResourceServiceWithTracing) ByQueriesAll(ctx context.Context, i interface{}) (res []ent.ResourceBaseGetRes, count int, err error) {

	var name = "ResourceService.ByQueriesAll"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"i": i},
				"result": map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.ByQueriesAll(ctx, i)
}

// ByQueriesOne implements ResourceService
func (_d ResourceServiceWithTracing) ByQueriesOne(ctx context.Context, i interface{}) (res ent.ResourceBaseGetRes, err error) {

	var name = "ResourceService.ByQueriesOne"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"i": i},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.ByQueriesOne(ctx, i)
}

// Create implements ResourceService
func (_d ResourceServiceWithTracing) Create(ctx context.Context, v ent.ResourceBaseCreateReq) (res *ent.Resource, err error) {

	var name = "ResourceService.Create"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"v": v},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.Create(ctx, v)
}

// CreateMany implements ResourceService
func (_d ResourceServiceWithTracing) CreateMany(ctx context.Context, vs []ent.ResourceBaseCreateReq) (r1 ent.Resources, err error) {

	var name = "ResourceService.CreateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"r1":  r1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.CreateMany(ctx, vs)
}

// CreatePreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) CreatePreByResourceId(ctx context.Context, id int, v ent.ResourceBaseCreateReq) (res *ent.Resource, err error) {

	var name = "ResourceService.CreatePreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.CreatePreByResourceId(ctx, id, v)
}

// CreateResourcesByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) CreateResourcesByResourceId(ctx context.Context, id int, vs []ent.ResourceBaseCreateReq) (res *ent.Resource, err error) {

	var name = "ResourceService.CreateResourcesByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"vs": vs},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.CreateResourcesByResourceId(ctx, id, vs)
}

// DeleteById implements ResourceService
func (_d ResourceServiceWithTracing) DeleteById(ctx context.Context, id int) (err error) {

	var name = "ResourceService.DeleteById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.DeleteById(ctx, id)
}

// DeleteMany implements ResourceService
func (_d ResourceServiceWithTracing) DeleteMany(ctx context.Context, ids []int) (err error) {

	var name = "ResourceService.DeleteMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"ids": ids},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.DeleteMany(ctx, ids)
}

// GetById implements ResourceService
func (_d ResourceServiceWithTracing) GetById(ctx context.Context, id int) (res ent.ResourceBaseGetRes, err error) {

	var name = "ResourceService.GetById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.GetById(ctx, id)
}

// GetNextByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) GetNextByResourceId(ctx context.Context, id int, i interface{}) (res []ent.ResourceBaseGetRes, count int, err error) {

	var name = "ResourceService.GetNextByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"i":  i},
				"result": map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.GetNextByResourceId(ctx, id, i)
}

// GetPreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) GetPreByResourceId(ctx context.Context, id int) (res ent.ResourceBaseGetRes, err error) {

	var name = "ResourceService.GetPreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.GetPreByResourceId(ctx, id)
}

// IdByResource implements ResourceService
func (_d ResourceServiceWithTracing) IdByResource(ctx context.Context, path string, action string) (i1 int, err error) {

	var name = "ResourceService.IdByResource"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"path":   path,
					"action": action},
				"result": map[string]interface{}{
					"i1":  i1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.IdByResource(ctx, path, action)
}

// RawAddBindNextByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawAddBindNextByResourceId(ctx context.Context, id int, addIds []int) (err error) {

	var name = "ResourceService.RawAddBindNextByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":     id,
					"addIds": addIds},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawAddBindNextByResourceId(ctx, id, addIds)
}

// RawAddBindPreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawAddBindPreByResourceId(ctx context.Context, id int, addId int) (err error) {

	var name = "ResourceService.RawAddBindPreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":    id,
					"addId": addId},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawAddBindPreByResourceId(ctx, id, addId)
}

// RawByQueriesAll implements ResourceService
func (_d ResourceServiceWithTracing) RawByQueriesAll(ctx context.Context, i interface{}) (res ent.Resources, count int, err error) {

	var name = "ResourceService.RawByQueriesAll"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"i": i},
				"result": map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawByQueriesAll(ctx, i)
}

// RawByQueriesOne implements ResourceService
func (_d ResourceServiceWithTracing) RawByQueriesOne(ctx context.Context, i interface{}) (res *ent.Resource, err error) {

	var name = "ResourceService.RawByQueriesOne"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"i": i},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawByQueriesOne(ctx, i)
}

// RawCreate implements ResourceService
func (_d ResourceServiceWithTracing) RawCreate(ctx context.Context, v *ent.Resource) (res *ent.Resource, err error) {

	var name = "ResourceService.RawCreate"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"v": v},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawCreate(ctx, v)
}

// RawCreateMany implements ResourceService
func (_d ResourceServiceWithTracing) RawCreateMany(ctx context.Context, vs ent.Resources) (r1 ent.Resources, err error) {

	var name = "ResourceService.RawCreateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"r1":  r1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawCreateMany(ctx, vs)
}

// RawCreatePreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawCreatePreByResourceId(ctx context.Context, id int, v *ent.Resource) (res *ent.Resource, err error) {

	var name = "ResourceService.RawCreatePreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawCreatePreByResourceId(ctx, id, v)
}

// RawCreateResourcesByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawCreateResourcesByResourceId(ctx context.Context, id int, vs ent.Resources) (res *ent.Resource, err error) {

	var name = "ResourceService.RawCreateResourcesByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"vs": vs},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawCreateResourcesByResourceId(ctx, id, vs)
}

// RawDeleteNextByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawDeleteNextByResourceId(ctx context.Context, id int, deleteIds []int) (err error) {

	var name = "ResourceService.RawDeleteNextByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":        id,
					"deleteIds": deleteIds},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawDeleteNextByResourceId(ctx, id, deleteIds)
}

// RawDeletePreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawDeletePreByResourceId(ctx context.Context, id int, deleteId int) (err error) {

	var name = "ResourceService.RawDeletePreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":       id,
					"deleteId": deleteId},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawDeletePreByResourceId(ctx, id, deleteId)
}

// RawGetById implements ResourceService
func (_d ResourceServiceWithTracing) RawGetById(ctx context.Context, id int) (res *ent.Resource, err error) {

	var name = "ResourceService.RawGetById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawGetById(ctx, id)
}

// RawGetNextByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawGetNextByResourceId(ctx context.Context, id int, i interface{}) (res ent.Resources, count int, err error) {

	var name = "ResourceService.RawGetNextByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"i":  i},
				"result": map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawGetNextByResourceId(ctx, id, i)
}

// RawGetPreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawGetPreByResourceId(ctx context.Context, id int) (res *ent.Resource, err error) {

	var name = "ResourceService.RawGetPreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id},
				"result": map[string]interface{}{
					"res": res,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawGetPreByResourceId(ctx, id)
}

// RawRemoveBindNextByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawRemoveBindNextByResourceId(ctx context.Context, id int, removeIds []int) (err error) {

	var name = "ResourceService.RawRemoveBindNextByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":        id,
					"removeIds": removeIds},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawRemoveBindNextByResourceId(ctx, id, removeIds)
}

// RawRemoveBindPreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawRemoveBindPreByResourceId(ctx context.Context, id int) (err error) {

	var name = "ResourceService.RawRemoveBindPreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawRemoveBindPreByResourceId(ctx, id)
}

// RawUpdateBindNextByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawUpdateBindNextByResourceId(ctx context.Context, id int, removeIds []int, addIds []int) (err error) {

	var name = "ResourceService.RawUpdateBindNextByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":        id,
					"removeIds": removeIds,
					"addIds":    addIds},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawUpdateBindNextByResourceId(ctx, id, removeIds, addIds)
}

// RawUpdateBindPreByResourceId implements ResourceService
func (_d ResourceServiceWithTracing) RawUpdateBindPreByResourceId(ctx context.Context, id int, updateId int) (err error) {

	var name = "ResourceService.RawUpdateBindPreByResourceId"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id":       id,
					"updateId": updateId},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawUpdateBindPreByResourceId(ctx, id, updateId)
}

// RawUpdateById implements ResourceService
func (_d ResourceServiceWithTracing) RawUpdateById(ctx context.Context, id int, v *ent.Resource) (rp1 *ent.Resource, err error) {

	var name = "ResourceService.RawUpdateById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"rp1": rp1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawUpdateById(ctx, id, v)
}

// RawUpdateMany implements ResourceService
func (_d ResourceServiceWithTracing) RawUpdateMany(ctx context.Context, vs ent.Resources) (err error) {

	var name = "ResourceService.RawUpdateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.RawUpdateMany(ctx, vs)
}

// UpdateById implements ResourceService
func (_d ResourceServiceWithTracing) UpdateById(ctx context.Context, id int, v ent.ResourceBaseUpdateReq) (rp1 *ent.Resource, err error) {

	var name = "ResourceService.UpdateById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"rp1": rp1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.UpdateById(ctx, id, v)
}

// UpdateMany implements ResourceService
func (_d ResourceServiceWithTracing) UpdateMany(ctx context.Context, vs []ent.ResourceBaseUpdateReq) (err error) {

	var name = "ResourceService.UpdateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.ResourceService.UpdateMany(ctx, vs)
}

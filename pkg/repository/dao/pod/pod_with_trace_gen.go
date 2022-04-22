package pod

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"encoding/json"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	"hello/pkg/ent"
)

// PodServiceWithTracing implements PodService interface instrumented with opentracing spans
type PodServiceWithTracing struct {
	PodService
}

// NewPodServiceWithTracing returns PodServiceWithTracing
func NewPodServiceWithTracing(base PodService) PodService {
	d := PodServiceWithTracing{
		PodService: base,
	}

	return d
}

// ByQueries implements PodService
func (_d PodServiceWithTracing) ByQueries(ctx context.Context, i interface{}) (res []ent.PodBaseGetRes, count int, err error) {

	var name = "PodService.ByQueries"
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

	return _d.PodService.ByQueries(ctx, i)
}

// Create implements PodService
func (_d PodServiceWithTracing) Create(ctx context.Context, v ent.PodBaseCreateReq) (res *ent.Pod, err error) {

	var name = "PodService.Create"
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

	return _d.PodService.Create(ctx, v)
}

// CreateMany implements PodService
func (_d PodServiceWithTracing) CreateMany(ctx context.Context, vs []ent.PodBaseCreateReq) (p1 ent.Pods, err error) {

	var name = "PodService.CreateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"p1":  p1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.PodService.CreateMany(ctx, vs)
}

// CreateServicetreeByPodId implements PodService
func (_d PodServiceWithTracing) CreateServicetreeByPodId(ctx context.Context, id int64, v ent.SpiderDevTblServicetreeBaseCreateReq) (res *ent.Pod, err error) {

	var name = "PodService.CreateServicetreeByPodId"
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

	return _d.PodService.CreateServicetreeByPodId(ctx, id, v)
}

// DeleteById implements PodService
func (_d PodServiceWithTracing) DeleteById(ctx context.Context, id int64) (err error) {

	var name = "PodService.DeleteById"
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

	return _d.PodService.DeleteById(ctx, id)
}

// DeleteMany implements PodService
func (_d PodServiceWithTracing) DeleteMany(ctx context.Context, ids []int64) (err error) {

	var name = "PodService.DeleteMany"
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

	return _d.PodService.DeleteMany(ctx, ids)
}

// GetById implements PodService
func (_d PodServiceWithTracing) GetById(ctx context.Context, id int64) (res ent.PodBaseGetRes, err error) {

	var name = "PodService.GetById"
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

	return _d.PodService.GetById(ctx, id)
}

// GetServicetreeByPodId implements PodService
func (_d PodServiceWithTracing) GetServicetreeByPodId(ctx context.Context, id int64) (res ent.SpiderDevTblServicetreeBaseGetRes, err error) {

	var name = "PodService.GetServicetreeByPodId"
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

	return _d.PodService.GetServicetreeByPodId(ctx, id)
}

// UpdateById implements PodService
func (_d PodServiceWithTracing) UpdateById(ctx context.Context, id int64, v ent.PodBaseUpdateReq) (pp1 *ent.Pod, err error) {

	var name = "PodService.UpdateById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"pp1": pp1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.PodService.UpdateById(ctx, id, v)
}

// UpdateMany implements PodService
func (_d PodServiceWithTracing) UpdateMany(ctx context.Context, vs []ent.PodBaseUpdateReq) (err error) {

	var name = "PodService.UpdateMany"
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

	return _d.PodService.UpdateMany(ctx, vs)
}

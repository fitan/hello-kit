package tblservicetree

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"encoding/json"
	"hello/pkg/ent"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// TblservicetreeServiceWithTracing implements TblservicetreeService interface instrumented with opentracing spans
type TblservicetreeServiceWithTracing struct {
	TblservicetreeService
}

// NewTblservicetreeServiceWithTracing returns TblservicetreeServiceWithTracing
func NewTblservicetreeServiceWithTracing(base TblservicetreeService) TblservicetreeService {
	d := TblservicetreeServiceWithTracing{
		TblservicetreeService: base,
	}

	return d
}

// ByQueries implements TblservicetreeService
func (_d TblservicetreeServiceWithTracing) ByQueries(ctx context.Context, i interface{}) (res ent.SpiderDevTblServicetrees, count int, err error) {
	var name = "TblservicetreeService.ByQueries"
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
	return _d.TblservicetreeService.ByQueries(ctx, i)
}

// Create implements TblservicetreeService
func (_d TblservicetreeServiceWithTracing) Create(ctx context.Context, v ent.SpiderDevTblServicetree) (res *ent.SpiderDevTblServicetree, err error) {
	var name = "TblservicetreeService.Create"
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
	return _d.TblservicetreeService.Create(ctx, v)
}

// DeleteById implements TblservicetreeService
func (_d TblservicetreeServiceWithTracing) DeleteById(ctx context.Context, id int32) (err error) {
	var name = "TblservicetreeService.DeleteById"
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
	return _d.TblservicetreeService.DeleteById(ctx, id)
}

// GetById implements TblservicetreeService
func (_d TblservicetreeServiceWithTracing) GetById(ctx context.Context, id int32) (res *ent.SpiderDevTblServicetree, err error) {
	var name = "TblservicetreeService.GetById"
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
	return _d.TblservicetreeService.GetById(ctx, id)
}

// UpdateById implements TblservicetreeService
func (_d TblservicetreeServiceWithTracing) UpdateById(ctx context.Context, id int32, v *ent.SpiderDevTblServicetree) (sp1 *ent.SpiderDevTblServicetree, err error) {
	var name = "TblservicetreeService.UpdateById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"sp1": sp1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()
	return _d.TblservicetreeService.UpdateById(ctx, id, v)
}

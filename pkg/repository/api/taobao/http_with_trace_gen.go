package taobao

// Code generated by gowrap. DO NOT EDIT.
// template: ../../../gowrap/templates/opentracing.tmpl
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// TaobaoApiWithTracing implements TaobaoApi interface instrumented with opentracing spans
type TaobaoApiWithTracing struct {
	TaobaoApi
}

// NewTaobaoApiWithTracing returns TaobaoApiWithTracing
func NewTaobaoApiWithTracing(base TaobaoApi) TaobaoApi {
	d := TaobaoApiWithTracing{
		TaobaoApi: base,
	}

	return d
}

// GetRoot implements TaobaoApi
func (_d TaobaoApiWithTracing) GetRoot(ctx context.Context) (rp1 *resty.Response, err error) {
	var name = "TaobaoApi.GetRoot"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{},
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
	return _d.TaobaoApi.GetRoot(ctx)
}

// GetRoot1 implements TaobaoApi
func (_d TaobaoApiWithTracing) GetRoot1(ctx context.Context) (rp1 *resty.Response, err error) {
	var name = "TaobaoApi.GetRoot1"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{},
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
	return _d.TaobaoApi.GetRoot1(ctx)
}
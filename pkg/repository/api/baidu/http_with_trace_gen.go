package baidu

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

// BaiduApiWithTracing implements BaiduApi interface instrumented with opentracing spans
type BaiduApiWithTracing struct {
	BaiduApi
}

// NewBaiduApiWithTracing returns BaiduApiWithTracing
func NewBaiduApiWithTracing(base BaiduApi) BaiduApi {
	d := BaiduApiWithTracing{
		BaiduApi: base,
	}

	return d
}

// GetRoot implements BaiduApi
func (_d BaiduApiWithTracing) GetRoot(ctx context.Context) (rp1 *resty.Response, err error) {
	var name = "BaiduApi.GetRoot"
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
	return _d.BaiduApi.GetRoot(ctx)
}

// GetRoot1 implements BaiduApi
func (_d BaiduApiWithTracing) GetRoot1(ctx context.Context) (rp1 *resty.Response, err error) {
	var name = "BaiduApi.GetRoot1"
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
	return _d.BaiduApi.GetRoot1(ctx)
}
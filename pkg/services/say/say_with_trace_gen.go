package say

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"encoding/json"
	"hello/pkg/services/hello/types"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

// SayServiceWithTracing implements SayService interface instrumented with opentracing spans
type SayServiceWithTracing struct {
	SayService
}

// NewSayServiceWithTracing returns SayServiceWithTracing
func NewSayServiceWithTracing(base SayService) SayService {
	d := SayServiceWithTracing{
		SayService: base,
	}

	return d
}

// Say implements SayService
func (_d SayServiceWithTracing) Say(ctx context.Context, req types.SayReq) (res types.SayRes, err error) {
	var name = "SayService.Say"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"req": req},
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
	return _d.SayService.Say(ctx, req)
}

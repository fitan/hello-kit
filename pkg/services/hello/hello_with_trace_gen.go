package hello

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

// HelloServiceWithTracing implements HelloService interface instrumented with opentracing spans
type HelloServiceWithTracing struct {
	HelloService
}

// NewHelloServiceWithTracing returns HelloServiceWithTracing
func NewHelloServiceWithTracing(base HelloService) HelloService {
	d := HelloServiceWithTracing{
		HelloService: base,
	}

	return d
}

// GetUser implements HelloService
func (_d HelloServiceWithTracing) GetUser(ctx context.Context, req GetUserReq) (pp1 *ent.Pod, err error) {

	var name = "HelloService.GetUser"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"req": req},
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

	return _d.HelloService.GetUser(ctx, req)
}

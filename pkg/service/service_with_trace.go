package service

// Code generated by gowrap. DO NOT EDIT.
// template: ../gowrap/templates/opentracing
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"encoding/json"

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

// Foo implements HelloService
func (_d HelloServiceWithTracing) Foo(ctx context.Context, s string) (rs string, err error) {
	var name = "HelloService.Foo"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"s": s},
				"result": map[string]interface{}{
					"rs":  rs,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()
	return _d.HelloService.Foo(ctx, s)
}

// Say implements HelloService
func (_d HelloServiceWithTracing) Say(ctx context.Context, req SayReq) (res SayRes, err error) {
	var name = "HelloService.Say"
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
	return _d.HelloService.Say(ctx, req)
}

// Say1 implements HelloService
func (_d HelloServiceWithTracing) Say1(ctx context.Context, req SayReq) (res SayRes, err error) {
	var name = "HelloService.Say1"
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
	return _d.HelloService.Say1(ctx, req)
}

// SayHello implements HelloService
func (_d HelloServiceWithTracing) SayHello(ctx context.Context, req SayReq) (res SayRes, err error) {
	var name = "HelloService.SayHello"
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
	return _d.HelloService.SayHello(ctx, req)
}

// SayHello1 implements HelloService
func (_d HelloServiceWithTracing) SayHello1(ctx context.Context, s1 string) (res SayRes, err error) {
	var name = "HelloService.SayHello1"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"s1": s1},
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
	return _d.HelloService.SayHello1(ctx, s1)
}

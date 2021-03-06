package casbin

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
)

// CasbinServiceWithTracing implements CasbinService interface instrumented with opentracing spans
type CasbinServiceWithTracing struct {
	CasbinService
}

// NewCasbinServiceWithTracing returns CasbinServiceWithTracing
func NewCasbinServiceWithTracing(base CasbinService) CasbinService {
	d := CasbinServiceWithTracing{
		CasbinService: base,
	}

	return d
}

// AddRoleAuthorization implements CasbinService
func (_d CasbinServiceWithTracing) AddRoleAuthorization(ctx context.Context, role Role) (err error) {

	var name = "CasbinService.AddRoleAuthorization"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"role": role},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.CasbinService.AddRoleAuthorization(ctx, role)
}

// BindPermission implements CasbinService
func (_d CasbinServiceWithTracing) BindPermission(ctx context.Context, permission Permission) (err error) {

	var name = "CasbinService.BindPermission"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"permission": permission},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.CasbinService.BindPermission(ctx, permission)
}

// CheckPermission implements CasbinService
func (_d CasbinServiceWithTracing) CheckPermission(ctx context.Context, permission CheckPermission) (b1 bool, err error) {

	var name = "CasbinService.CheckPermission"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"permission": permission},
				"result": map[string]interface{}{
					"b1":  b1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.CasbinService.CheckPermission(ctx, permission)
}

// DelRoleAuthorization implements CasbinService
func (_d CasbinServiceWithTracing) DelRoleAuthorization(ctx context.Context, role Role) (err error) {

	var name = "CasbinService.DelRoleAuthorization"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"role": role},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.CasbinService.DelRoleAuthorization(ctx, role)
}

// GetPermissionByUser implements CasbinService
func (_d CasbinServiceWithTracing) GetPermissionByUser(ctx context.Context, permission Permission) (res []Permission) {

	return _d.CasbinService.GetPermissionByUser(ctx, permission)
}

// UnBindPermission implements CasbinService
func (_d CasbinServiceWithTracing) UnBindPermission(ctx context.Context, permission Permission) (err error) {

	var name = "CasbinService.UnBindPermission"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"permission": permission},
				"result": map[string]interface{}{
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.CasbinService.UnBindPermission(ctx, permission)
}

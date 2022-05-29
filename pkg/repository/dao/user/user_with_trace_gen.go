package user

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

// UserServiceWithTracing implements UserService interface instrumented with opentracing spans
type UserServiceWithTracing struct {
	UserService
}

// NewUserServiceWithTracing returns UserServiceWithTracing
func NewUserServiceWithTracing(base UserService) UserService {
	d := UserServiceWithTracing{
		UserService: base,
	}

	return d
}

// ByQueriesAll implements UserService
func (_d UserServiceWithTracing) ByQueriesAll(ctx context.Context, i interface{}) (res []ent.UserBaseGetRes, count int, err error) {

	var name = "UserService.ByQueriesAll"
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

	return _d.UserService.ByQueriesAll(ctx, i)
}

// ByQueriesOne implements UserService
func (_d UserServiceWithTracing) ByQueriesOne(ctx context.Context, i interface{}) (res ent.UserBaseGetRes, err error) {

	var name = "UserService.ByQueriesOne"
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

	return _d.UserService.ByQueriesOne(ctx, i)
}

// Create implements UserService
func (_d UserServiceWithTracing) Create(ctx context.Context, v ent.UserBaseCreateReq) (res *ent.User, err error) {

	var name = "UserService.Create"
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

	return _d.UserService.Create(ctx, v)
}

// CreateMany implements UserService
func (_d UserServiceWithTracing) CreateMany(ctx context.Context, vs []ent.UserBaseCreateReq) (u1 ent.Users, err error) {

	var name = "UserService.CreateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"u1":  u1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.UserService.CreateMany(ctx, vs)
}

// DeleteById implements UserService
func (_d UserServiceWithTracing) DeleteById(ctx context.Context, id int) (err error) {

	var name = "UserService.DeleteById"
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

	return _d.UserService.DeleteById(ctx, id)
}

// DeleteMany implements UserService
func (_d UserServiceWithTracing) DeleteMany(ctx context.Context, ids []int) (err error) {

	var name = "UserService.DeleteMany"
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

	return _d.UserService.DeleteMany(ctx, ids)
}

// GetById implements UserService
func (_d UserServiceWithTracing) GetById(ctx context.Context, id int) (res ent.UserBaseGetRes, err error) {

	var name = "UserService.GetById"
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

	return _d.UserService.GetById(ctx, id)
}

// RawByQueriesAll implements UserService
func (_d UserServiceWithTracing) RawByQueriesAll(ctx context.Context, i interface{}) (res ent.Users, count int, err error) {

	var name = "UserService.RawByQueriesAll"
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

	return _d.UserService.RawByQueriesAll(ctx, i)
}

// RawByQueriesOne implements UserService
func (_d UserServiceWithTracing) RawByQueriesOne(ctx context.Context, i interface{}) (res *ent.User, err error) {

	var name = "UserService.RawByQueriesOne"
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

	return _d.UserService.RawByQueriesOne(ctx, i)
}

// RawCreate implements UserService
func (_d UserServiceWithTracing) RawCreate(ctx context.Context, v *ent.User) (res *ent.User, err error) {

	var name = "UserService.RawCreate"
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

	return _d.UserService.RawCreate(ctx, v)
}

// RawCreateMany implements UserService
func (_d UserServiceWithTracing) RawCreateMany(ctx context.Context, vs ent.Users) (u1 ent.Users, err error) {

	var name = "UserService.RawCreateMany"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"vs": vs},
				"result": map[string]interface{}{
					"u1":  u1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.UserService.RawCreateMany(ctx, vs)
}

// RawGetById implements UserService
func (_d UserServiceWithTracing) RawGetById(ctx context.Context, id int) (res *ent.User, err error) {

	var name = "UserService.RawGetById"
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

	return _d.UserService.RawGetById(ctx, id)
}

// RawUpdateById implements UserService
func (_d UserServiceWithTracing) RawUpdateById(ctx context.Context, id int, v *ent.User) (up1 *ent.User, err error) {

	var name = "UserService.RawUpdateById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"up1": up1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.UserService.RawUpdateById(ctx, id, v)
}

// RawUpdateMany implements UserService
func (_d UserServiceWithTracing) RawUpdateMany(ctx context.Context, vs ent.Users) (err error) {

	var name = "UserService.RawUpdateMany"
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

	return _d.UserService.RawUpdateMany(ctx, vs)
}

// UpdateById implements UserService
func (_d UserServiceWithTracing) UpdateById(ctx context.Context, id int, v ent.UserBaseUpdateReq) (up1 *ent.User, err error) {

	var name = "UserService.UpdateById"
	_, span := otel.Tracer(name).Start(ctx, name)
	defer func() {
		if err != nil {
			l := map[string]interface{}{
				"params": map[string]interface{}{
					"id": id,
					"v":  v},
				"result": map[string]interface{}{
					"up1": up1,
					"err": err},
			}
			s, _ := json.Marshal(l)
			span.AddEvent(semconv.ExceptionEventName, trace.WithAttributes(semconv.ExceptionTypeKey.String("context"), semconv.ExceptionMessageKey.String(string(s))))
			span.SetStatus(codes.Error, err.Error())
		}
		span.End()
	}()

	return _d.UserService.UpdateById(ctx, id, v)
}

// UpdateMany implements UserService
func (_d UserServiceWithTracing) UpdateMany(ctx context.Context, vs []ent.UserBaseUpdateReq) (err error) {

	var name = "UserService.UpdateMany"
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

	return _d.UserService.UpdateMany(ctx, vs)
}

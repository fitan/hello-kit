package user

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"hello/pkg/ent"
)

// UserServiceWithLog implements UserService that is instrumented with logging
type UserServiceWithLog struct {
	_log  *zap.SugaredLogger
	_base UserService
}

// NewUserServiceWithLog instruments an implementation of the UserService with simple logging
func NewUserServiceWithLog(base UserService, log *zap.SugaredLogger) UserServiceWithLog {
	return UserServiceWithLog{
		_base: base,
		_log:  log,
	}
}

// ByQueries implements UserService
func (_d UserServiceWithLog) ByQueries(ctx context.Context, req ent.UserRestByQueriesReq) (res ent.UserRestByQueriesRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling ByQueries", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.ByQueries(ctx, req)
}

// Create implements UserService
func (_d UserServiceWithLog) Create(ctx context.Context, req ent.UserRestCreateReq) (res *ent.User, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling Create", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.Create(ctx, req)
}

// CreateMany implements UserService
func (_d UserServiceWithLog) CreateMany(ctx context.Context, req ent.UserRestCreateManyReq) (res ent.Users, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling CreateMany", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.CreateMany(ctx, req)
}

// CreatePodsSliceByUserId implements UserService
func (_d UserServiceWithLog) CreatePodsSliceByUserId(ctx context.Context, req ent.UserRestCreatePodsSliceByUserIdReq) (res *ent.User, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling CreatePodsSliceByUserId", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.CreatePodsSliceByUserId(ctx, req)
}

// DeleteById implements UserService
func (_d UserServiceWithLog) DeleteById(ctx context.Context, req ent.UserRestDeleteByIdReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling DeleteById", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"success": success,
			"err":     err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}

	}()
	return _d._base.DeleteById(ctx, req)
}

// DeleteMany implements UserService
func (_d UserServiceWithLog) DeleteMany(ctx context.Context, req ent.UserRestDeleteManyReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling DeleteMany", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"success": success,
			"err":     err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}

	}()
	return _d._base.DeleteMany(ctx, req)
}

// GetById implements UserService
func (_d UserServiceWithLog) GetById(ctx context.Context, req ent.UserRestGetByIdReq) (res ent.UserBaseGetRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling GetById", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.GetById(ctx, req)
}

// GetPodsSliceByUserId implements UserService
func (_d UserServiceWithLog) GetPodsSliceByUserId(ctx context.Context, req ent.UserRestGetPodsSliceByUserIdReq) (res ent.UserRestGetPodsSliceByUserIdRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling GetPodsSliceByUserId", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.GetPodsSliceByUserId(ctx, req)
}

// UpdateById implements UserService
func (_d UserServiceWithLog) UpdateById(ctx context.Context, req ent.UserRestUpdateByIdReq) (res *ent.User, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling UpdateById", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"res": res,
			"err": err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

	}()
	return _d._base.UpdateById(ctx, req)
}

// UpdateMany implements UserService
func (_d UserServiceWithLog) UpdateMany(ctx context.Context, req ent.UserRestUpdateManyReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("UserServiceWithLog calling UpdateMany", "params", map[string]interface{}{
			"req": req}, "results", map[string]interface{}{
			"success": success,
			"err":     err})
		if err != nil {
			_log.Errorw("with_log", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}

	}()
	return _d._base.UpdateMany(ctx, req)
}

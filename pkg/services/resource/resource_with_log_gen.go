package resource

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"fmt"

	ginkHttp "github.com/fitan/gink/transport/http"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"hello/pkg/ent"
)

// ResourceServiceWithLog implements ResourceService that is instrumented with logging
type ResourceServiceWithLog struct {
	_log  *zap.SugaredLogger
	_base ResourceService
}

// NewResourceServiceWithLog instruments an implementation of the ResourceService with simple logging
func NewResourceServiceWithLog(base ResourceService, log *zap.SugaredLogger) ResourceServiceWithLog {
	return ResourceServiceWithLog{
		_base: base,
		_log:  log,
	}
}

// ResourceRestByQueriesAll implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestByQueriesAll(ctx context.Context, req ent.ResourceRestByQueriesAllReq) (res ent.ResourceRestByQueriesAllRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestByQueriesAll", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestByQueriesAll", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestByQueriesAll", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestByQueriesAll", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestByQueriesAll(ctx, req)
}

// ResourceRestCreate implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestCreate(ctx context.Context, req ent.ResourceRestCreateReq) (res *ent.Resource, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestCreate", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestCreate", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestCreate", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestCreate", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestCreate(ctx, req)
}

// ResourceRestCreateMany implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestCreateMany(ctx context.Context, req ent.ResourceRestCreateManyReq) (res ent.Resources, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestCreateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestCreateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestCreateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestCreateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestCreateMany(ctx, req)
}

// ResourceRestDeleteById implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestDeleteById(ctx context.Context, req ent.ResourceRestDeleteByIdReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestDeleteById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestDeleteById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestDeleteById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestDeleteById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestDeleteById(ctx, req)
}

// ResourceRestDeleteMany implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestDeleteMany(ctx context.Context, req ent.ResourceRestDeleteManyReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestDeleteMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestDeleteMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestDeleteMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestDeleteMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestDeleteMany(ctx, req)
}

// ResourceRestGetById implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestGetById(ctx context.Context, req ent.ResourceRestGetByIdReq) (res ent.ResourceBaseGetRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestGetById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestGetById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestGetById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestGetById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestGetById(ctx, req)
}

// ResourceRestUpdateById implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestUpdateById(ctx context.Context, req ent.ResourceRestUpdateByIdReq) (res *ent.Resource, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestUpdateById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestUpdateById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestUpdateById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestUpdateById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestUpdateById(ctx, req)
}

// ResourceRestUpdateMany implements ResourceService
func (_d ResourceServiceWithLog) ResourceRestUpdateMany(ctx context.Context, req ent.ResourceRestUpdateManyReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ResourceRestUpdateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ResourceRestUpdateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ResourceRestUpdateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ResourceRestUpdateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ResourceRestUpdateMany(ctx, req)
}

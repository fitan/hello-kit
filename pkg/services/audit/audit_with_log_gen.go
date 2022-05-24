package audit

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"

	ginkHttp "github.com/fitan/gink/transport/http"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"

	"hello/pkg/ent"
)

// AuditServiceWithLog implements AuditService that is instrumented with logging
type AuditServiceWithLog struct {
	_log  *zap.SugaredLogger
	_base AuditService
}

// NewAuditServiceWithLog instruments an implementation of the AuditService with simple logging
func NewAuditServiceWithLog(base AuditService, log *zap.SugaredLogger) AuditServiceWithLog {
	return AuditServiceWithLog{
		_base: base,
		_log:  log,
	}
}

// AuditRestByQueriesAll implements AuditService
func (_d AuditServiceWithLog) AuditRestByQueriesAll(ctx context.Context, req ent.AuditRestByQueriesAllReq) (res ent.AuditRestByQueriesAllRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestByQueriesAll", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestByQueriesAll", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestByQueriesAll", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestByQueriesAll", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}
	}()
	return _d._base.AuditRestByQueriesAll(ctx, req)
}

// AuditRestCreate implements AuditService
func (_d AuditServiceWithLog) AuditRestCreate(ctx context.Context, req ent.AuditRestCreateReq) (res *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestCreate", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestCreate", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestCreate", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestCreate", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}
	}()
	return _d._base.AuditRestCreate(ctx, req)
}

// AuditRestCreateMany implements AuditService
func (_d AuditServiceWithLog) AuditRestCreateMany(ctx context.Context, req ent.AuditRestCreateManyReq) (res ent.Audits, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestCreateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestCreateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestCreateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestCreateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}
	}()
	return _d._base.AuditRestCreateMany(ctx, req)
}

// AuditRestDeleteById implements AuditService
func (_d AuditServiceWithLog) AuditRestDeleteById(ctx context.Context, req ent.AuditRestDeleteByIdReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestDeleteById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestDeleteById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestDeleteById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestDeleteById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}
	}()
	return _d._base.AuditRestDeleteById(ctx, req)
}

// AuditRestDeleteMany implements AuditService
func (_d AuditServiceWithLog) AuditRestDeleteMany(ctx context.Context, req ent.AuditRestDeleteManyReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestDeleteMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestDeleteMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestDeleteMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestDeleteMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}
	}()
	return _d._base.AuditRestDeleteMany(ctx, req)
}

// AuditRestGetById implements AuditService
func (_d AuditServiceWithLog) AuditRestGetById(ctx context.Context, req ent.AuditRestGetByIdReq) (res ent.AuditBaseGetRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestGetById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestGetById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestGetById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestGetById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}
	}()
	return _d._base.AuditRestGetById(ctx, req)
}

// AuditRestUpdateById implements AuditService
func (_d AuditServiceWithLog) AuditRestUpdateById(ctx context.Context, req ent.AuditRestUpdateByIdReq) (res *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestUpdateById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestUpdateById", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"res": res,
					"err": err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestUpdateById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestUpdateById", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"res": res,
				"err": err})
		}
	}()
	return _d._base.AuditRestUpdateById(ctx, req)
}

// AuditRestUpdateMany implements AuditService
func (_d AuditServiceWithLog) AuditRestUpdateMany(ctx context.Context, req ent.AuditRestUpdateManyReq) (success bool, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling AuditRestUpdateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     err})
			}

			if err != nil {
				_log.Errorw("with_log calling AuditRestUpdateMany", "params", map[string]interface{}{
					"req": req}, "results", map[string]interface{}{
					"success": success,
					"err":     err})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling AuditRestUpdateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling AuditRestUpdateMany", "params", map[string]interface{}{
				"req": req}, "results", map[string]interface{}{
				"success": success,
				"err":     err})
		}
	}()
	return _d._base.AuditRestUpdateMany(ctx, req)
}

package audit

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

// ByQueriesAll implements AuditService
func (_d AuditServiceWithLog) ByQueriesAll(ctx context.Context, i interface{}) (res []ent.AuditBaseGetRes, count int, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ByQueriesAll", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ByQueriesAll", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ByQueriesAll", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res":   res,
				"count": count,
				"err":   fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ByQueriesAll", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res":   res,
				"count": count,
				"err":   fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ByQueriesAll(ctx, i)
}

// ByQueriesOne implements AuditService
func (_d AuditServiceWithLog) ByQueriesOne(ctx context.Context, i interface{}) (res ent.AuditBaseGetRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling ByQueriesOne", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling ByQueriesOne", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling ByQueriesOne", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling ByQueriesOne", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.ByQueriesOne(ctx, i)
}

// Create implements AuditService
func (_d AuditServiceWithLog) Create(ctx context.Context, v ent.AuditBaseCreateReq) (res *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling Create", "params", map[string]interface{}{
					"v": v}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling Create", "params", map[string]interface{}{
					"v": v}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling Create", "params", map[string]interface{}{
				"v": v}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling Create", "params", map[string]interface{}{
				"v": v}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.Create(ctx, v)
}

// CreateMany implements AuditService
func (_d AuditServiceWithLog) CreateMany(ctx context.Context, vs []ent.AuditBaseCreateReq) (a1 ent.Audits, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling CreateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"a1":  a1,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling CreateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"a1":  a1,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling CreateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"a1":  a1,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling CreateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"a1":  a1,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.CreateMany(ctx, vs)
}

// DeleteById implements AuditService
func (_d AuditServiceWithLog) DeleteById(ctx context.Context, id int) (err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling DeleteById", "params", map[string]interface{}{
					"id": id}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling DeleteById", "params", map[string]interface{}{
					"id": id}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling DeleteById", "params", map[string]interface{}{
				"id": id}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling DeleteById", "params", map[string]interface{}{
				"id": id}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.DeleteById(ctx, id)
}

// DeleteMany implements AuditService
func (_d AuditServiceWithLog) DeleteMany(ctx context.Context, ids []int) (err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling DeleteMany", "params", map[string]interface{}{
					"ids": ids}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling DeleteMany", "params", map[string]interface{}{
					"ids": ids}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling DeleteMany", "params", map[string]interface{}{
				"ids": ids}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling DeleteMany", "params", map[string]interface{}{
				"ids": ids}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.DeleteMany(ctx, ids)
}

// GetById implements AuditService
func (_d AuditServiceWithLog) GetById(ctx context.Context, id int) (res ent.AuditBaseGetRes, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling GetById", "params", map[string]interface{}{
					"id": id}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling GetById", "params", map[string]interface{}{
					"id": id}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling GetById", "params", map[string]interface{}{
				"id": id}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling GetById", "params", map[string]interface{}{
				"id": id}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.GetById(ctx, id)
}

// RawByQueriesAll implements AuditService
func (_d AuditServiceWithLog) RawByQueriesAll(ctx context.Context, i interface{}) (res ent.Audits, count int, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawByQueriesAll", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawByQueriesAll", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res":   res,
					"count": count,
					"err":   fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawByQueriesAll", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res":   res,
				"count": count,
				"err":   fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawByQueriesAll", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res":   res,
				"count": count,
				"err":   fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawByQueriesAll(ctx, i)
}

// RawByQueriesOne implements AuditService
func (_d AuditServiceWithLog) RawByQueriesOne(ctx context.Context, i interface{}) (res *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawByQueriesOne", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawByQueriesOne", "params", map[string]interface{}{
					"i": i}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawByQueriesOne", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawByQueriesOne", "params", map[string]interface{}{
				"i": i}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawByQueriesOne(ctx, i)
}

// RawCreate implements AuditService
func (_d AuditServiceWithLog) RawCreate(ctx context.Context, v *ent.Audit) (res *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawCreate", "params", map[string]interface{}{
					"v": v}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawCreate", "params", map[string]interface{}{
					"v": v}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawCreate", "params", map[string]interface{}{
				"v": v}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawCreate", "params", map[string]interface{}{
				"v": v}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawCreate(ctx, v)
}

// RawCreateMany implements AuditService
func (_d AuditServiceWithLog) RawCreateMany(ctx context.Context, vs ent.Audits) (a1 ent.Audits, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawCreateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"a1":  a1,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawCreateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"a1":  a1,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawCreateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"a1":  a1,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawCreateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"a1":  a1,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawCreateMany(ctx, vs)
}

// RawGetById implements AuditService
func (_d AuditServiceWithLog) RawGetById(ctx context.Context, id int) (res *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawGetById", "params", map[string]interface{}{
					"id": id}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawGetById", "params", map[string]interface{}{
					"id": id}, "results", map[string]interface{}{
					"res": res,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawGetById", "params", map[string]interface{}{
				"id": id}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawGetById", "params", map[string]interface{}{
				"id": id}, "results", map[string]interface{}{
				"res": res,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawGetById(ctx, id)
}

// RawUpdateById implements AuditService
func (_d AuditServiceWithLog) RawUpdateById(ctx context.Context, id int, v *ent.Audit) (ap1 *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawUpdateById", "params", map[string]interface{}{
					"id": id,
					"v":  v}, "results", map[string]interface{}{
					"ap1": ap1,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawUpdateById", "params", map[string]interface{}{
					"id": id,
					"v":  v}, "results", map[string]interface{}{
					"ap1": ap1,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawUpdateById", "params", map[string]interface{}{
				"id": id,
				"v":  v}, "results", map[string]interface{}{
				"ap1": ap1,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawUpdateById", "params", map[string]interface{}{
				"id": id,
				"v":  v}, "results", map[string]interface{}{
				"ap1": ap1,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawUpdateById(ctx, id, v)
}

// RawUpdateMany implements AuditService
func (_d AuditServiceWithLog) RawUpdateMany(ctx context.Context, vs ent.Audits) (err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling RawUpdateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling RawUpdateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling RawUpdateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling RawUpdateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.RawUpdateMany(ctx, vs)
}

// UpdateById implements AuditService
func (_d AuditServiceWithLog) UpdateById(ctx context.Context, id int, v ent.AuditBaseUpdateReq) (ap1 *ent.Audit, err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling UpdateById", "params", map[string]interface{}{
					"id": id,
					"v":  v}, "results", map[string]interface{}{
					"ap1": ap1,
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling UpdateById", "params", map[string]interface{}{
					"id": id,
					"v":  v}, "results", map[string]interface{}{
					"ap1": ap1,
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling UpdateById", "params", map[string]interface{}{
				"id": id,
				"v":  v}, "results", map[string]interface{}{
				"ap1": ap1,
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling UpdateById", "params", map[string]interface{}{
				"id": id,
				"v":  v}, "results", map[string]interface{}{
				"ap1": ap1,
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.UpdateById(ctx, id, v)
}

// UpdateMany implements AuditService
func (_d AuditServiceWithLog) UpdateMany(ctx context.Context, vs []ent.AuditBaseUpdateReq) (err error) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	debug, _ := ctx.Value(ginkHttp.ContextKeyRequestDebug).(bool)

	defer func() {
		if debug {
			if err == nil {
				_log.Infow("with_log calling UpdateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

			if err != nil {
				_log.Errorw("with_log calling UpdateMany", "params", map[string]interface{}{
					"vs": vs}, "results", map[string]interface{}{
					"err": fmt.Sprintf("%v", err)})
			}

		}
		if !debug && err == nil {
			_log.Debugw("with_log calling UpdateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}

		if err != nil && !debug {
			_log.Errorw("with_log calling UpdateMany", "params", map[string]interface{}{
				"vs": vs}, "results", map[string]interface{}{
				"err": fmt.Sprintf("%v", err)})
		}
	}()
	return _d._base.UpdateMany(ctx, vs)
}

package taobao

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

// TaobaoServiceWithLog implements TaobaoService that is instrumented with logging
type TaobaoServiceWithLog struct {
	_log  *zap.SugaredLogger
	_base TaobaoService
}

// NewTaobaoServiceWithLog instruments an implementation of the TaobaoService with simple logging
func NewTaobaoServiceWithLog(base TaobaoService, log *zap.SugaredLogger) TaobaoServiceWithLog {
	return TaobaoServiceWithLog{
		_base: base,
		_log:  log,
	}
}

// GetRoot implements TaobaoService
func (_d TaobaoServiceWithLog) GetRoot(ctx context.Context) (s1 string) {

	_log := _d._log.With(zap.String("traceId", trace.SpanFromContext(ctx).SpanContext().TraceID().String()))

	defer func() {
		_log.Debugw("TaobaoServiceWithLog calling GetRoot", "params", map[string]interface{}{}, "results", map[string]interface{}{
			"s1": s1})

	}()
	return _d._base.GetRoot(ctx)
}

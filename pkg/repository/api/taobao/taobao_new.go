package taobao

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Middleware func(TaobaoService) TaobaoService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(TaobaoService TaobaoService) TaobaoService {
		return NewTaobaoServiceWithPrometheus(TaobaoService)
	})
	mw = append(mw, func(TaobaoService TaobaoService) TaobaoService {
		return NewTaobaoServiceWithLog(TaobaoService, logger)
	})
	mw = append(mw, func(TaobaoService TaobaoService) TaobaoService {
		return NewTaobaoServiceWithTracing(TaobaoService)
	})
	return
}

// New returns a TaobaoService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) TaobaoService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

var TaobaoServiceSet = wire.NewSet(NewBasicService, NewService, NewServiceMiddleware)

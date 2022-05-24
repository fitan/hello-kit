package service

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Middleware func(ServiceService) ServiceService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(ServiceService ServiceService) ServiceService {
		return NewServiceServiceWithPrometheus(ServiceService)
	})
	mw = append(mw, func(ServiceService ServiceService) ServiceService {
		return NewServiceServiceWithLog(ServiceService, logger)
	})
	mw = append(mw, func(ServiceService ServiceService) ServiceService {
		return NewServiceServiceWithTracing(ServiceService)
	})
	return
}

// New returns a ServiceService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) ServiceService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

var ServiceServiceSet = wire.NewSet(NewBasicService, NewService, NewServiceMiddleware)

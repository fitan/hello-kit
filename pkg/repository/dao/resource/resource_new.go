package resource

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Middleware func(ResourceService) ResourceService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(ResourceService ResourceService) ResourceService {
		return NewResourceServiceWithPrometheus(ResourceService)
	})
	mw = append(mw, func(ResourceService ResourceService) ResourceService {
		return NewResourceServiceWithLog(ResourceService, logger)
	})
	mw = append(mw, func(ResourceService ResourceService) ResourceService {
		return NewResourceServiceWithTracing(ResourceService)
	})
	return
}

// New returns a ResourceService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) ResourceService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

var ResourceServiceSet = wire.NewSet(NewBasicService, NewService, NewServiceMiddleware)

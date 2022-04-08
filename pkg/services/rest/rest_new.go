package rest

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

type Middleware func(RestService) RestService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(RestService RestService) RestService {
		return NewRestServiceWithPrometheus(RestService)
	})
	mw = append(mw, func(RestService RestService) RestService {
		return NewRestServiceWithLog(RestService, logger)
	})
	mw = append(mw, func(RestService RestService) RestService {
		return NewRestServiceWithTracing(RestService)
	})
	return
}

// New returns a RestService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) RestService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

var RestServiceSet = wire.NewSet(NewBasicService, NewService, NewServiceMiddleware)

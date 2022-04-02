package casbin

import "go.uber.org/zap"

type Middleware func(CasbinService) CasbinService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(CasbinService CasbinService) CasbinService {
		return NewCasbinServiceWithPrometheus(CasbinService)
	})
	mw = append(mw, func(CasbinService CasbinService) CasbinService {
		return NewCasbinServiceWithLog(CasbinService, logger)
	})
	mw = append(mw, func(CasbinService CasbinService) CasbinService {
		return NewCasbinServiceWithTracing(CasbinService)
	})
	return
}

// New returns a CasbinService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) CasbinService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

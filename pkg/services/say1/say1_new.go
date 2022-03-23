package say1

import "go.uber.org/zap"

type Middleware func(Say1Service) Say1Service

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(Say1Service Say1Service) Say1Service {
		return NewSay1ServiceWithPrometheus(Say1Service, "pkg.services.Say1")
	})
	mw = append(mw, func(Say1Service Say1Service) Say1Service {
		return NewSay1ServiceWithLog(Say1Service, logger)
	})
	mw = append(mw, func(Say1Service Say1Service) Say1Service {
		return NewSay1ServiceWithTracing(Say1Service)
	})
	return
}

// New returns a Say1Service with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) Say1Service {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

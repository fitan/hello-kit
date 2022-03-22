package say

import "go.uber.org/zap"

type Middleware func(SayService) SayService

func NewServiceMiddleware(logger *zap.SugaredLogger, appName string) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(SayService SayService) SayService {
		return NewSayServiceWithPrometheus(SayService, appName)
	})
	mw = append(mw, func(SayService SayService) SayService {
		return NewSayServiceWithLog(SayService, logger)
	})
	mw = append(mw, func(SayService SayService) SayService {
		return NewSayServiceWithTracing(SayService)
	})
	return
}

// New returns a SayService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) SayService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

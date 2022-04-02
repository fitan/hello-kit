package hello

import "go.uber.org/zap"

type Middleware func(HelloService) HelloService

func NewServiceMiddleware(logger *zap.SugaredLogger, appName string) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(HelloService HelloService) HelloService {
		return NewHelloServiceWithPrometheus(HelloService, appName)
	})
	mw = append(mw, func(HelloService HelloService) HelloService {
		return NewHelloServiceWithLog(HelloService, logger)
	})
	mw = append(mw, func(HelloService HelloService) HelloService {
		return NewHelloServiceWithTracing(HelloService)
	})
	return
}

// New returns a HelloService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) HelloService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

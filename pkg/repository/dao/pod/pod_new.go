package pod

import "go.uber.org/zap"

type Middleware func(PodService) PodService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(PodService PodService) PodService {
		return NewPodServiceWithPrometheus(PodService)
	})
	mw = append(mw, func(PodService PodService) PodService {
		return NewPodServiceWithLog(PodService, logger)
	})
	mw = append(mw, func(PodService PodService) PodService {
		return NewPodServiceWithTracing(PodService)
	})
	return
}

// New returns a PodService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) PodService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

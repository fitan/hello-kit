package project

import "go.uber.org/zap"

type Middleware func(ProjectService) ProjectService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(ProjectService ProjectService) ProjectService {
		return NewProjectServiceWithPrometheus(ProjectService, "repo.dao")
	})
	mw = append(mw, func(ProjectService ProjectService) ProjectService {
		return NewProjectServiceWithLog(ProjectService, logger)
	})
	mw = append(mw, func(ProjectService ProjectService) ProjectService {
		return NewProjectServiceWithTracing(ProjectService)
	})
	return
}

// New returns a ProjectService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) ProjectService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

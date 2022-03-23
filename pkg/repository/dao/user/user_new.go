package user

import "go.uber.org/zap"

type Middleware func(UserService) UserService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(UserService UserService) UserService {
		return NewUserServiceWithPrometheus(UserService, "repo.dao")
	})
	mw = append(mw, func(UserService UserService) UserService {
		return NewUserServiceWithLog(UserService, logger)
	})
	mw = append(mw, func(UserService UserService) UserService {
		return NewUserServiceWithTracing(UserService)
	})
	return
}

// New returns a UserService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) UserService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

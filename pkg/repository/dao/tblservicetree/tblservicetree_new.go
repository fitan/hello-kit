package tblservicetree

import "go.uber.org/zap"

type Middleware func(TblservicetreeService) TblservicetreeService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(TblservicetreeService TblservicetreeService) TblservicetreeService {
		return NewTblservicetreeServiceWithPrometheus(TblservicetreeService, "repo.dao")
	})
	mw = append(mw, func(TblservicetreeService TblservicetreeService) TblservicetreeService {
		return NewTblservicetreeServiceWithLog(TblservicetreeService, logger)
	})
	mw = append(mw, func(TblservicetreeService TblservicetreeService) TblservicetreeService {
		return NewTblservicetreeServiceWithTracing(TblservicetreeService)
	})
	return
}

// New returns a TblservicetreeService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) TblservicetreeService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

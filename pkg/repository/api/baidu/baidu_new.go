package baidu

import (
	"go.uber.org/zap"
)

type Middleware func(BaiduService) BaiduService

func NewServiceMiddleware(logger *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(BaiduService BaiduService) BaiduService {
		return NewBaiduServiceWithPrometheus(BaiduService)
	})
	mw = append(mw, func(BaiduService BaiduService) BaiduService {
		return NewBaiduServiceWithLog(BaiduService, logger)
	})
	mw = append(mw, func(BaiduService BaiduService) BaiduService {
		return NewBaiduServiceWithTracing(BaiduService)
	})
	return
}

// New returns a BaiduService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) BaiduService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

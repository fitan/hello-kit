package baidu

import "go.uber.org/zap"

func NewBaiduApi(baseHttp Base, middleware []Middleware) BaiduApi {
	for _, m := range middleware {
		baseHttp = m(baseHttp)
	}

	return baseHttp
}

type Middleware func(BaiduApi) BaiduApi

func NewBaiduApiMiddleware(log *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}

	mw = append(mw, func(http BaiduApi) BaiduApi {
		return NewBaiduApiWithPrometheus(http, "repo.baidu")
	})
	mw = append(mw, func(http BaiduApi) BaiduApi {
		return NewBaiduApiWithLog(http, log)
	})
	mw = append(mw, func(http BaiduApi) BaiduApi {
		return NewBaiduApiWithTracing(http)
	})

	return
}

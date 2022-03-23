package taobao

import "go.uber.org/zap"

func NewTaoApi(baseHttp Base, middleware []Middleware) TaobaoApi {
	for _, m := range middleware {
		baseHttp = m(baseHttp)
	}

	return baseHttp
}

type Middleware func(TaobaoApi) TaobaoApi

func NewTaobaoApiMiddleware(log *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}

	mw = append(mw, func(http TaobaoApi) TaobaoApi {
		return NewTaobaoApiWithPrometheus(http, "repo.api")
	})
	mw = append(mw, func(http TaobaoApi) TaobaoApi {
		return NewTaobaoApiWithLog(http, log)
	})
	mw = append(mw, func(http TaobaoApi) TaobaoApi {
		return NewTaobaoApiWithTracing(http)
	})

	return
}

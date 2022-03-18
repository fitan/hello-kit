package baidu

import "go.uber.org/zap"

func NewHttp(baseHttp BaseHttp, middleware []Middleware) Http {
	for _, m := range middleware {
		baseHttp = m(baseHttp)
	}

	return baseHttp
}

type Middleware func(Http) Http

func NewHttpMiddleware(log *zap.SugaredLogger) (mw []Middleware) {
	mw = []Middleware{}

	mw = append(mw, func(http Http) Http {
		return NewHttpWithPrometheus(http, "repo.baidu")
	})
	mw = append(mw, func(http Http) Http {
		return NewHttpWithLog(http, log)
	})
	mw = append(mw, func(http Http) Http {
		return NewHttpWithTracing(http)
	})

	return
}

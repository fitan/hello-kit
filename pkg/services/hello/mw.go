package hello

import (
	"github.com/go-kit/log"
)

type Middleware func(HelloService) HelloService

func getServiceMiddleware(logger log.Logger) (mw []Middleware) {
	mw = []Middleware{}
	// Append your middleware here
	mw = append(mw, func(helloService HelloService) HelloService {
		return NewHelloServiceWithPrometheus(helloService, "hello-kit")
	})
	mw = append(mw, func(helloService HelloService) HelloService {
		return NewHelloServiceWithLog(helloService, logger)
	})
	mw = append(mw, func(helloService HelloService) HelloService {
		return NewHelloServiceWithTracing(helloService)
	})

	return
}

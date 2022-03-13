package services

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/metrics/prometheus"
	http3 "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/go-kit/kit/otelkit"
	"hello/pkg/mid"
	"hello/pkg/services/hello"
	http2 "net/http"
)

func helloServiceMiddleware(logger log.Logger) (mw []hello.Middleware) {
	mw = []hello.Middleware{}
	// Append your middleware here
	mw = append(mw, func(helloService hello.HelloService) hello.HelloService {
		return hello.NewHelloServiceWithPrometheus(helloService, "hello-kit")
	})
	mw = append(mw, func(helloService hello.HelloService) hello.HelloService {
		return hello.NewHelloServiceWithLog(helloService, logger)
	})
	mw = append(mw, func(helloService hello.HelloService) hello.HelloService {
		return hello.NewHelloServiceWithTracing(helloService)
	})

	return
}

func helloEndpointMiddleware(logger log.Logger) (mw map[string][]endpoint1.Middleware) {
	mw = map[string][]endpoint1.Middleware{}
	// Add you endpoint middleware here
	otelkitEMW := func(n string) endpoint1.Middleware {
		return otelkit.EndpointMiddleware(otelkit.WithOperation(n))
	}
	logEMW := func(n string) endpoint1.Middleware {
		return mid.LoggingMiddleware(logger)
	}
	instrumentingEMW := func(n string) endpoint1.Middleware {
		return mid.InstrumentingMiddleware(prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "hello",
			Subsystem: n,
			Name:      "request_duration_seconds",
			Help:      "Total time spent serving requests.",
		}, []string{"success"}))
	}
	hello.AddEndpointMiddlewareToAllMethods(mw, otelkitEMW)
	hello.AddEndpointMiddlewareToAllMethods(mw, logEMW)
	hello.AddEndpointMiddlewareToAllMethods(mw, instrumentingEMW)

	return
}

func InitHttpHandler(m *http2.ServeMux, log log.Logger) {

	helloSvc := hello.New(helloServiceMiddleware(log))
	helloEps := hello.NewEndpoints(helloSvc, helloEndpointMiddleware(log))
	pc := http3.ServerBefore(http3.PopulateRequestContext)
	helloPptions := make(map[string][]http3.ServerOption)
	hello.AddHttpOptionToAllMethods(helloPptions, pc)
	hello.NewHTTPHandler(m, helloEps, helloPptions)
}

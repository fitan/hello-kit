package hello

import (
	"context"
	"hello/pkg/services/hello/types"
)

// HelloService describes the service.
//go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/prometheus.tmpl -o service_with_prometheus.go
//go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/log.tmpl -o service_with_log.go
//go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/opentracing.tmpl -o service_with_trace.go
//go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/endpoint.tmpl -o endpoint.go
//go:generate gowrap gen -g -ps -p ./ -i HelloService -t ../../gowrap/templates/http-gin.tmpl -o http.go
type HelloService interface {
	// @http-gin /foo/:id POST
	Foo(ctx context.Context, s types.SayReq) (rs string, err error)
	// @http-gin /say/:id PUT
	Say(ctx context.Context, req types.SayReq) (res types.SayRes, err error)
	// @http-gin /say1/:id DELETE
	Say1(ctx context.Context, req types.SayReq) (res types.SayRes, err error)
	// @http-gin /sayhello/:id GET
	SayHello(ctx context.Context, req types.SayReq) (res types.SayRes, err error)
	// @http-gin /sayhello1/:id GET
	SayHello1(ctx context.Context, s1 types.SayReq) (res types.SayRes, err error)
}

type basicHelloService struct{}

func (b *basicHelloService) Foo(ctx context.Context, s types.SayReq) (rs string, err error) {
	return
}

func (b *basicHelloService) Say(ctx context.Context, req types.SayReq) (res types.SayRes, err error) {
	return
}

func (b *basicHelloService) Say1(ctx context.Context, req types.SayReq) (res types.SayRes, err error) {
	return
}

func (b *basicHelloService) SayHello(ctx context.Context, req types.SayReq) (res types.SayRes, err error) {
	return
}

func (b *basicHelloService) SayHello1(ctx context.Context, s1 types.SayReq) (res types.SayRes, err error) {
	return
}

type Middleware func(HelloService) HelloService

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

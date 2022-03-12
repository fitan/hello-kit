package service

import (
	"context"
	"errors"
)

// HelloService describes the service.
//go:generate gowrap gen -p ./ -i HelloService -t ../gowrap/templates/prometheus -o service_with_prometheus.go
//go:generate gowrap gen -p ./ -i HelloService -t ../gowrap/templates/log -o service_with_log.go
//go:generate gowrap gen -p ./ -i HelloService -t ../gowrap/templates/opentracing -o service_with_trace.go
type HelloService interface {
	Foo(ctx context.Context, s string) (rs string, err error)
	Say(ctx context.Context, req SayReq) (res SayRes, err error)
	Say1(ctx context.Context, req SayReq) (res SayRes, err error)
	SayHello(ctx context.Context, req SayReq) (res SayRes, err error)
	SayHello1(ctx context.Context, s1 string, s2 string) (res SayRes, err error)
}

type basicHelloService struct{}

func (b *basicHelloService) Foo(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Foo
	//r := ctx.Value(http.ContextKeyRequestHost).(string)
	if s != "" {
		return "", errors.New("this is erorr")
	}
	return s, err
}

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

func (b *basicHelloService) Say(ctx context.Context, req SayReq) (res SayRes, err error) {
	// TODO implement the business logic of Say
	return res, err
}

func (b *basicHelloService) Say1(ctx context.Context, req SayReq) (res SayRes, err error) {
	// TODO implement the business logic of Say1
	return res, err
}

func (b *basicHelloService) SayHello(ctx context.Context, req SayReq) (res SayRes, err error) {
	// TODO implement the business logic of SayHello
	return res, err
}

func (b *basicHelloService) SayHello1(ctx context.Context, s1 string, s2 string) (res SayRes, err error) {
	// TODO implement the business logic of SayHello1
	return res, err
}

func (b *basicHelloService) SayMan(ctx context.Context, id int, name string, email string) (res SayRes, err error) {
	// TODO implement the business logic of SayMan
	return res, err
}

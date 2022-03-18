package hello

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"hello/pkg/repository"
	"hello/pkg/services/hello/types"
)

// HelloService describes the service.
////go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/prometheus.tmpl:service_with_prometheus.go
////go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/log.tmpl:service_with_log.go
////go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/opentracing.tmpl:service_with_trace.go
////go:generate gowrap gen -g -p ./ -i HelloService -t ../../gowrap/templates/endpoint.tmpl:endpoint.go
//go:generate gowrap gen -g -p ./ -i HelloService -bt "../../gowrap/templates/http-gin.tmpl:http_gen.go ../../gowrap/templates/prometheus.tmpl:service_with_prometheus_gen.go ../../gowrap/templates/log.tmpl:service_with_log_gen.go ../../gowrap/templates/opentracing.tmpl:service_with_trace_gen.go ../../gowrap/templates/endpoint.tmpl:endpoint_gen.go"
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

type basicHelloService struct {
	log  *zap.SugaredLogger
	repo *repository.Repository
}

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
	root, err := b.repo.Baidu.GetRoot(ctx)
	if err != nil {
		return types.SayRes{}, err
	}
	res.Name = root.String()
	return
}

func (b *basicHelloService) SayHello1(ctx context.Context, s1 types.SayReq) (res types.SayRes, err error) {
	return res, fmt.Errorf("sayhello1 error")
}

type Middleware func(HelloService) HelloService

type BaseService HelloService

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService(log *zap.SugaredLogger, repo *repository.Repository) BaseService {
	return &basicHelloService{
		log:  log,
		repo: repo,
	}
}

// New returns a HelloService with all of the expected middleware wired in.
func NewService(svc BaseService, middleware []Middleware) HelloService {

	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}

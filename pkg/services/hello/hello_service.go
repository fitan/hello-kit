package hello

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/repository"
)

// HelloService describes the service.
//go:generate gowrap gen -g -p ./ -i HelloService -bt "prometheus:hello_with_prometheus_gen.go log:hello_with_log_gen.go opentracing:hello_with_trace_gen.go http-gin:hello_http_gen.go endpoint:hello_endpoint_gen.go"
type HelloService interface {
	// @http-gin /hello/:id GET
	GetUser(ctx context.Context, req GetUserReq) (ent.PodBaseGetRes, error)
}

type basicHelloService struct {
	repo repository.Repository
}

type GetUserReq struct {
	Uri struct{
		Id int64 `uri:"id" json:"id"`
	} `json:"uri"`
}

func (b *basicHelloService)GetUser(ctx context.Context, req GetUserReq) (ent.PodBaseGetRes, error) {
	return b.repo.Pod.GetById(ctx, req.Uri.Id)
}

type BaseService HelloService

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicService(repo repository.Repository) BaseService {
	return &basicHelloService{repo: repo}
}

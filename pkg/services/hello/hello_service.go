package hello

import (
	"context"
	"hello/pkg/ent"
)

// HelloService describes the service.
//go:generate gowrap gen -g -p ./ -i HelloService -bt "prometheus:hello_with_prometheus_gen.go log:hello_with_log_gen.go opentracing:hello_with_trace_gen.go http-gin:hello_http_gen.go endpoint:hello_endpoint_gen.go"
type HelloService interface {
	// @http-gin /user/:id GET
	Say(ctx context.Context, req SayReq) (*ent.User, error)
}

type basicHelloService struct {
	db *ent.Client
}

type SayReq struct {
	Uri struct {
		Id int `uri:"id"`
	}
}

func (s *basicHelloService) Say(ctx context.Context, req SayReq) (*ent.User, error) {
	return s.db.User.Get(ctx, req.Uri.Id)
}

type BaseService HelloService

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicHelloService{db: db}
}

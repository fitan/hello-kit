package say1

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/repository"
)

// Say1Service describes the service.
//go:generate gowrap gen -g -p ./ -i Say1Service -bt "prometheus:say1_with_prometheus_gen.go log:say1_with_log_gen.go opentracing:say1_with_trace_gen.go http-gin:say1_http_gen.go endpoint:say1_endpoint_gen.go"
type Say1Service interface {
	// @http-gin /say1pod/:id GET
	SayPod(ctx context.Context, req SayPodReq) (ent.PodBaseGetRes, error)
}

type basicSay1Service struct {

	repo repository.Repository
}

type SayPodReq struct {
	Uri struct{
		Id int64 `uri:"id" json:"id"`
	}
}

func (b *basicSay1Service) SayPod(ctx context.Context, req SayPodReq) (ent.PodBaseGetRes, error) {
	return b.repo.Pod.GetById(ctx, req.Uri.Id)
}

type BaseService Say1Service

// NewBasicSay1Service returns a naive, stateless implementation of Say1Service.
func NewBasicService(repo repository.Repository) BaseService {
	return &basicSay1Service{repo: repo}
}

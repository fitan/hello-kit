package say

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/repository"
)

// SayService describes the service.
//go:generate gowrap gen -g -p ./ -i SayService -bt "prometheus:say_with_prometheus_gen.go log:say_with_log_gen.go opentracing:say_with_trace_gen.go http-gin:say_http_gen.go endpoint:say_endpoint_gen.go"
type SayService interface {
    // @http-gin /saypod/:id GET
	SayPod(ctx context.Context, req SayPodReq) (*ent.Pod, error)
}

type basicSayService struct {
	repo repository.Repository
}

type SayPodReq struct {
	Uri struct{
		Id int64 `uri:"id" json:"id"`
	}
}

func (b *basicSayService) SayPod(ctx context.Context, req SayPodReq) (*ent.Pod, error) {
	return b.repo.Pod.GetById(ctx, req.Uri.Id)
}

type BaseService SayService

// NewBasicSayService returns a naive, stateless implementation of SayService.
func NewBasicService(repo repository.Repository) BaseService {
	return &basicSayService{repo: repo}
}

package say

import (
	"context"
	"hello/pkg/services/hello/types"
)

// SayService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i SayService -bt "prometheus:say_with_prometheus_gen.go log:say_with_log_gen.go opentracing:say_with_trace_gen.go http-gin:say_http_gen.go endpoint:say_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i SayService -bt "prometheus:say_with_prometheus_gen.go log:say_with_log_gen.go opentracing:say_with_trace_gen.go"
type SayService interface {
	// @http-gin /say/foo/:id POST
	Say(ctx context.Context, req types.SayReq) (res types.SayRes, err error)
}

type basicSayService struct {
}

func (b basicSayService) Say(ctx context.Context, req types.SayReq) (res types.SayRes, err error) {
	panic("implement me")
}

type BaseService SayService

// NewBasicSayService returns a naive, stateless implementation of SayService.
func NewBasicService() BaseService {
	return &basicSayService{}
}

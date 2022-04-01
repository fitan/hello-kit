package pod

import "hello/pkg/ent"

// PodService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i PodService -bt "prometheus:pod_with_prometheus_gen.go log:pod_with_log_gen.go opentracing:pod_with_trace_gen.go http-gin:pod_http_gen.go endpoint:pod_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i PodService -bt "prometheus:pod_with_prometheus_gen.go log:pod_with_log_gen.go opentracing:pod_with_trace_gen.go"
type PodService interface {
	ent.PodBaseInterface
}

type basicPodService struct {
	ent.PodBaseInterface
}

type BaseService PodService

// NewBasicPodService returns a naive, stateless implementation of PodService.
func NewBasicService(client *ent.Client) BaseService {
	return &basicPodService{ent.NewPodBase(client.Pod)}
}

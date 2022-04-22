package pod

import "hello/pkg/ent"

// PodService describes the service.
//go:generate gowrap gen -g -p ./ -i PodService -bt "prometheus:pod_with_prometheus_gen.go log:pod_with_log_gen.go opentracing:pod_with_trace_gen.go http-gin:pod_http_gen.go endpoint:pod_endpoint_gen.go"
type PodService interface {
	ent.PodRestInterface
}

type basicPodService struct {
	ent.PodRestInterface
}

type BaseService PodService

// NewBasicPodService returns a naive, stateless implementation of PodService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicPodService{ent.NewPodRest(db)}
}

package resource

import "hello/pkg/ent"

// ResourceService describes the service.
//go:generate gowrap gen -g -p ./ -i ResourceService -bt "prometheus:resource_with_prometheus_gen.go log:resource_with_log_gen.go opentracing:resource_with_trace_gen.go http-gin:resource_http_gen.go endpoint:resource_endpoint_gen.go"
type ResourceService interface {
	ent.ResourceRestInterface
}

type basicResourceService struct {
	ent.ResourceRestInterface
}

type BaseService ResourceService

// NewBasicResourceService returns a naive, stateless implementation of ResourceService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicResourceService{ent.NewResourceRest(db)}
}

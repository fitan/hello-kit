package service

import "hello/pkg/ent"

// ServiceService describes the service.
//go:generate gowrap gen -g -p ./ -i ServiceService -bt "prometheus:service_with_prometheus_gen.go log:service_with_log_gen.go opentracing:service_with_trace_gen.go http-gin:service_http_gen.go endpoint:service_endpoint_gen.go"
type ServiceService interface {
	ent.ServiceRestInterface
}

type basicServiceService struct {
	ent.ServiceRestInterface
}

type BaseService ServiceService

// NewBasicServiceService returns a naive, stateless implementation of ServiceService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicServiceService{ent.NewServiceRest(db)}
}

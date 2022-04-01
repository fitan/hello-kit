package project

import "hello/pkg/ent"

// ProjectService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i ProjectService -bt "prometheus:project_with_prometheus_gen.go log:project_with_log_gen.go opentracing:project_with_trace_gen.go http-gin:project_http_gen.go endpoint:project_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i ProjectService -bt "prometheus:project_with_prometheus_gen.go log:project_with_log_gen.go opentracing:project_with_trace_gen.go"
type ProjectService interface {
	ent.ProjectBaseInterface
}

type basicProjectService struct {
	ent.ProjectBaseInterface
}

type BaseService ProjectService

// NewBasicProjectService returns a naive, stateless implementation of ProjectService.
func NewBasicService(client *ent.Client) BaseService {
	return &basicProjectService{ent.NewProjectBase(client.Project)}
}

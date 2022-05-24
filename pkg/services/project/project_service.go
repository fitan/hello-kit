package project

import (
	"hello/pkg/ent"
	"hello/pkg/repository"
)

// ProjectService describes the service.
//go:generate gowrap gen -g -p ./ -i ProjectService -bt "prometheus:project_with_prometheus_gen.go log:project_with_log_gen.go opentracing:project_with_trace_gen.go http-gin:project_http_gen.go endpoint:project_endpoint_gen.go"
type ProjectService interface {
	ent.ProjectRestInterface
}

type basicProjectService struct {
	repo repository.Repository
	ent.ProjectRestInterface
}

type BaseService ProjectService

// NewBasicProjectService returns a naive, stateless implementation of ProjectService.
func NewBasicService(repo repository.Repository, db *ent.Client) BaseService {
	return &basicProjectService{repo, ent.NewProjectRest(db)}
}

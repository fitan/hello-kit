package project

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/ent/project"
)

// ProjectService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i ProjectService -bt "prometheus:project_with_prometheus_gen.go log:project_with_log_gen.go opentracing:project_with_trace_gen.go"
type ProjectService interface {
	ent.ProjectBaseInterface
	GetOneWithServiceQsById(ctx context.Context, id int, qs interface{}) (*ent.Project, error)
}

type basicProjectService struct {
	ent.ProjectBaseInterface
	db *ent.Client
}

type BaseService ProjectService

func (s *basicProjectService) GetOneWithServiceQsById(ctx context.Context, id int, qs interface{}) (
	*ent.Project,
	error,
) {
	return s.db.Project.Query().Where(project.IDEQ(id)).WithServices(
		func(query *ent.ServiceQuery) {
			query.Queries(qs)
		}).First(ctx)
}

// NewBasicProjectService returns a naive, stateless implementation of ProjectService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicProjectService{ent.NewProjectBase(db), db}
}

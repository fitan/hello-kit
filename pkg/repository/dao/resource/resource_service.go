package resource

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/ent/resource"
)

// ResourceService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i ResourceService -bt "prometheus:resource_with_prometheus_gen.go log:resource_with_log_gen.go opentracing:resource_with_trace_gen.go"
type ResourceService interface {
	ent.ResourceBaseInterface
	IdByResource(ctx context.Context, path string, action string) (int, error)
}

type basicResourceService struct {
	ent.ResourceBaseInterface
	db *ent.Client
}

func (s *basicResourceService) IdByResource(ctx context.Context, path, action string) (int, error) {
	return s.db.Resource.Query().Where(resource.PathEQ(path), resource.ActionEQ(action)).FirstID(ctx)
}

type BaseService ResourceService

// NewBasicResourceService returns a naive, stateless implementation of ResourceService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicResourceService{
		ent.NewResourceBase(db),
		db,
	}
}

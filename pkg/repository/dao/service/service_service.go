package service

import (
	"context"
	"hello/pkg/ent"
)

// ServiceService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i ServiceService -bt "prometheus:service_with_prometheus_gen.go log:service_with_log_gen.go opentracing:service_with_trace_gen.go"
type ServiceService interface {
	ent.ServiceBaseInterface
}

type basicServiceService struct {
	ent.ServiceBaseInterface
	db *ent.Client
}

type BaseService ServiceService

func (s *basicServiceService) ListByQs(ctx context.Context, qs interface{}) (res ent.Services, total int, err error) {
	return s.db.Service.Query().WithProject().ByQueriesAll(ctx, nil)
}

// NewBasicServiceService returns a naive, stateless implementation of ServiceService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicServiceService{ent.NewServiceBase(db), db}
}

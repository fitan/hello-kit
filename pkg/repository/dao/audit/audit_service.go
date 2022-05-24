package audit

import "hello/pkg/ent"

// AuditService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i AuditService -bt "prometheus:audit_with_prometheus_gen.go log:audit_with_log_gen.go opentracing:audit_with_trace_gen.go"
type AuditService interface {
	ent.AuditBaseInterface
}

type basicAuditService struct {
	ent.AuditBaseInterface
}

type BaseService AuditService

// NewBasicAuditService returns a naive, stateless implementation of AuditService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicAuditService{ent.NewAuditBase(db)}
}

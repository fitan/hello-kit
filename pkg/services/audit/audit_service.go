package audit

import "hello/pkg/ent"

// AuditService describes the service.
//go:generate gowrap gen -g -p ./ -i AuditService -bt "prometheus:audit_with_prometheus_gen.go log:audit_with_log_gen.go opentracing:audit_with_trace_gen.go http-gin:audit_http_gen.go endpoint:audit_endpoint_gen.go"
type AuditService interface {
	ent.AuditRestInterface
}

type basicAuditService struct {
	ent.AuditRestInterface
}

type BaseService AuditService

// NewBasicAuditService returns a naive, stateless implementation of AuditService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicAuditService{
		ent.NewAuditRest(db),
	}
}

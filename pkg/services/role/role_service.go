package role

import "hello/pkg/ent"

// RoleService describes the service.
//go:generate gowrap gen -g -p ./ -i RoleService -bt "prometheus:role_with_prometheus_gen.go log:role_with_log_gen.go opentracing:role_with_trace_gen.go http-gin:role_http_gen.go endpoint:role_endpoint_gen.go"
type RoleService interface {
	ent.RoleRestInterface
}

type basicRoleService struct {
	ent.RoleRestInterface
}

type BaseService RoleService

// NewBasicRoleService returns a naive, stateless implementation of RoleService.
func NewBasicService(db *ent.Client) BaseService {
	return &basicRoleService{ent.NewRest(db)}
}

package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/pkg/errors"
	"hello/pkg/ent"
)

// CasbinService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i CasbinService -bt "prometheus:casbin_with_prometheus_gen.go log:casbin_with_log_gen.go opentracing:casbin_with_trace_gen.go http-gin:casbin_http_gen.go endpoint:casbin_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i CasbinService -bt "prometheus:casbin_with_prometheus_gen.go log:casbin_with_log_gen.go opentracing:casbin_with_trace_gen.go"
type CasbinService interface {
	BindPermission(ctx context.Context, userId, roleId, resourceId int) error
	UnBindPermission(ctx context.Context, userId, roleId, resourceId int) error
}

type basicCasbinService struct {
	e  *casbin.Enforcer
	db *ent.Client
}

func (s *basicCasbinService) BindPermission(ctx context.Context, userId, roleId, resourceId int) error {
	user := CasbinUserKeyWrap(userId)
	role := CasbinRoleKeyWrap(roleId)
	resource := CasbinResourceWrap(resourceId)

	_, err := s.e.AddGroupingPolicy(user, role, resource)
	if err != nil {
		return errors.Wrap(err, "AddGroupingPolicy")
	}
	return nil
}

func (s *basicCasbinService) UnBindPermission(ctx context.Context, userId, roleId, resourceId int) error {
	user := CasbinUserKeyWrap(userId)
	role := CasbinRoleKeyWrap(roleId)
	resource := CasbinResourceWrap(resourceId)
	_, err := s.e.RemovePolicy(user, role, resource)
	if err != nil {
		return errors.Wrap(err, "RemovePolicy")
	}
	return nil
}

type BaseService CasbinService

// NewBasicCasbinService returns a naive, stateless implementation of CasbinService.
func NewBasicService() BaseService {
	return &basicCasbinService{}
}

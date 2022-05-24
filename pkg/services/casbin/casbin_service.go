package casbin

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/pkg/errors"
	"hello/pkg/ent"
	"strconv"
	"strings"
)

// CasbinService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i CasbinService -bt "prometheus:casbin_with_prometheus_gen.go log:casbin_with_log_gen.go opentracing:casbin_with_trace_gen.go http-gin:casbin_http_gen.go endpoint:casbin_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i CasbinService -bt "prometheus:casbin_with_prometheus_gen.go log:casbin_with_log_gen.go opentracing:casbin_with_trace_gen.go"
type CasbinService interface {
	GetPermissionByUser(ctx context.Context, permission Permission) (res []Permission)
	CheckPermission(ctx context.Context, permission CheckPermission) (bool, error)
	BindPermission(ctx context.Context, permission Permission) error
	UnBindPermission(ctx context.Context, permission Permission) error
	AddRoleAuthorization(ctx context.Context, role Role) error
	DelRoleAuthorization(ctx context.Context, role Role) error
}

type basicCasbinService struct {
	e  *casbin.Enforcer
	db *ent.Client
}

func (s *basicCasbinService) GetPermissionByUser(ctx context.Context, permission Permission) (res []Permission) {
	policies := s.e.GetFilteredGroupingPolicy(1, permission.UserKey())
	for _, v := range policies {
		userKey := v[0]
		roleKey := v[1]
		domainKey := v[2]

		userIdStr := strings.TrimPrefix(userKey, UserPre)
		roleIdStr := strings.TrimPrefix(roleKey, RolePre)
		userId, _ := strconv.Atoi(userIdStr)
		roleId, _ := strconv.Atoi(roleIdStr)

		p := Permission{
			UserId: userId,
			RoleId: roleId,
			Domain: domainKey,
		}
		res = append(res, p)
	}
	return
}

func (s *basicCasbinService) CheckPermission(ctx context.Context, permission CheckPermission) (bool, error) {

	has, err := s.e.Enforce(permission.CasbinObj()...)
	return has, errors.Wrap(err, "EnforceEx")
}

func (s *basicCasbinService) BindPermission(ctx context.Context, permission Permission) error {

	_, err := s.e.AddGroupingPolicy(permission.CasbinObj()...)
	if err != nil {
		return errors.Wrap(err, "AddGroupingPolicy")
	}
	return nil
}

func (s *basicCasbinService) UnBindPermission(ctx context.Context, permission Permission) error {
	_, err := s.e.RemovePolicy(permission.CasbinObj()...)
	if err != nil {
		return errors.Wrap(err, "RemovePolicy")
	}
	return nil
}

func (s *basicCasbinService) AddRoleAuthorization(ctx context.Context, role Role) error {
	_, err := s.e.AddPolicy(role.CasbinObj()...)
	if err != nil {
		err = errors.Wrap(err, "AddPolicy")
		return err
	}
	return nil
}

func (s *basicCasbinService) DelRoleAuthorization(ctx context.Context, role Role) error {
	_, err := s.e.RemovePolicy(role.CasbinObj()...)
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

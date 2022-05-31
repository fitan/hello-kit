package services

import (
	"github.com/google/wire"
	"hello/pkg/services/audit"
	"hello/pkg/services/casbin"
	"hello/pkg/services/project"
	"hello/pkg/services/resource"
	"hello/pkg/services/role"
	"hello/pkg/services/service"
	"hello/pkg/services/user"
)

type HttpHandler struct {
	User     user.HttpHandler
	Audit    audit.HttpHandler
	Project  project.HttpHandler
	Service  service.HttpHandler
	Resource resource.HttpHandler
	Role     role.HttpHandler
}

type Services struct {
	User     user.UserService
	Audit    audit.AuditService
	Project  project.ProjectService
	Casbin   casbin.CasbinService
	Service  service.ServiceService
	Resource resource.ResourceService
	Role     role.RoleService
}

var ServicesSet = wire.NewSet(
	role.RoleKitSet,
	resource.ResourceKitSet,
	service.ServiceKitSet,
	project.ProjectKitSet,
	audit.AuditKitSet,
	user.UserServiceSet,
	casbin.CasbinSet,
	wire.Struct(new(HttpHandler), "*"),
	wire.Struct(new(Services), "*"),
)

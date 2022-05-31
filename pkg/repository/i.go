package repository

import (
	"github.com/google/wire"
	"hello/pkg/repository/api/baidu"
	"hello/pkg/repository/api/taobao"
	"hello/pkg/repository/dao/audit"
	"hello/pkg/repository/dao/project"
	"hello/pkg/repository/dao/resource"
	"hello/pkg/repository/dao/service"
	"hello/pkg/repository/dao/user"
)

type Repository struct {
	Baidu    baidu.BaiduService
	Taobao   taobao.TaobaoService
	User     user.UserService
	Audit    audit.AuditService
	Project  project.ProjectService
	Service  service.ServiceService
	Resource resource.ResourceService
}

var RepoSet = wire.NewSet(
	baidu.BaiduServiceSet,
	taobao.TaobaoServiceSet,

	resource.ResourceServiceSet,
	service.ServiceServiceSet,
	project.ProjectServiceSet,
	audit.AuditServiceSet,
	user.UserServiceSet,
	wire.Struct(new(Repository), "*"),
)

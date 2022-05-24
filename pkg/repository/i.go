package repository

import (
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

//go:build wireinject
// +build wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/oklog/oklog/pkg/group"
	"hello/pkg/repository"
	"hello/pkg/repository/api/baidu"
	"hello/pkg/repository/api/taobao"
	auditD "hello/pkg/repository/dao/audit"
	projectD "hello/pkg/repository/dao/project"
	"hello/pkg/repository/dao/resource"
	serviceD "hello/pkg/repository/dao/service"
	userD "hello/pkg/repository/dao/user"
	"hello/pkg/services"
	"hello/pkg/services/audit"
	"hello/pkg/services/casbin"
	"hello/pkg/services/project"
	"hello/pkg/services/user"
)

var confSet = wire.NewSet(initConf)
var logSet = wire.NewSet(initLog)
var traceSet = wire.NewSet(initTracer)
var dbSet = wire.NewSet(initDb)
var entSet = wire.NewSet(initEnt)
var pyroscopeSet = wire.NewSet(initPyroscope)
var casbinSet = wire.NewSet(initCasbin)
var auditMidSet = wire.NewSet(initAuditMid)
var debugSwitchSet = wire.NewSet(initDebugSwitch)

// repo.api.service
var baiduHttpSet = wire.NewSet(baidu.NewBasicService, baidu.NewServiceMiddleware, baidu.NewService)
var taobaoHttpSet = wire.NewSet(taobao.NewBasicService, taobao.NewServiceMiddleware, taobao.NewService)

// repo.dao.service
var userDaoSet = wire.NewSet(userD.NewBasicService, userD.NewServiceMiddleware, userD.NewService)

//var projectDaoSet = wire.NewSet(projectD.NewBasicService, projectD.NewServiceMiddleware, projectD.NewService)

var repoSet = wire.NewSet(
	resource.ResourceServiceSet,
	serviceD.ServiceServiceSet,
	projectD.ProjectServiceSet, userDaoSet, baiduHttpSet, taobaoHttpSet, auditD.AuditServiceSet, wire.Struct(new(repository.Repository), "*"))

// http service
var casbinServiceSet = wire.NewSet(casbin.NewBasicService, casbin.NewService, casbin.NewServiceMiddleware)

//var userServiceSet = wire.NewSet(user.NewBasicService, user.NewService, user.NewEndpointMiddleware, user.NewServiceMiddleware, user.NewEndpoints, user.NewServiceOption, user.NewHTTPHandler)

var servicesSet = wire.NewSet(project.ProjectKitSet, audit.AuditKitSet, casbinServiceSet, user.UserServiceSet, wire.Struct(new(services.HttpHandler), "*"), wire.Struct(new(services.Services), "*"))

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initMicro)

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, g *group.Group, name ConfName) (App, error) {
	wire.Build(debugSwitchSet, auditMidSet, casbinSet, entSet, pyroscopeSet, dbSet, gSet, mwSet, confSet, logSet, traceSet, repoSet, servicesSet, wire.Struct(new(App), "*"))
	return App{}, nil
}

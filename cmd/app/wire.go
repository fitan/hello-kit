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
	podD "hello/pkg/repository/dao/pod"
	projectD "hello/pkg/repository/dao/project"
	tblservicetreeD "hello/pkg/repository/dao/tblservicetree"
	userD "hello/pkg/repository/dao/user"
	"hello/pkg/services"
	"hello/pkg/services/casbin"
	"hello/pkg/services/hello"
	"hello/pkg/services/say"
	"hello/pkg/services/say1"
	"hello/pkg/services/user"
)

var confSet = wire.NewSet(initConf)
var logSet = wire.NewSet(initLog)
var traceSet = wire.NewSet(initTracer)
var dbSet = wire.NewSet(initDb)
var entSet = wire.NewSet(initEnt)
var pyroscopeSet = wire.NewSet(initPyroscope)
var casbinSet = wire.NewSet(initCasbin)

// repo.api.service
var baiduHttpSet = wire.NewSet(baidu.NewBasicService, baidu.NewServiceMiddleware, baidu.NewService)
var taobaoHttpSet = wire.NewSet(taobao.NewBasicService, taobao.NewServiceMiddleware, taobao.NewService)

// repo.dao.service
var userDaoSet = wire.NewSet(userD.NewBasicService, userD.NewServiceMiddleware, userD.NewService)
var podDaoSet = wire.NewSet(podD.NewBasicService, podD.NewServiceMiddleware, podD.NewService)
var projectDaoSet = wire.NewSet(projectD.NewBasicService, projectD.NewServiceMiddleware, projectD.NewService)
var tblservicetreeDaoSet = wire.NewSet(tblservicetreeD.NewBasicService, tblservicetreeD.NewServiceMiddleware, tblservicetreeD.NewService)




var repoSet = wire.NewSet(podDaoSet,projectDaoSet,tblservicetreeDaoSet,userDaoSet, baiduHttpSet, taobaoHttpSet, wire.Struct(new(repository.Repository), "*"))

// http service
var casbinServiceSet = wire.NewSet(casbin.NewBasicService, casbin.NewService, casbin.NewServiceMiddleware)

var helloServiceSet = wire.NewSet(hello.NewBasicService, hello.NewService, hello.NewEndpointMiddleware, hello.NewServiceMiddleware, hello.NewEndpoints, hello.NewServiceOption, hello.NewHTTPHandler)
//var userServiceSet = wire.NewSet(user.NewBasicService, user.NewService, user.NewEndpointMiddleware, user.NewServiceMiddleware, user.NewEndpoints, user.NewServiceOption, user.NewHTTPHandler)
var servicesSet = wire.NewSet(casbinServiceSet, say.SayKitSet,say1.Say1KitSet,user.UserServiceSet,helloServiceSet, wire.Struct(new(services.Services), "*"))

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initMicro)

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, g *group.Group, name ConfName) (App, error) {
	wire.Build(casbinSet, entSet, pyroscopeSet, dbSet, gSet, mwSet, confSet, logSet, traceSet, repoSet, servicesSet, wire.Struct(new(App), "*"))
	return App{}, nil
}

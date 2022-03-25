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
	userD "hello/pkg/repository/dao/user"
	"hello/pkg/services"
	"hello/pkg/services/casbin"
	"hello/pkg/services/hello"
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
var baiduHttpSet = wire.NewSet(baidu.NewBaiduApi, baidu.NewBase, baidu.NewBaiduApiMiddleware)
var taobaoHttpSet = wire.NewSet(taobao.NewTaoApi, taobao.NewBase, taobao.NewTaobaoApiMiddleware)

// repo.dao.service
var userDaoSet = wire.NewSet(userD.NewBasicService, userD.NewServiceMiddleware, userD.NewService)

var repoSet = wire.NewSet(userDaoSet, baiduHttpSet, taobaoHttpSet, wire.Struct(new(repository.Repository), "*"))

// http service
var casbinServiceSet = wire.NewSet(casbin.NewBasicService, casbin.NewService, casbin.NewServiceMiddleware)

var helloServiceSet = wire.NewSet(hello.NewBasicHelloService, hello.NewService, hello.NewEndpointMiddleware, hello.NewServiceMiddleware, hello.NewEndpoints, hello.NewServiceOption, hello.NewHTTPHandler)
var userServiceSet = wire.NewSet(user.NewBasicService, user.NewService, user.NewEndpointMiddleware, user.NewServiceMiddleware, user.NewEndpoints, user.NewServiceOption, user.NewHTTPHandler)
var servicesSet = wire.NewSet(casbinServiceSet,userServiceSet,helloServiceSet, wire.Struct(new(services.Services), "*"))

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initMicro)

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, g *group.Group, name ConfName) (App, error) {
	wire.Build(casbinSet, entSet, pyroscopeSet, dbSet, gSet, mwSet, confSet, logSet, traceSet, repoSet, servicesSet, wire.Struct(new(App), "*"))
	return App{}, nil
}

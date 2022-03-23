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
	"hello/pkg/repository/dao/user"
	"hello/pkg/services"
	"hello/pkg/services/hello"
)

var confSet = wire.NewSet(initMicroConf)
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
var userServiceSet = wire.NewSet(user.NewBasicService, user.NewServiceMiddleware, user.NewService)

var repoSet = wire.NewSet(userServiceSet, baiduHttpSet, taobaoHttpSet, wire.Struct(new(repository.Repository), "*"))

// http service
var helloServiceSet = wire.NewSet(hello.NewBasicHelloService, hello.NewService, hello.NewEndpointMiddleware, hello.NewServiceMiddleware, hello.NewEndpoints, hello.NewServiceOption, hello.NewHTTPHandler)
var servicesSet = wire.NewSet(helloServiceSet, wire.Struct(new(services.Services), "*"))

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initMicro)

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, g *group.Group, name ConfName) (App, error) {
	wire.Build(casbinSet, entSet, pyroscopeSet, dbSet, gSet, mwSet, confSet, logSet, traceSet, repoSet, servicesSet, wire.Struct(new(App), "*"))
	return App{}, nil
}

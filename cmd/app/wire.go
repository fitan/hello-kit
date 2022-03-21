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
	"hello/pkg/services"
	"hello/pkg/services/hello"
)

var confSet = wire.NewSet(initConf)
var logSet = wire.NewSet(initLog)
var traceSet = wire.NewSet(initTracer)
var dbSet = wire.NewSet(initDb)
var pyroscopeSet = wire.NewSet(initPyroscope)

var baiduHttpSet = wire.NewSet(baidu.NewBaiduApi, baidu.NewBase, baidu.NewBaiduApiMiddleware)
var taobaoHttpSet = wire.NewSet(taobao.NewTaoApi, taobao.NewBase, taobao.NewTaobaoApiMiddleware)
var repoSet = wire.NewSet(baiduHttpSet, taobaoHttpSet, wire.Struct(new(repository.Repository), "*"))

var helloServiceSet = wire.NewSet(hello.NewBasicHelloService, hello.NewService, hello.NewEndpointMiddleware, hello.NewServiceMiddleware, hello.NewEndpoints, hello.NewServiceOption, hello.NewHTTPHandler)
var servicesSet = wire.NewSet(helloServiceSet, wire.Struct(new(services.Services), "*"))

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initHttpHandler)

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, g *group.Group, name ConfName) (App, error) {
	wire.Build(pyroscopeSet, dbSet, gSet, mwSet, confSet, logSet, traceSet, repoSet, servicesSet, wire.Struct(new(App), "*"))
	return App{}, nil
}

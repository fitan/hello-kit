// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/oklog/run"
	"hello/pkg/repository"
	"hello/pkg/repository/api/baidu"
	"hello/pkg/repository/api/taobao"
	"hello/pkg/repository/dao/user"
	"hello/pkg/services"
	"hello/pkg/services/hello"
)

// Injectors from wire.go:

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, g *run.Group, name ConfName) (App, error) {
	myConf := initConf(name)
	base := baidu.NewBase(myConf)
	sugaredLogger := initLog(myConf)
	v := baidu.NewBaiduApiMiddleware(sugaredLogger)
	baiduApi := baidu.NewBaiduApi(base, v)
	taobaoBase := taobao.NewBase(myConf)
	v2 := taobao.NewTaobaoApiMiddleware(sugaredLogger)
	taobaoApi := taobao.NewTaoApi(taobaoBase, v2)
	client, err := initEnt(myConf)
	if err != nil {
		return App{}, err
	}
	baseService := user.NewBasicService(client)
	v3 := user.NewServiceMiddleware(sugaredLogger)
	userService := user.NewService(baseService, v3)
	repositoryRepository := &repository.Repository{
		Baidu:  baiduApi,
		Taobao: taobaoApi,
		User:   userService,
	}
	helloBaseService := hello.NewBasicHelloService(sugaredLogger, repositoryRepository)
	v4 := hello.NewServiceMiddleware(sugaredLogger, myConf)
	helloService := hello.NewService(helloBaseService, v4)
	v5 := initEndpointMiddleware()
	v6 := hello.NewEndpointMiddleware(sugaredLogger, v5)
	endpoints := hello.NewEndpoints(helloService, v6)
	v7 := initHttpServerOption()
	v8 := hello.NewServiceOption(v7)
	wireHttpHandler := hello.NewHTTPHandler(r, endpoints, v8)
	servicesServices := &services.Services{
		Hello: wireHttpHandler,
	}
	tracerProvider := initTracer(myConf)
	db, err := initDb(myConf)
	if err != nil {
		return App{}, err
	}
	profiler, err := initPyroscope(myConf)
	if err != nil {
		return App{}, err
	}
	appInitCancelInterrupt := initCancelInterrupt(g)
	appInitMetricsEndpoint := initMetricsEndpoint(g, myConf)
	appInitHttpHandler := initHttpHandler(r, g, sugaredLogger, myConf)
	app := App{
		r:                   r,
		repository:          repositoryRepository,
		services:            servicesServices,
		g:                   g,
		conf:                myConf,
		log:                 sugaredLogger,
		tp:                  tracerProvider,
		db:                  db,
		ent:                 client,
		pyroscope:           profiler,
		InitCancelInterrupt: appInitCancelInterrupt,
		InitMetricsEndpoint: appInitMetricsEndpoint,
		InitHttpHandler:     appInitHttpHandler,
	}
	return app, nil
}

// wire.go:

var confSet = wire.NewSet(initConf)

var logSet = wire.NewSet(initLog)

var traceSet = wire.NewSet(initTracer)

var dbSet = wire.NewSet(initDb)

var entSet = wire.NewSet(initEnt)

var pyroscopeSet = wire.NewSet(initPyroscope)

// http api
var baiduHttpSet = wire.NewSet(baidu.NewBaiduApi, baidu.NewBase, baidu.NewBaiduApiMiddleware)

var taobaoHttpSet = wire.NewSet(taobao.NewTaoApi, taobao.NewBase, taobao.NewTaobaoApiMiddleware)

// dao service
var userServiceSet = wire.NewSet(user.NewBasicService, user.NewServiceMiddleware, user.NewService)

var repoSet = wire.NewSet(userServiceSet, baiduHttpSet, taobaoHttpSet, wire.Struct(new(repository.Repository), "*"))

// http service
var helloServiceSet = wire.NewSet(hello.NewBasicHelloService, hello.NewService, hello.NewEndpointMiddleware, hello.NewServiceMiddleware, hello.NewEndpoints, hello.NewServiceOption, hello.NewHTTPHandler)

var servicesSet = wire.NewSet(helloServiceSet, wire.Struct(new(services.Services), "*"))

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initHttpHandler)

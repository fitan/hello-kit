//go:build wireinject
// +build wireinject

package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/oklog/oklog/pkg/group"
	"hello/pkg/repository"
	"hello/pkg/services"
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
var DefaultHttpClient = wire.NewSet(initDefaultHttpClient)
var initNatsSet = wire.NewSet(initNats)

var mwSet = wire.NewSet(initEndpointMiddleware, initHttpServerOption)

var gSet = wire.NewSet(initCancelInterrupt, initMetricsEndpoint, initMicro)

var initSet = wire.NewSet(
	initNatsSet,
	DefaultHttpClient,
	debugSwitchSet,
	auditMidSet,
	casbinSet,
	entSet,
	pyroscopeSet,
	dbSet,
	gSet,
	mwSet,
	confSet,
	logSet,
	traceSet,
)

//var appSet = wire.NewSet(wire.Struct(App{}, "*"))
func InitApp(r *gin.Engine, router *gin.RouterGroup, g *group.Group, name ConfName) (App, error) {
	wire.Build(
		repository.RepoSet,
		services.ServicesSet,
		initSet,
		wire.Struct(new(App), "*"),
	)
	return App{}, nil
}

//go:build wireinject
// +build wireinject

package services

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"go.uber.org/zap"
	"hello/pkg/services/hello"
)

type wireServices struct {
	hello hello.WireHttpHandler
}

var helloSet = wire.NewSet(hello.NewBasicHelloService, hello.NewService, hello.NewEndpointMiddleware, hello.NewServiceMiddleware, hello.NewEndpoints, hello.NewServiceOption, hello.NewHTTPHandler)

var set = wire.NewSet(helloSet, wire.Struct(new(wireServices), "*"))

func InitHttp(
	r *gin.Engine, log *zap.SugaredLogger, appName string, eps []endpoint.Middleware, ops []http.ServerOption,
) wireServices {
	wire.Build(set)
	return wireServices{}
}

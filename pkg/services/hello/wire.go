//+build wireinject

package hello

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitHttpHandler(r *gin.Engine, log *zap.SugaredLogger, appName string, eps []endpoint.Middleware, ops []http.ServerOption) WireHttpHandler {
	wire.Build(NewBasicHelloService, NewService, NewEndpointMiddleware, NewServiceMiddleware, NewEndpoints, NewServiceOption, NewHTTPHandler)
	return WireHttpHandler{}
}

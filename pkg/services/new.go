package services

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
	"hello/pkg/services/hello"
)

func InitServices(r *gin.Engine, log *zap.SugaredLogger, appName string, epsmw []endpoint.Middleware, opts []http.ServerOption) {
	hello.InitHttpHandler(r, log, appName, epsmw, opts)
}

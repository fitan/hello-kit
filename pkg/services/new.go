package services

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"hello/pkg/services/hello"
)

func InitServices(r *gin.Engine, log log.Logger, appName string, epsmw []endpoint.Middleware, opts []http.ServerOption) {
	hello.InitHttpHandler(r, log, appName, epsmw, opts)
}

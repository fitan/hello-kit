package services

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"hello/pkg/services/hello"
)

func InitServices(r *gin.Engine, log log.Logger, appName string) {
	hello.InitHttpHandler(r, log, appName)
}

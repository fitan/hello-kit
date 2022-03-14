package services

import (
	"github.com/go-kit/log"
	"hello/pkg/services/hello"
	"net/http"
)

func InitServices(m *http.ServeMux, log log.Logger, appName string) {
	hello.InitHttpHandler(m, log, appName)
}

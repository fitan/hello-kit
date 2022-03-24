package services

import (
	"hello/pkg/services/hello"
	"hello/pkg/services/user"
)

type Services struct {
	Hello hello.HttpHandler
	User  user.HttpHandler
}

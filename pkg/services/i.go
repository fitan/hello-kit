package services

import (
	"hello/pkg/services/hello"
	"hello/pkg/services/pod"
	"hello/pkg/services/say"
	"hello/pkg/services/say1"
	"hello/pkg/services/user"
)

type Services struct {
	Hello hello.HttpHandler
	User  user.HttpHandler
	Say say.HttpHandler
	Say1 say1.HttpHandler
	Pod pod.HttpHandler
}

// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package service

import (
	endpoint1 "github.com/go-kit/kit/endpoint"
	log "github.com/go-kit/kit/log"
	http "github.com/go-kit/kit/transport/http"
	group "github.com/oklog/oklog/pkg/group"
	"hello/pkg/service"
)

func createService(endpoints service.Endpoints) (g *group.Group) {
	g = &group.Group{}
	initHttpHandler(endpoints, g)
	return g
}
func defaultHttpOptions(logger log.Logger) map[string][]http.ServerOption {
	options := map[string][]http.ServerOption{
		"Foo":       {},
		"Say":       {},
		"Say1":      {},
		"SayHello":  {},
		"SayHello1": {},
	}
	return options
}
func addEndpointMiddlewareToAllMethods(mw map[string][]endpoint1.Middleware, m endpoint1.Middleware) {
	methods := []string{"Foo", "Say", "Say1", "SayHello", "SayHello1"}
	for _, v := range methods {
		mw[v] = append(mw[v], m)
	}
}

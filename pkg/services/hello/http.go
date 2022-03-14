package hello

// Code generated by gowrap. DO NOT EDIT.
// template: ../../gowrap/templates/http-gin.tmpl
// gowrap: http://github.com/fitan/gowrap

import (
	"context"

	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
)

func AddHttpOptionToAllMethods(options map[string][]http.ServerOption, option http.ServerOption) {
	methods := []string{

		"SayHello",
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

func NewHTTPHandler(r *gin.Engine, endpoints Endpoints, options map[string][]http.ServerOption) {

	makeSayHelloHandler(r, endpoints, options["SayHello"])

}

func makeSayHelloHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/sayhello/:id", http.NewServer(endpoints.SayHelloEndpoint, decodeSayHelloRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeSayHelloRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req SayReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

package hello

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"

	"context"
)

type Ops map[string][]http.ServerOption

func AddHttpOptionToAllMethods(options map[string][]http.ServerOption, option http.ServerOption) {
	methods := []string{

		"GetUser",
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

type HttpHandler struct {
}

func NewHTTPHandler(r *gin.Engine, endpoints Endpoints, options Ops) HttpHandler {

	makeGetUserHandler(r, endpoints, options["GetUser"])

	return HttpHandler{}
}

type SwagResponse struct {
	TraceId string      `json:"traceId"`
	Data    interface{} `json:"data"`
}

// @Accept  json
// @Tags HelloService
// @Param id path string true " "
// @Success 200 {object} SwagResponse{data=ent.Pod}
// @Router /hello/{id} [get]
func makeGetUserHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/hello/:id", http.NewServer(endpoints.GetUserEndpoint, decodeGetUserRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeGetUserRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req GetUserReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

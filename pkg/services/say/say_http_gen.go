package say

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

		"SayPod",
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

type HttpHandler struct {
}

func NewHTTPHandler(r *gin.Engine, endpoints Endpoints, options Ops) HttpHandler {

	makeSayPodHandler(r, endpoints, options["SayPod"])

	return HttpHandler{}
}

type SwagResponse struct {
	TraceId string      `json:"traceId"`
	Data    interface{} `json:"data"`
}

// @Accept  json
// @Tags SayService
// @Param id path string true " "
// @Success 200 {object} SwagResponse{data=ent.Pod}
// @Router /saypod/{id} [get]
func makeSayPodHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/saypod/:id", http.NewServer(endpoints.SayPodEndpoint, decodeSayPodRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeSayPodRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req SayPodReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

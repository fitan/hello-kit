package user

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"hello/pkg/ent"

	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"
)

type Ops map[string][]http.ServerOption

func AddHttpOptionToAllMethods(options map[string][]http.ServerOption, option http.ServerOption) {
	methods := []string{

		"GetById",

		"GetList",
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

type HttpHandler struct {
}

func NewHTTPHandler(r *gin.Engine, endpoints Endpoints, options Ops) HttpHandler {

	makeGetByIdHandler(r, endpoints, options["GetById"])

	makeGetListHandler(r, endpoints, options["GetList"])

	return HttpHandler{}
}

type SwagResponse struct {
	TraceId string      `json:"traceId"`
	Data    interface{} `json:"data"`
}

// @Accept  json
// @Tags UserService
// @Param id path string true " "
// @Success 200 {object} SwagResponse{data=ent.User}
// @Router /user/{id} [get]
func makeGetByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/user/:id", http.NewServer(endpoints.GetByIdEndpoint, decodeGetByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeGetByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req GetByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type GetListQuerySwag struct {
	*ent.PodTableServiceNameEQForm
	*ent.PodTableNamespaceEQForm
	*ent.PodTableServiceNameInForm
	*ent.PodTablePagingForm `binding:"required"`
	*ent.PodTableOrderForm
	*ent.PodTableHostIPEQForm
}

type GetListBodySwag ent.User

// @Accept  json
// @Tags UserService
// @Param body body GetListBodySwag true " "
// @Param query query GetListQuerySwag false " "
// @Param project header string false " "
// @Param service header string false " "
// @Success 200 {object} SwagResponse{data=GetListRes}
// @Router /user [get]
func makeGetListHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/user", http.NewServer(endpoints.GetListEndpoint, decodeGetListRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeGetListRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req GetListReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindHeader(&req.Header)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

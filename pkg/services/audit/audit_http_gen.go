package audit

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	debug2 "hello/utils/debug"

	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"

	"hello/pkg/ent"
)

type Ops map[string][]http.ServerOption

func AddHttpOptionToAllMethods(options map[string][]http.ServerOption, option http.ServerOption) {
	methods := []string{

		AuditRestByQueriesAllMethodName,

		AuditRestCreateMethodName,

		AuditRestCreateManyMethodName,

		AuditRestDeleteByIdMethodName,

		AuditRestDeleteManyMethodName,

		AuditRestGetByIdMethodName,

		AuditRestUpdateByIdMethodName,

		AuditRestUpdateManyMethodName,
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

type HttpHandler struct {
}

func NewHTTPHandler(r *gin.Engine, endpoints Endpoints, options Ops, debugSwitch *debug2.DebugSwitch) HttpHandler {

	debugSwitch.Register("AuditRestByQueriesAll", "/audits", "GET")

	debugSwitch.Register("AuditRestCreate", "/audit", "POST")

	debugSwitch.Register("AuditRestCreateMany", "/audits", "POST")

	debugSwitch.Register("AuditRestDeleteById", "/audits/:auditId", "DELETE")

	debugSwitch.Register("AuditRestDeleteMany", "/audits", "DELETE")

	debugSwitch.Register("AuditRestGetById", "/audits/:auditId", "GET")

	debugSwitch.Register("AuditRestUpdateById", "/audits/:auditId", "PUT")

	debugSwitch.Register("AuditRestUpdateMany", "/audits", "PUT")

	makeAuditRestByQueriesAllHandler(r, endpoints, options[AuditRestByQueriesAllMethodName])

	makeAuditRestCreateHandler(r, endpoints, options[AuditRestCreateMethodName])

	makeAuditRestCreateManyHandler(r, endpoints, options[AuditRestCreateManyMethodName])

	makeAuditRestDeleteByIdHandler(r, endpoints, options[AuditRestDeleteByIdMethodName])

	makeAuditRestDeleteManyHandler(r, endpoints, options[AuditRestDeleteManyMethodName])

	makeAuditRestGetByIdHandler(r, endpoints, options[AuditRestGetByIdMethodName])

	makeAuditRestUpdateByIdHandler(r, endpoints, options[AuditRestUpdateByIdMethodName])

	makeAuditRestUpdateManyHandler(r, endpoints, options[AuditRestUpdateManyMethodName])

	return HttpHandler{}
}

type SwagResponse struct {
	TraceId string      `json:"traceId"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type AuditRestByQueriesAllQuerySwag ent.AuditQueryOps

// @Accept  json
// @Tags AuditService
// @Param query query AuditRestByQueriesAllQuerySwag false " "
// @Success 200 {object} SwagResponse{data=ent.AuditRestByQueriesAllRes}
// @Router /audits [get]
func makeAuditRestByQueriesAllHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/audits", http.NewServer(endpoints.AuditRestByQueriesAllEndpoint, decodeAuditRestByQueriesAllRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestByQueriesAllRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestByQueriesAllReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	return req, err
}

type AuditRestCreateBodySwag ent.AuditBaseCreateReq

// @Accept  json
// @Tags AuditService
// @Param body body AuditRestCreateBodySwag true " "
// @Success 200 {object} SwagResponse{data=ent.Audit}
// @Router /audit [post]
func makeAuditRestCreateHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/audit", http.NewServer(endpoints.AuditRestCreateEndpoint, decodeAuditRestCreateRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestCreateRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestCreateReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

type AuditRestCreateManyBodySwag []ent.AuditBaseCreateReq

// @Accept  json
// @Tags AuditService
// @Param body body AuditRestCreateManyBodySwag true " "
// @Success 200 {object} SwagResponse{data=ent.Audits}
// @Router /audits [post]
func makeAuditRestCreateManyHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/audits", http.NewServer(endpoints.AuditRestCreateManyEndpoint, decodeAuditRestCreateManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestCreateManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestCreateManyReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

// @Accept  json
// @Tags AuditService
// @Param auditId path string true " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /audits/{auditId} [delete]
func makeAuditRestDeleteByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/audits/:auditId", http.NewServer(endpoints.AuditRestDeleteByIdEndpoint, decodeAuditRestDeleteByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestDeleteByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestDeleteByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type AuditRestDeleteManyQuerySwag struct {
	Ids []int `json:"ids" form:"ids"`
}

// @Accept  json
// @Tags AuditService
// @Param query query AuditRestDeleteManyQuerySwag false " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /audits [delete]
func makeAuditRestDeleteManyHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/audits", http.NewServer(endpoints.AuditRestDeleteManyEndpoint, decodeAuditRestDeleteManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestDeleteManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestDeleteManyReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	return req, err
}

// @Accept  json
// @Tags AuditService
// @Param auditId path string true " "
// @Success 200 {object} SwagResponse{data=ent.AuditBaseGetRes}
// @Router /audits/{auditId} [get]
func makeAuditRestGetByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/audits/:auditId", http.NewServer(endpoints.AuditRestGetByIdEndpoint, decodeAuditRestGetByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestGetByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestGetByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type AuditRestUpdateByIdBodySwag ent.AuditBaseUpdateReq

// @Accept  json
// @Tags AuditService
// @Param body body AuditRestUpdateByIdBodySwag true " "
// @Param auditId path string true " "
// @Success 200 {object} SwagResponse{data=ent.Audit}
// @Router /audits/{auditId} [put]
func makeAuditRestUpdateByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/audits/:auditId", http.NewServer(endpoints.AuditRestUpdateByIdEndpoint, decodeAuditRestUpdateByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestUpdateByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestUpdateByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

type AuditRestUpdateManyBodySwag []ent.AuditBaseUpdateReq

// @Accept  json
// @Tags AuditService
// @Param body body AuditRestUpdateManyBodySwag true " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /audits [put]
func makeAuditRestUpdateManyHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/audits", http.NewServer(endpoints.AuditRestUpdateManyEndpoint, decodeAuditRestUpdateManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeAuditRestUpdateManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.AuditRestUpdateManyReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}
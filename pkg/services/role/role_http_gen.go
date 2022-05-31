package role

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

		RoleRestAddBindResourcesByRoleIdMethodName,

		RoleRestByQueriesAllMethodName,

		RoleRestCreateMethodName,

		RoleRestCreateManyMethodName,

		RoleRestCreateResourcesByRoleIdMethodName,

		RoleRestDeleteByIdMethodName,

		RoleRestDeleteManyMethodName,

		RoleRestDeleteResourcesByRoleIdMethodName,

		RoleRestGetByIdMethodName,

		RoleRestGetResourcesByRoleIdMethodName,

		RoleRestRemoveBindResourcesByRoleIdMethodName,

		RoleRestUpdateBindResourcesByRoleIdMethodName,

		RoleRestUpdateByIdMethodName,

		RoleRestUpdateManyMethodName,
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

type HttpHandler struct {
}

func NewHTTPHandler(r *gin.Engine, endpoints Endpoints, options Ops, debugSwitch *debug2.DebugSwitch) HttpHandler {

	debugSwitch.Register("RoleRestAddBindResourcesByRoleId", "/roles/:roleId/resources/bind/add", "PUT")

	debugSwitch.Register("RoleRestByQueriesAll", "/roles", "GET")

	debugSwitch.Register("RoleRestCreate", "/role", "POST")

	debugSwitch.Register("RoleRestCreateMany", "/roles", "POST")

	debugSwitch.Register("RoleRestCreateResourcesByRoleId", "/roles/:roleId/resources", "POST")

	debugSwitch.Register("RoleRestDeleteById", "/roles/:roleId", "DELETE")

	debugSwitch.Register("RoleRestDeleteMany", "/roles", "DELETE")

	debugSwitch.Register("RoleRestDeleteResourcesByRoleId", "/roles/:roleId/resources", "DELETE")

	debugSwitch.Register("RoleRestGetById", "/roles/:roleId", "GET")

	debugSwitch.Register("RoleRestGetResourcesByRoleId", "/roles/:roleId/resources", "GET")

	debugSwitch.Register("RoleRestRemoveBindResourcesByRoleId", "/roles/:roleId/resources/bind/remove", "PUT")

	debugSwitch.Register("RoleRestUpdateBindResourcesByRoleId", "/roles/:roleId/resources/bind/update", "PUT")

	debugSwitch.Register("RoleRestUpdateById", "/roles/:roleId", "PUT")

	debugSwitch.Register("RoleRestUpdateMany", "/roles", "PUT")

	makeRoleRestAddBindResourcesByRoleIdHandler(r, endpoints, options[RoleRestAddBindResourcesByRoleIdMethodName])

	makeRoleRestByQueriesAllHandler(r, endpoints, options[RoleRestByQueriesAllMethodName])

	makeRoleRestCreateHandler(r, endpoints, options[RoleRestCreateMethodName])

	makeRoleRestCreateManyHandler(r, endpoints, options[RoleRestCreateManyMethodName])

	makeRoleRestCreateResourcesByRoleIdHandler(r, endpoints, options[RoleRestCreateResourcesByRoleIdMethodName])

	makeRoleRestDeleteByIdHandler(r, endpoints, options[RoleRestDeleteByIdMethodName])

	makeRoleRestDeleteManyHandler(r, endpoints, options[RoleRestDeleteManyMethodName])

	makeRoleRestDeleteResourcesByRoleIdHandler(r, endpoints, options[RoleRestDeleteResourcesByRoleIdMethodName])

	makeRoleRestGetByIdHandler(r, endpoints, options[RoleRestGetByIdMethodName])

	makeRoleRestGetResourcesByRoleIdHandler(r, endpoints, options[RoleRestGetResourcesByRoleIdMethodName])

	makeRoleRestRemoveBindResourcesByRoleIdHandler(r, endpoints, options[RoleRestRemoveBindResourcesByRoleIdMethodName])

	makeRoleRestUpdateBindResourcesByRoleIdHandler(r, endpoints, options[RoleRestUpdateBindResourcesByRoleIdMethodName])

	makeRoleRestUpdateByIdHandler(r, endpoints, options[RoleRestUpdateByIdMethodName])

	makeRoleRestUpdateManyHandler(r, endpoints, options[RoleRestUpdateManyMethodName])

	return HttpHandler{}
}

type SwagResponse struct {
	TraceId string      `json:"traceId"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type RoleRestAddBindResourcesByRoleIdBodySwag struct {
	ResourcesIds []int `json:"resourcesIds"`
}

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestAddBindResourcesByRoleIdBodySwag true " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /roles/{roleId}/resources/bind/add [put]
func makeRoleRestAddBindResourcesByRoleIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/roles/:roleId/resources/bind/add", http.NewServer(endpoints.RoleRestAddBindResourcesByRoleIdEndpoint, decodeRoleRestAddBindResourcesByRoleIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestAddBindResourcesByRoleIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestAddBindResourcesByRoleIdReq
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

type RoleRestByQueriesAllQuerySwag ent.RoleQueryOps

// @Accept  json
// @Tags RoleService
// @Param query query RoleRestByQueriesAllQuerySwag false " "
// @Success 200 {object} SwagResponse{data=ent.RoleRestByQueriesAllRes}
// @Router /roles [get]
func makeRoleRestByQueriesAllHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/roles", http.NewServer(endpoints.RoleRestByQueriesAllEndpoint, decodeRoleRestByQueriesAllRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestByQueriesAllRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestByQueriesAllReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestCreateBodySwag ent.RoleBaseCreateReq

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestCreateBodySwag true " "
// @Success 200 {object} SwagResponse{data=ent.Role}
// @Router /role [post]
func makeRoleRestCreateHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/role", http.NewServer(endpoints.RoleRestCreateEndpoint, decodeRoleRestCreateRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestCreateRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestCreateReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestCreateManyBodySwag []ent.RoleBaseCreateReq

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestCreateManyBodySwag true " "
// @Success 200 {object} SwagResponse{data=ent.Roles}
// @Router /roles [post]
func makeRoleRestCreateManyHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/roles", http.NewServer(endpoints.RoleRestCreateManyEndpoint, decodeRoleRestCreateManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestCreateManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestCreateManyReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestCreateResourcesByRoleIdBodySwag []ent.ResourceBaseCreateReq

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestCreateResourcesByRoleIdBodySwag true " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=ent.Role}
// @Router /roles/{roleId}/resources [post]
func makeRoleRestCreateResourcesByRoleIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/roles/:roleId/resources", http.NewServer(endpoints.RoleRestCreateResourcesByRoleIdEndpoint, decodeRoleRestCreateResourcesByRoleIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestCreateResourcesByRoleIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestCreateResourcesByRoleIdReq
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

// @Accept  json
// @Tags RoleService
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /roles/{roleId} [delete]
func makeRoleRestDeleteByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/roles/:roleId", http.NewServer(endpoints.RoleRestDeleteByIdEndpoint, decodeRoleRestDeleteByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestDeleteByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestDeleteByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestDeleteManyQuerySwag struct {
	Ids []int `json:"ids" form:"ids"`
}

// @Accept  json
// @Tags RoleService
// @Param query query RoleRestDeleteManyQuerySwag false " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /roles [delete]
func makeRoleRestDeleteManyHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/roles", http.NewServer(endpoints.RoleRestDeleteManyEndpoint, decodeRoleRestDeleteManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestDeleteManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestDeleteManyReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestDeleteResourcesByRoleIdQuerySwag struct {
	ResourcesIds []int `json:"resourcesIds" form:"resourcesIds"`
}

// @Accept  json
// @Tags RoleService
// @Param query query RoleRestDeleteResourcesByRoleIdQuerySwag false " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /roles/{roleId}/resources [delete]
func makeRoleRestDeleteResourcesByRoleIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/roles/:roleId/resources", http.NewServer(endpoints.RoleRestDeleteResourcesByRoleIdEndpoint, decodeRoleRestDeleteResourcesByRoleIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestDeleteResourcesByRoleIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestDeleteResourcesByRoleIdReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

// @Accept  json
// @Tags RoleService
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=ent.RoleBaseGetRes}
// @Router /roles/{roleId} [get]
func makeRoleRestGetByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/roles/:roleId", http.NewServer(endpoints.RoleRestGetByIdEndpoint, decodeRoleRestGetByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestGetByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestGetByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestGetResourcesByRoleIdQuerySwag ent.ResourceQueryOps

// @Accept  json
// @Tags RoleService
// @Param query query RoleRestGetResourcesByRoleIdQuerySwag false " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=ent.RoleRestGetResourcesByRoleIdRes}
// @Router /roles/{roleId}/resources [get]
func makeRoleRestGetResourcesByRoleIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/roles/:roleId/resources", http.NewServer(endpoints.RoleRestGetResourcesByRoleIdEndpoint, decodeRoleRestGetResourcesByRoleIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestGetResourcesByRoleIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestGetResourcesByRoleIdReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type RoleRestRemoveBindResourcesByRoleIdBodySwag struct {
	ResourcesIds []int `json:"resourcesIds"`
}

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestRemoveBindResourcesByRoleIdBodySwag true " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /roles/{roleId}/resources/bind/remove [put]
func makeRoleRestRemoveBindResourcesByRoleIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/roles/:roleId/resources/bind/remove", http.NewServer(endpoints.RoleRestRemoveBindResourcesByRoleIdEndpoint, decodeRoleRestRemoveBindResourcesByRoleIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestRemoveBindResourcesByRoleIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestRemoveBindResourcesByRoleIdReq
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

type RoleRestUpdateBindResourcesByRoleIdBodySwag struct {
	OldIds []int `json:"OldIds"`
	NewIds []int `json:"NewIds"`
}

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestUpdateBindResourcesByRoleIdBodySwag true " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /roles/{roleId}/resources/bind/update [put]
func makeRoleRestUpdateBindResourcesByRoleIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/roles/:roleId/resources/bind/update", http.NewServer(endpoints.RoleRestUpdateBindResourcesByRoleIdEndpoint, decodeRoleRestUpdateBindResourcesByRoleIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestUpdateBindResourcesByRoleIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestUpdateBindResourcesByRoleIdReq
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

type RoleRestUpdateByIdBodySwag ent.RoleBaseUpdateReq

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestUpdateByIdBodySwag true " "
// @Param roleId path string true " "
// @Success 200 {object} SwagResponse{data=ent.Role}
// @Router /roles/{roleId} [put]
func makeRoleRestUpdateByIdHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/roles/:roleId", http.NewServer(endpoints.RoleRestUpdateByIdEndpoint, decodeRoleRestUpdateByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestUpdateByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestUpdateByIdReq
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

type RoleRestUpdateManyBodySwag []ent.RoleBaseUpdateReq

// @Accept  json
// @Tags RoleService
// @Param body body RoleRestUpdateManyBodySwag true " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /roles [put]
func makeRoleRestUpdateManyHandler(r *gin.Engine, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/roles", http.NewServer(endpoints.RoleRestUpdateManyEndpoint, decodeRoleRestUpdateManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeRoleRestUpdateManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.RoleRestUpdateManyReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

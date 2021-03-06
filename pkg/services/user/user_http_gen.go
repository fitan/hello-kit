package user

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"hello/utils/debug"

	"github.com/fitan/gink/transport/http"
	"github.com/gin-gonic/gin"

	"hello/pkg/ent"
)

type Ops map[string][]http.ServerOption

func AddHttpOptionToAllMethods(options map[string][]http.ServerOption, option http.ServerOption) {
	methods := []string{

		UserRestAddBindRolesByUserIdMethodName,

		UserRestByQueriesAllMethodName,

		UserRestCreateMethodName,

		UserRestCreateManyMethodName,

		UserRestCreateRolesByUserIdMethodName,

		UserRestDeleteByIdMethodName,

		UserRestDeleteManyMethodName,

		UserRestDeleteRolesByUserIdMethodName,

		UserRestGetByIdMethodName,

		UserRestGetRolesByUserIdMethodName,

		UserRestRemoveBindRolesByUserIdMethodName,

		UserRestUpdateBindRolesByUserIdMethodName,

		UserRestUpdateByIdMethodName,

		UserRestUpdateManyMethodName,
	}
	for _, v := range methods {
		options[v] = append(options[v], option)
	}
}

type HttpHandler struct {
}

func NewHTTPHandler(r *gin.RouterGroup, endpoints Endpoints, options Ops, debugSwitch *debug.DebugSwitch) HttpHandler {

	debugSwitch.Register("UserRestAddBindRolesByUserId", r.BasePath()+"/users/:userId/roles/bind/add", "PUT")

	debugSwitch.Register("UserRestByQueriesAll", r.BasePath()+"/users", "GET")

	debugSwitch.Register("UserRestCreate", r.BasePath()+"/user", "POST")

	debugSwitch.Register("UserRestCreateMany", r.BasePath()+"/users", "POST")

	debugSwitch.Register("UserRestCreateRolesByUserId", r.BasePath()+"/users/:userId/roles", "POST")

	debugSwitch.Register("UserRestDeleteById", r.BasePath()+"/users/:userId", "DELETE")

	debugSwitch.Register("UserRestDeleteMany", r.BasePath()+"/users", "DELETE")

	debugSwitch.Register("UserRestDeleteRolesByUserId", r.BasePath()+"/users/:userId/roles", "DELETE")

	debugSwitch.Register("UserRestGetById", r.BasePath()+"/users/:userId", "GET")

	debugSwitch.Register("UserRestGetRolesByUserId", r.BasePath()+"/users/:userId/roles", "GET")

	debugSwitch.Register("UserRestRemoveBindRolesByUserId", r.BasePath()+"/users/:userId/roles/bind/remove", "PUT")

	debugSwitch.Register("UserRestUpdateBindRolesByUserId", r.BasePath()+"/users/:userId/roles/bind/update", "PUT")

	debugSwitch.Register("UserRestUpdateById", r.BasePath()+"/users/:userId", "PUT")

	debugSwitch.Register("UserRestUpdateMany", r.BasePath()+"/users", "PUT")

	makeUserRestAddBindRolesByUserIdHandler(r, endpoints, options[UserRestAddBindRolesByUserIdMethodName])

	makeUserRestByQueriesAllHandler(r, endpoints, options[UserRestByQueriesAllMethodName])

	makeUserRestCreateHandler(r, endpoints, options[UserRestCreateMethodName])

	makeUserRestCreateManyHandler(r, endpoints, options[UserRestCreateManyMethodName])

	makeUserRestCreateRolesByUserIdHandler(r, endpoints, options[UserRestCreateRolesByUserIdMethodName])

	makeUserRestDeleteByIdHandler(r, endpoints, options[UserRestDeleteByIdMethodName])

	makeUserRestDeleteManyHandler(r, endpoints, options[UserRestDeleteManyMethodName])

	makeUserRestDeleteRolesByUserIdHandler(r, endpoints, options[UserRestDeleteRolesByUserIdMethodName])

	makeUserRestGetByIdHandler(r, endpoints, options[UserRestGetByIdMethodName])

	makeUserRestGetRolesByUserIdHandler(r, endpoints, options[UserRestGetRolesByUserIdMethodName])

	makeUserRestRemoveBindRolesByUserIdHandler(r, endpoints, options[UserRestRemoveBindRolesByUserIdMethodName])

	makeUserRestUpdateBindRolesByUserIdHandler(r, endpoints, options[UserRestUpdateBindRolesByUserIdMethodName])

	makeUserRestUpdateByIdHandler(r, endpoints, options[UserRestUpdateByIdMethodName])

	makeUserRestUpdateManyHandler(r, endpoints, options[UserRestUpdateManyMethodName])

	return HttpHandler{}
}

type SwagResponse struct {
	TraceId string      `json:"traceId"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type UserRestAddBindRolesByUserIdBodySwag struct {
	RolesIds []int `json:"rolesIds"`
}

// @Accept  json
// @Tags UserService
// @Param body body UserRestAddBindRolesByUserIdBodySwag true " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /users/{userId}/roles/bind/add [put]
func makeUserRestAddBindRolesByUserIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/users/:userId/roles/bind/add", http.NewServer(endpoints.UserRestAddBindRolesByUserIdEndpoint, decodeUserRestAddBindRolesByUserIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestAddBindRolesByUserIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestAddBindRolesByUserIdReq
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

type UserRestByQueriesAllQuerySwag ent.UserQueryOps

// @Accept  json
// @Tags UserService
// @Param query query UserRestByQueriesAllQuerySwag false " "
// @Success 200 {object} SwagResponse{data=ent.UserRestByQueriesAllRes}
// @Router /users [get]
func makeUserRestByQueriesAllHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/users", http.NewServer(endpoints.UserRestByQueriesAllEndpoint, decodeUserRestByQueriesAllRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestByQueriesAllRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestByQueriesAllReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	return req, err
}

type UserRestCreateBodySwag ent.UserBaseCreateReq

// @Accept  json
// @Tags UserService
// @Param body body UserRestCreateBodySwag true " "
// @Success 200 {object} SwagResponse{data=ent.User}
// @Router /user [post]
func makeUserRestCreateHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/user", http.NewServer(endpoints.UserRestCreateEndpoint, decodeUserRestCreateRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestCreateRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestCreateReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

type UserRestCreateManyBodySwag []ent.UserBaseCreateReq

// @Accept  json
// @Tags UserService
// @Param body body UserRestCreateManyBodySwag true " "
// @Success 200 {object} SwagResponse{data=ent.Users}
// @Router /users [post]
func makeUserRestCreateManyHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/users", http.NewServer(endpoints.UserRestCreateManyEndpoint, decodeUserRestCreateManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestCreateManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestCreateManyReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

type UserRestCreateRolesByUserIdBodySwag []ent.RoleBaseCreateReq

// @Accept  json
// @Tags UserService
// @Param body body UserRestCreateRolesByUserIdBodySwag true " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=ent.User}
// @Router /users/{userId}/roles [post]
func makeUserRestCreateRolesByUserIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.POST("/users/:userId/roles", http.NewServer(endpoints.UserRestCreateRolesByUserIdEndpoint, decodeUserRestCreateRolesByUserIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestCreateRolesByUserIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestCreateRolesByUserIdReq
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
// @Tags UserService
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /users/{userId} [delete]
func makeUserRestDeleteByIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/users/:userId", http.NewServer(endpoints.UserRestDeleteByIdEndpoint, decodeUserRestDeleteByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestDeleteByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestDeleteByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type UserRestDeleteManyQuerySwag struct {
	Ids []int `json:"ids" form:"ids"`
}

// @Accept  json
// @Tags UserService
// @Param query query UserRestDeleteManyQuerySwag false " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /users [delete]
func makeUserRestDeleteManyHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/users", http.NewServer(endpoints.UserRestDeleteManyEndpoint, decodeUserRestDeleteManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestDeleteManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestDeleteManyReq
	var err error

	err = ctx.ShouldBindQuery(&req.Query)
	if err != nil {
		return nil, err
	}

	return req, err
}

type UserRestDeleteRolesByUserIdQuerySwag struct {
	RolesIds []int `json:"rolesIds" form:"rolesIds"`
}

// @Accept  json
// @Tags UserService
// @Param query query UserRestDeleteRolesByUserIdQuerySwag false " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /users/{userId}/roles [delete]
func makeUserRestDeleteRolesByUserIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.DELETE("/users/:userId/roles", http.NewServer(endpoints.UserRestDeleteRolesByUserIdEndpoint, decodeUserRestDeleteRolesByUserIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestDeleteRolesByUserIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestDeleteRolesByUserIdReq
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
// @Tags UserService
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=ent.UserBaseGetRes}
// @Router /users/{userId} [get]
func makeUserRestGetByIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/users/:userId", http.NewServer(endpoints.UserRestGetByIdEndpoint, decodeUserRestGetByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestGetByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestGetByIdReq
	var err error

	err = ctx.ShouldBindUri(&req.Uri)
	if err != nil {
		return nil, err
	}

	return req, err
}

type UserRestGetRolesByUserIdQuerySwag ent.RoleQueryOps

// @Accept  json
// @Tags UserService
// @Param query query UserRestGetRolesByUserIdQuerySwag false " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=ent.UserRestGetRolesByUserIdRes}
// @Router /users/{userId}/roles [get]
func makeUserRestGetRolesByUserIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.GET("/users/:userId/roles", http.NewServer(endpoints.UserRestGetRolesByUserIdEndpoint, decodeUserRestGetRolesByUserIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestGetRolesByUserIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestGetRolesByUserIdReq
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

type UserRestRemoveBindRolesByUserIdBodySwag struct {
	RolesIds []int `json:"rolesIds"`
}

// @Accept  json
// @Tags UserService
// @Param body body UserRestRemoveBindRolesByUserIdBodySwag true " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /users/{userId}/roles/bind/remove [put]
func makeUserRestRemoveBindRolesByUserIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/users/:userId/roles/bind/remove", http.NewServer(endpoints.UserRestRemoveBindRolesByUserIdEndpoint, decodeUserRestRemoveBindRolesByUserIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestRemoveBindRolesByUserIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestRemoveBindRolesByUserIdReq
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

type UserRestUpdateBindRolesByUserIdBodySwag struct {
	OldIds []int `json:"OldIds"`
	NewIds []int `json:"NewIds"`
}

// @Accept  json
// @Tags UserService
// @Param body body UserRestUpdateBindRolesByUserIdBodySwag true " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=string}
// @Router /users/{userId}/roles/bind/update [put]
func makeUserRestUpdateBindRolesByUserIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/users/:userId/roles/bind/update", http.NewServer(endpoints.UserRestUpdateBindRolesByUserIdEndpoint, decodeUserRestUpdateBindRolesByUserIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestUpdateBindRolesByUserIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestUpdateBindRolesByUserIdReq
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

type UserRestUpdateByIdBodySwag ent.UserBaseUpdateReq

// @Accept  json
// @Tags UserService
// @Param body body UserRestUpdateByIdBodySwag true " "
// @Param userId path string true " "
// @Success 200 {object} SwagResponse{data=ent.User}
// @Router /users/{userId} [put]
func makeUserRestUpdateByIdHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/users/:userId", http.NewServer(endpoints.UserRestUpdateByIdEndpoint, decodeUserRestUpdateByIdRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestUpdateByIdRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestUpdateByIdReq
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

type UserRestUpdateManyBodySwag []ent.UserBaseUpdateReq

// @Accept  json
// @Tags UserService
// @Param body body UserRestUpdateManyBodySwag true " "
// @Success 200 {object} SwagResponse{data=bool}
// @Router /users [put]
func makeUserRestUpdateManyHandler(r *gin.RouterGroup, endpoints Endpoints, options []http.ServerOption) {
	r.PUT("/users", http.NewServer(endpoints.UserRestUpdateManyEndpoint, decodeUserRestUpdateManyRequest, http.EncodeJSONResponse, options...).ServeHTTP)
}

func decodeUserRestUpdateManyRequest(_ context.Context, ctx *gin.Context) (interface{}, error) {
	var req ent.UserRestUpdateManyReq
	var err error

	err = ctx.ShouldBindJSON(&req.Body)
	if err != nil {
		return nil, err
	}

	return req, err
}

package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UrlRequestContext() func(ctx context.Context, r *gin.Context) context.Context {
	return func(ctx context.Context, r *gin.Context) context.Context {
		projectIdStr := r.Param(ContextKeyProjectId)
		serviceIdStr := r.Param(ContextKeyServiceId)
		var projectId int
		var serviceId int

		projectId, _ = strconv.Atoi(projectIdStr)
		serviceId, _ = strconv.Atoi(serviceIdStr)
		ctx = context.WithValue(ctx, ContextKeyProjectId, projectId)
		ctx = context.WithValue(ctx, ContextKeyServiceId, serviceId)

		return ctx
	}

}

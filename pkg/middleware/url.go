package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UrlRequestContext() func(ctx context.Context, r *gin.Context) context.Context {
	return func(ctx context.Context, r *gin.Context) context.Context {
		projectIdStr, ok := r.Params.Get(ContextKeyProjectId)
		if ok {
			var projectId int
			projectId, _ = strconv.Atoi(projectIdStr)
			ctx = context.WithValue(ctx, ContextKeyProjectId, projectId)
		}

		serviceIdStr, ok := r.Params.Get(ContextKeyServiceId)
		if ok {
			var serviceId int
			serviceId, _ = strconv.Atoi(serviceIdStr)
			ctx = context.WithValue(ctx, ContextKeyServiceId, serviceId)
		}

		return ctx
	}

}

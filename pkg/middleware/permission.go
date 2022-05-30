package middleware

import (
	"context"
	"fmt"
	ginkHttp "github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/pkg/errors"
	"hello/pkg/ent"
	resource2 "hello/pkg/ent/resource"
	"hello/pkg/repository"
	"hello/pkg/services"
	"hello/pkg/services/casbin"
)

func PermissionMiddleware(services *services.Services, repository repository.Repository) endpoint.Middleware {
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			projectId, projectOK := ctx.Value(ContextKeyProjectId).(int)
			serviceId, serviceOK := ctx.Value(ContextKeyServiceId).(int)
			UserId, _ := ctx.Value("userId").(int)
			if UserId == 0 {
				UserId = 99999999
			}
			//if !userOk {
			//	err = fmt.Errorf("not logged in")
			//	return
			//}
			if !projectOK && !serviceOK {
				return e(ctx, request)
			}

			var domain string

			if projectOK && serviceOK {
				domain = fmt.Sprintf("/%v/%v", projectId, serviceId)
			}

			if !projectOK {
				p, err := repository.Service.GetProjectByServiceId(ctx, serviceId)
				if err != nil {
					err = errors.Wrap(err, "Service.GetProjectByServiceId")
					return nil, err
				}
				projectId = p.ID
				domain = fmt.Sprintf("/%v/%v", projectId, serviceId)
			}

			if !serviceOK {
				domain = fmt.Sprintf("/%v/*", projectId)
			}

			method, _ := ctx.Value(ginkHttp.ContextKeyRequestMethod).(string)
			path, _ := ctx.Value(ginkHttp.ContextKeyRequestFullPath).(string)
			t := resource2.TypeAPI
			one, err := repository.Resource.ByQueriesOne(
				ctx, struct {
					ent.ResourceTableActionEQForm
					ent.ResourceTableTypeEQForm
					ent.ResourceTablePathEQForm
				}{
					ResourceTableActionEQForm: ent.ResourceTableActionEQForm{ActionEQ: &method},
					ResourceTableTypeEQForm:   ent.ResourceTableTypeEQForm{TypeEQ: &t},
					ResourceTablePathEQForm:   ent.ResourceTablePathEQForm{PathEQ: &path},
				},
			)
			if err != nil {
				return nil, errors.Wrap(err, "Resource.ByQueriesOne")
			}

			permission, err := services.Casbin.CheckPermission(
				ctx, casbin.CheckPermission{
					User:       UserId,
					Domain:     domain,
					ResourceId: one.ID,
				},
			)
			if err != nil {
				err = errors.Wrap(err, "Casbin.CheckPermission")
				return nil, err
			}

			if !permission {
				err = fmt.Errorf("do not have permission")
				return
			}

			return e(ctx, request)
		}
	}
}

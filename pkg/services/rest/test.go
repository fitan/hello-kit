package rest

import (
	"context"
	"hello/pkg/ent"
)

type UserRest interface {
	Create(ctx context.Context, req ent.UserRestCreateReq) (res *ent.User, err error)
	CreateMany(ctx context.Context, req ent.UserRestCreateManyReq) (res ent.Users, err error)
	GetById(ctx context.Context, req ent.UserRestGetByIdReq) (res *ent.User, err error)
	ByQueries(ctx context.Context, req ent.UserRestByQueriesReq) (res ent.Users, count int, err error)
	UpdateById(ctx context.Context, req ent.UserRestUpdateByIdReq) (res *ent.User, err error)
	UpdateMany(ctx context.Context, req ent.UserRestUpdateManyReq) (success bool, err error)
	DeleteById(ctx context.Context, req ent.UserRestDeleteByIdReq) (success bool, err error)
	DeleteMany(ctx context.Context, req ent.UserRestDeleteManyReq) (success bool, err error)
	CreatePodsSliceByUserId(ctx context.Context, req ent.UserRestCreatePodsSliceByUserIdReq) (res *ent.User, err error)
	GetPodsSliceByUserId(ctx context.Context, req ent.UserRestGetPodsSliceByUserIdReq) (
		res ent.Pods, count int, err error,
	)
}

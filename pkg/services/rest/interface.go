package rest

import (
	"context"
	"hello/pkg/ent"
)

type RestService interface {
	CreateUser(ctx context.Context, req CreateUserReq) (res CreateUserRes, err error)
	CreateUsers(ctx context.Context, req CreateUsersReq) (res CreateUsersRes, err error)
	UpdateUserById(ctx context.Context, req UpdateUserIdReq) (res UpdateUserIdRes, err error)
	UpdateUsers(ctx context.Context, req UpdateUsersReq) (res UpdateUsersRes, err error)
	ReadUserById(ctx context.Context, req ReadUserByIdReq) (res ReadUserByIdRes, err error)
	ReadUsers(ctx context.Context, req ReadUsersReq) (res ReadUsersRes, err error)
	DeleteUserById(ctx context.Context, req DeleteUserByIdReq) (res DeleteUserByIdRes, err error)
	DeleteUsers(ctx context.Context, req DeleteUsersReq) (res DeleteUsersRes, err error)
}

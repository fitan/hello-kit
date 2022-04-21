package rest

import (
	"context"
	"hello/pkg/ent"
)

// RestService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i RestService -bt "prometheus:rest_with_prometheus_gen.go log:rest_with_log_gen.go opentracing:rest_with_trace_gen.go"

type basicRestService struct {
	db *ent.Client
	userRepo ent.UserBaseInterface
}

func CreateCopy(o *ent.UserCreate,v ent.User) {
	o.SetAge(v.Age).SetName(v.Name)
}

func ReadCopy(q *ent.UserQuery) {
	q.Select()
}

func UpdateCopy(o *ent.UserUpdate,v ent.User) {
	o.SetAge(v.Age)
	o.SetName(v.Name)
}

func ResultOneCopy(user *ent.User) (res ent.User) {
	return
}

func ResultListCopy(users ent.Users) (res ent.Users) {
	return
}


type CreateUserReq struct {
	Body ent.User
}

func (b *basicRestService) CreateUser(ctx context.Context, req CreateUserReq) (ent.User, error) {
	c := b.db.User.Create()
	CreateCopy(c, req.Body)
	data, err := c.Save(ctx)
	return ResultOneCopy(data), err
}


type CreateUsersReq struct {
	Body ent.Users
}

type CreateUsersRes struct {

}

func (b *basicRestService) CreateUsers(ctx context.Context, req CreateUsersReq) (res CreateUsersRes, err error) {
	b.userRepo.CreateMany(ctx, req.Body)

}

type UpdateUserIdReq struct {

}

type UpdateUserIdRes struct {

}

func (b *basicRestService) UpdateUserById(ctx context.Context, req UpdateUserIdReq) (res UpdateUserIdRes, err error) {
	panic("implement me")
}

type UpdateUsersReq struct {

}

type UpdateUsersRes struct {

}

func (b *basicRestService) UpdateUsers(ctx context.Context, req UpdateUsersReq) (res UpdateUsersRes, err error) {
	panic("implement me")
}

type ReadUserByIdReq struct {

}

type ReadUserByIdRes struct {

}

func (b *basicRestService) ReadUserById(ctx context.Context, req ReadUserByIdReq) (res ReadUserByIdRes, err error) {
	panic("implement me")
}

type ReadUsersReq struct {

}

type ReadUsersRes struct {

}

func (b *basicRestService) ReadUsers(ctx context.Context, req ReadUsersReq) (res ReadUsersRes, err error) {
	panic("implement me")
}

type DeleteUserByIdReq struct {

}

type DeleteUserByIdRes struct {

}

func (b *basicRestService) DeleteUserById(ctx context.Context, req DeleteUserByIdReq) (res DeleteUserByIdRes, err error) {
	panic("implement me")
}

type DeleteUsersReq struct {

}

type DeleteUsersRes struct {

}

func (b *basicRestService) DeleteUsers(ctx context.Context, req DeleteUsersReq) (res DeleteUsersRes, err error) {
	panic("implement me")
}

type BaseService RestService

// NewBasicRestService returns a naive, stateless implementation of RestService.
func NewBasicService() BaseService {
	return &basicRestService{}
}

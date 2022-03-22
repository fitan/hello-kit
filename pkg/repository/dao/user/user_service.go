package user

import (
	"context"
	"hello/pkg/repository/dao/ent"
)

// UserService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i UserService -bt "prometheus:user_with_prometheus_gen.go log:user_with_log_gen.go opentracing:user_with_trace_gen.go http-gin:user_http_gen.go endpoint:user_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i UserService -bt "prometheus:user_with_prometheus_gen.go log:user_with_log_gen.go opentracing:user_with_trace_gen.go"
type UserService interface {
	GetById(ctx context.Context, id int) (*ent.User, error)
	DeleteById(ctx context.Context, id int) error
	GetList(ctx context.Context) ([]*ent.User, error)
	Update(ctx context.Context, u *ent.User) (*ent.User, error)
}

type basicUserService struct {
	client *ent.Client
}

func (b *basicUserService) Update(ctx context.Context, u *ent.User) (*ent.User, error) {
	return b.client.User.UpdateOne(u).Save(ctx)
}

func (b *basicUserService) DeleteById(ctx context.Context, id int) error {
	return b.client.User.DeleteOneID(id).Exec(ctx)
}

func (b *basicUserService) GetById(ctx context.Context, id int) (*ent.User, error) {
	return b.client.User.Get(ctx, id)
}

func (b *basicUserService) GetList(ctx context.Context) ([]*ent.User, error) {
	return b.client.User.Query().All(ctx)
}

type BaseService UserService

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicService(client *ent.Client) BaseService {
	return &basicUserService{client: client}
}

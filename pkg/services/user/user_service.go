package user

import (
	"context"
	"hello/pkg/ent"
	"hello/pkg/repository"
	"hello/pkg/services/casbin"
)

// UserService describes the service.
//go:generate gowrap gen -g -p ./ -i UserService -bt "prometheus:user_with_prometheus_gen.go log:user_with_log_gen.go opentracing:user_with_trace_gen.go http-gin:user_http_gen.go endpoint:user_endpoint_gen.go"
type UserService interface {
	// @http-gin /user/:id GET
	GetById(ctx context.Context, req GetByIdReq) (*ent.User, error)
	// @http-gin /user GET
	GetList(ctx context.Context, req GetListReq) ([]*ent.User, error)
}

type basicUserService struct {
	repo   repository.Repository
	casbin casbin.CasbinService
	db     *ent.Client
}

type GetByIdReq struct {
	Uri struct {
		Id int `uri:"id" json:"id"`
	} `json:"uri"`
}

func (s *basicUserService) GetById(ctx context.Context, req GetByIdReq) (*ent.User, error) {
	return s.repo.User.GetById(ctx, req.Uri.Id)
}

type GetListReq struct {
	Query struct {
		say string
		ent.UserTableNameInForm
		ent.UserTablePagingForm
		ent.UserTableOrderForm
	} `json:"query"`
}

func (s *basicUserService) GetList(ctx context.Context, req GetListReq) ([]*ent.User, error) {
	q := s.db.User.Query()
	ent.SetUserFormQueries(req, q)
	return q.All(ctx)
}

type BaseService UserService

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicService(repo repository.Repository, casbin casbin.CasbinService, db *ent.Client) BaseService {
	return &basicUserService{
		repo:   repo,
		casbin: casbin,
		db:     db,
	}
}

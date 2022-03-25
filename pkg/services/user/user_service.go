package user

import (
	"context"
	"fmt"
	"hello/pkg/ent"
	"hello/pkg/repository"
	"hello/pkg/services/casbin"
	"reflect"
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
	} `json:"query"`
}

func (s *basicUserService) GetList(ctx context.Context, req GetListReq) ([]*ent.User, error) {
	q := s.db.User.Query()
	req.Query.UserTableNameInForm.Query(q)
	req.Query.UserTablePagingForm.Query(q)
	fmt.Println("all: ",GetUserQueries(req))
	return q.All(ctx)
}


func GetUserQueries(o interface{}) []ent.UserTableFormer {
	l := make([]ent.UserTableFormer,0)
	v := reflect.ValueOf(o)
	t := reflect.TypeOf(o)
	former := reflect.TypeOf((*ent.UserTableFormer)(nil)).Elem()
	DepValue(v, t,former, &l)
	return l
}

func DepValue(v reflect.Value, t reflect.Type,former reflect.Type, l *[]ent.UserTableFormer)  {


	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.IsZero() {
			continue
		}
		fmt.Printf("f itnerface %#v\n former %#v\n", f.Interface(), former)
		if f.Type().Implements(former) {
			fmt.Println("find: usertableformer")
			*l = append(*l, f.Interface().(ent.UserTableFormer))
			continue
		}
		fmt.Printf("types kind: %s", f.Type().Kind())
		if f.Type().Kind() == reflect.Struct {
			fmt.Println("types kind is struct")
			DepValue(f, t.Field(i).Type,former, l)
		}
	}
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

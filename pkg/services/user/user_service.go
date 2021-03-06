package user

import (
	"github.com/nats-io/nats.go"
	"hello/pkg/ent"
	"hello/pkg/repository"
	"hello/pkg/services/casbin"
)

// UserService describes the service.
//go:generate gowrap gen -g -p ./ -i UserService -bt "prometheus:user_with_prometheus_gen.go log:user_with_log_gen.go opentracing:user_with_trace_gen.go http-gin:user_http_gen.go endpoint:user_endpoint_gen.go"
type UserService interface {
	ent.UserRestInterface
}

type basicUserService struct {
	repo   repository.Repository
	casbin casbin.CasbinService
	db     *ent.Client
	ent.UserRestInterface
	natsConn *nats.Conn
}

type NatsGetReq struct {
}

//type GetByIdReq struct {
//	Uri struct {
//		Id int `uri:"id" json:"id"`
//	} `json:"uri"`
//}
//
//func (s *basicUserService) GetById(ctx context.Context, req GetByIdReq) (*ent.User, error) {
//	return s.repo.User.GetById(ctx, req.Uri.Id)
//}
//
//type GetListReq struct {
//	Query struct {
//		*ent.PodTableServiceNameEQForm
//		*ent.PodTableNamespaceEQForm
//		*ent.PodTableServiceNameInForm
//		*ent.PodTablePagingForm `binding:"required"`
//		*ent.PodTableOrderForm
//		*ent.PodTableHostIPEQForm
//	} `json:"query"`
//}
//
//type GetListRes struct {
//	Count int `json:"count"`
//	List ent.Pods `json:"data"`
//}
//
//func (s *basicUserService) GetList(ctx context.Context, req GetListReq) (res GetListRes,err error) {
//	res.List, res.Count, err = s.db.Pod.Query().ByQueries(ctx, req)
//	return
//	ent.SetUserFormQueries(req, q)
//}
//
type BaseService UserService

//
// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicService(repo repository.Repository, casbin casbin.CasbinService, db *ent.Client) BaseService {
	return &basicUserService{
		repo:              repo,
		casbin:            casbin,
		db:                db,
		UserRestInterface: ent.NewUserRest(db),
	}
}

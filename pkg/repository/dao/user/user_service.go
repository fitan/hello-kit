package user

import (
	"hello/pkg/ent"
)

// UserService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i UserService -bt "prometheus:user_with_prometheus_gen.go log:user_with_log_gen.go opentracing:user_with_trace_gen.go http-gin:user_http_gen.go endpoint:user_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i UserService -bt "prometheus:user_with_prometheus_gen.go log:user_with_log_gen.go opentracing:user_with_trace_gen.go"
type UserService interface {
	ent.UserBaseInterface
}

type basicUserService struct {
	client *ent.Client
	ent.UserBaseInterface
}



type BaseService UserService

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicService(client *ent.Client) BaseService {
	return &basicUserService{client: client, UserBaseInterface: ent.NewUserBase(client.User)}
}

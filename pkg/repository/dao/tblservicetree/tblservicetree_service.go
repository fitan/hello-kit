package tblservicetree

import "hello/pkg/ent"

// TblservicetreeService describes the service.
// if use kit. go:generate gowrap gen -g -p ./ -i TblservicetreeService -bt "prometheus:tblservicetree_with_prometheus_gen.go log:tblservicetree_with_log_gen.go opentracing:tblservicetree_with_trace_gen.go http-gin:tblservicetree_http_gen.go endpoint:tblservicetree_endpoint_gen.go"
//go:generate gowrap gen -ps=false -g -p ./ -i TblservicetreeService -bt "prometheus:tblservicetree_with_prometheus_gen.go log:tblservicetree_with_log_gen.go opentracing:tblservicetree_with_trace_gen.go"
type TblservicetreeService interface {
	ent.SpiderDevTblServicetreeBaseInterface
}

type basicTblservicetreeService struct {
	ent.SpiderDevTblServicetreeBaseInterface
}

type BaseService TblservicetreeService

// NewBasicTblservicetreeService returns a naive, stateless implementation of TblservicetreeService.
func NewBasicService(client *ent.Client) BaseService {
	return &basicTblservicetreeService{SpiderDevTblServicetreeBaseInterface:ent.NewSpiderDevTblServicetreeBase(client)}
}

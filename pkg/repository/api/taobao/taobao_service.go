package taobao

import (
	"context"
	"github.com/go-resty/resty/v2"
	"hello/utils/conf"
)

// TaobaoService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i TaobaoService -bt "prometheus:taobao_with_prometheus_gen.go log:taobao_with_log_gen.go opentracing:taobao_with_trace_gen.go"
type TaobaoService interface {
	GetRoot(ctx context.Context) string
}

type basicTaobaoService struct {
	client *resty.Client
}

func (b *basicTaobaoService) GetRoot(ctx context.Context) string {
	get, err := b.client.R().Get("/")
	if err != nil {
		return ""
	}
	return get.String()
}

type BaseService TaobaoService

// NewBasicTaobaoService returns a naive, stateless implementation of TaobaoService.
func NewBasicService(conf *conf.MyConf) BaseService {
	return &basicTaobaoService{client: resty.New().SetBaseURL(conf.Apis.Taobao.Url)}
}

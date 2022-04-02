package baidu

import (
	"context"
	"github.com/go-resty/resty/v2"
	"hello/utils/conf"
)

// BaiduService describes the service.
//go:generate gowrap gen -ps=false -g -p ./ -i BaiduService -bt "prometheus:baidu_with_prometheus_gen.go log:baidu_with_log_gen.go opentracing:baidu_with_trace_gen.go"
type BaiduService interface {
	GetRoot(ctx context.Context) string
}

type basicBaiduService struct {
	client *resty.Client
}

func (b *basicBaiduService) GetRoot(ctx context.Context) string {
	get, err := b.client.R().Get("/")
	if err != nil {
		return ""
	}
	return get.String()
}

type BaseService BaiduService

// NewBasicBaiduService returns a naive, stateless implementation of BaiduService.
func NewBasicService(conf *conf.MyConf) BaseService {
	return &basicBaiduService{client: resty.New().SetBaseURL(conf.Apis.Baidu.Url)}
}

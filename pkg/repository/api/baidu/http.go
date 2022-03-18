package baidu

import (
	"context"
	"github.com/go-resty/resty/v2"
	"hello/utils/conf"
	"hello/utils/httpclient"
)

//go:generate gowrap gen -ps  -g -p ./ -i Http -bt "../../../gowrap/templates/prometheus.tmpl:http_with_prometheus_gen.go ../../../gowrap/templates/log.tmpl:http_with_log_gen.go ../../../gowrap/templates/opentracing.tmpl:http_with_trace_gen.go"
type Http interface {
	GetRoot(ctx context.Context) (*resty.Response, error)
	GetRoot1(ctx context.Context) (*resty.Response, error)
}

type http struct {
	client *resty.Client
}

func (a *http) GetRoot(ctx context.Context) (*resty.Response, error) {
	return a.client.R().Get("/")
}

func (a *http) GetRoot1(ctx context.Context) (*resty.Response, error) {
	return a.client.R().Get("/1")
}

type BaseHttp Http

func NewBaseHttp(conf *conf.MyConf) BaseHttp {
	return &http{client: httpclient.NewHttpClient(httpclient.WithDebug(conf.Apis.Baidu.TraceDebug), httpclient.WithHost(conf.Apis.Baidu.Url))}
}

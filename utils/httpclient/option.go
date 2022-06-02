package httpclient

import (
	"github.com/go-resty/resty/v2"
	"go-micro.dev/v4/registry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
	"net"
	"net/http"
	"runtime"
	"time"
)

type Option func(client *resty.Client)

type NewFunc func(fs ...Option) *resty.Client

func New(fs ...Option) *resty.Client {
	client := resty.New()
	for _, f := range fs {
		f(client)
	}

	return client
}

func WithTransport(transport *http.Transport) Option {
	if transport == nil {
		transport = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConnsPerHost:   10,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}
	return func(client *resty.Client) {
		client.SetTransport(transport)
	}
}

func WithHost(host string) Option {
	return func(client *resty.Client) {
		client.SetBaseURL(host)
	}
}

func WithMicroHost(serviceName string, r registry.Registry) Option {
	return func(client *resty.Client) {
		client.OnBeforeRequest(BeforeMicroSelect(serviceName, r))
	}
}

func WithDebug(debug bool) Option {
	return func(client *resty.Client) {
		client.SetDebug(debug)
	}
}

func WithTimeOut(timeOut time.Duration) Option {
	return func(client *resty.Client) {
		client.SetTimeout(timeOut)
	}
}

func WithRetry(retryCount int, retryWaitTime, retryMaxWaitTime time.Duration) Option {
	return func(client *resty.Client) {
		client.SetRetryCount(retryCount)
		client.SetRetryWaitTime(retryWaitTime)
		client.SetRetryMaxWaitTime(retryMaxWaitTime)
	}
}

func WithDebugMid(log *zap.SugaredLogger, bodyLimit int) Option {
	return func(client *resty.Client) {
		client.OnBeforeRequest(BeforeDebug(log, bodyLimit))
		client.OnAfterResponse(AfterDebug(log, int64(bodyLimit)))
	}
}

func WithTrace() Option {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}

	t := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
	return func(client *resty.Client) {
		client.SetTransport(otelhttp.NewTransport(t))
	}
}

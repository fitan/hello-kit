package httpclient

import (
	"context"
	"fmt"
	"github.com/fitan/gink/transport/http"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"io"
	"net/url"
	"os"
	"strconv"
	"time"

	//"github.com/fitan/magic/pkg/types"
	"github.com/go-resty/resty/v2"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/registry/cache"
	"go-micro.dev/v4/selector"
)

func BeforeDebug(log *zap.SugaredLogger, bodyLimit int) resty.RequestMiddleware  {
	return func(c *resty.Client, req *resty.Request) error {
		debug, ok := req.Context().Value(http.ContextKeyRequestDebug).(bool)
		if ok && debug {
			c.EnableTrace()

			var body string
			lenValue := req.Header.Get("Content-Length")
			if lenValue == "" {
				body = "header Content-Length unknown"
			}
			lenValueInt, err := strconv.Atoi(lenValue)
			if err != nil {
				body = fmt.Sprintf("header Content-Length strconv err: %s", err.Error())
			}
			if lenValueInt > bodyLimit {
				body = fmt.Sprintf("***** REQUEST TOO LARGE (size - %d) *****", lenValueInt)
			}

			_log := log.With(
				zap.String(
					"traceId", trace.SpanFromContext(req.Context()).SpanContext().TraceID().String(),
				),
			)

			_log.Infow(
				"resty response debug",
				req.Method, req.RawRequest.URL.RequestURI(),
				"HOST", req.RawRequest.URL.Host,
				"BODY", body,
			)
		}
		return nil
	}
}

func AfterDebug(log *zap.SugaredLogger, bodyLimit int64) resty.ResponseMiddleware {
	return func(c *resty.Client, res *resty.Response) error {

		debug, ok := res.Request.Context().Value(http.ContextKeyRequestDebug).(bool)
		if ok && debug {
			var body string
			if res.Size() > bodyLimit {
				body = fmt.Sprintf("***** RESPONSE TOO LARGE (size - %d) *****", res.Size())
			}
			_log := log.With(zap.String("traceId", trace.SpanFromContext(res.Request.Context()).SpanContext().TraceID().String()))

			_log.Infow("resty response debug",
				"STATUS", res.Status(),
				"PROTO", res.RawResponse.Proto,
				"RECEIVED AT", res.ReceivedAt().Format(time.RFC3339Nano),
				"TIME DURATION", res.Time(),
				"BODY", body,
			)

			info := res.Request.TraceInfo()

			_log.Infow("resty trace",
				"DNSLookup", info.DNSLookup,
				"ConnTime", info.ConnTime,
				"TCPConnTime", info.TCPConnTime,
				"TLSHandshake", info.TLSHandshake,
				"ServerTime", info.ServerTime,
				"ResponseTime", info.ResponseTime,
				"TotalTime", info.TotalTime,
				"IsConnReused", info.IsConnReused,
				"IsConnWasIdle", info.IsConnWasIdle,
				"ConnIdleTime", info.ConnIdleTime,
				"RequestAttempt", info.RequestAttempt,
				"RemoteAddr", info.RemoteAddr.String(),
			)
		}

		return nil
	}
}

func ipFactory() sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			return instance, nil
		}, nil, nil
	}
}

type BeforeKitLBConf struct {
	balancerFunc func(s sd.Endpointer) lb.Balancer
}

type BeforeKitLbOption func(conf *BeforeKitLBConf)

func WithBeforeKitLbRandom(seed int64) BeforeKitLbOption {
	return func(conf *BeforeKitLBConf) {
		conf.balancerFunc = func(e sd.Endpointer) lb.Balancer {
			return lb.NewRandom(e,seed)
		}
	}
}

func BeforeKitLb(instance sd.Instancer,options ...BeforeKitLbOption) resty.RequestMiddleware {
	conf := &BeforeKitLBConf{}
	for _, option := range options {
		option(conf)
	}

	e := sd.NewEndpointer(instance, ipFactory(), log.NewLogfmtLogger(os.Stdout))

	if conf.balancerFunc == nil {
		conf.balancerFunc = lb.NewRoundRobin
	}

	balancer := lb.NewRoundRobin(e)

	return func(c *resty.Client, req *resty.Request) error {
		be,err := balancer.Endpoint()
		if err != nil {
			err = errors.Wrap(err, "endpoint error")
			return err
		}
		ip, _ := be(req.Context(), nil)

		urlParse,err := url.Parse(req.URL)
		if err != nil {
			err = errors.Wrap(err, "url parse error")
			return err
		}


		urlParse.Host = ip.(string)

		req.URL = urlParse.String()

		return nil
	}
}


func BeforeMicroSelect(serviceName string, r registry.Registry, options ...selector.SelectOption) resty.RequestMiddleware {
	s := selector.NewSelector(selector.Registry(cache.New(r)), selector.SetStrategy(selector.RoundRobin))
	next, _ := s.Select(serviceName, options...)
	return func(client *resty.Client, request *resty.Request) error {


		node, err := next()
		if err != nil {
			err = errors.Wrap(err, "select node next")
			return err
		}

		reqURL, err := url.Parse(request.URL)
		if err != nil {
			return err
		}

		// If Request.URL is relative path then added c.HostURL into
		// the request URL otherwise Request.URL will be used as-is
		if reqURL.IsAbs() {
			return nil
		}

		request.URL = reqURL.String()
		if len(request.URL) > 0 && request.URL[0] != '/' {
			request.URL = "/" + request.URL
		}


		request.URL = "http://" + node.Address +  request.URL

		return nil
	}
}


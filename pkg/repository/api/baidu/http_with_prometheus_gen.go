package baidu

// Code generated by gowrap. DO NOT EDIT.
// template: ../../../gowrap/templates/prometheus.tmpl
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// BaiduApiWithPrometheus implements BaiduApi interface with all methods wrapped
// with Prometheus metrics
type BaiduApiWithPrometheus struct {
	base         BaiduApi
	instanceName string
}

var baiduapiDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "baiduapi_duration_seconds",
		Help:       "baiduapi runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewBaiduApiWithPrometheus returns an instance of the BaiduApi decorated with prometheus summary metric
func NewBaiduApiWithPrometheus(base BaiduApi, instanceName string) BaiduApiWithPrometheus {
	return BaiduApiWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// GetRoot implements BaiduApi
func (_d BaiduApiWithPrometheus) GetRoot(ctx context.Context) (rp1 *resty.Response, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		baiduapiDurationSummaryVec.WithLabelValues(_d.instanceName, "GetRoot", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetRoot(ctx)
}

// GetRoot1 implements BaiduApi
func (_d BaiduApiWithPrometheus) GetRoot1(ctx context.Context) (rp1 *resty.Response, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		baiduapiDurationSummaryVec.WithLabelValues(_d.instanceName, "GetRoot1", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetRoot1(ctx)
}
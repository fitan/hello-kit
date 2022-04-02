package baidu

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"time"

	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// BaiduServiceWithPrometheus implements BaiduService interface with all methods wrapped
// with Prometheus metrics
type BaiduServiceWithPrometheus struct {
	base         BaiduService
	instanceName string
}

var baiduserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "api_baidu_duration_seconds",
		Help:       "baiduservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewBaiduServiceWithPrometheus returns an instance of the BaiduService decorated with prometheus summary metric
func NewBaiduServiceWithPrometheus(base BaiduService) BaiduServiceWithPrometheus {
	return BaiduServiceWithPrometheus{
		base:         base,
		instanceName: "(down .Interface.Name)",
	}
}

// GetRoot implements BaiduService
func (_d BaiduServiceWithPrometheus) GetRoot(ctx context.Context) (s1 string) {
	_since := time.Now()
	defer func() {
		result := "ok"
		baiduserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "GetRoot", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetRoot(ctx)
}

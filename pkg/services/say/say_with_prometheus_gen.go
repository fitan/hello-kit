package say

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"hello/pkg/services/hello/types"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// SayServiceWithPrometheus implements SayService interface with all methods wrapped
// with Prometheus metrics
type SayServiceWithPrometheus struct {
	base         SayService
	instanceName string
}

var sayserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "sayservice_duration_seconds",
		Help:       "sayservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewSayServiceWithPrometheus returns an instance of the SayService decorated with prometheus summary metric
func NewSayServiceWithPrometheus(base SayService, instanceName string) SayServiceWithPrometheus {
	return SayServiceWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// Say implements SayService
func (_d SayServiceWithPrometheus) Say(ctx context.Context, req types.SayReq) (res types.SayRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		sayserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Say", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Say(ctx, req)
}
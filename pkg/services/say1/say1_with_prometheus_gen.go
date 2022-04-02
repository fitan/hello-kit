package say1

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"hello/pkg/ent"
	"time"

	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Say1ServiceWithPrometheus implements Say1Service interface with all methods wrapped
// with Prometheus metrics
type Say1ServiceWithPrometheus struct {
	base         Say1Service
	instanceName string
}

var say1serviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "services_say1_duration_seconds",
		Help:       "say1service runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewSay1ServiceWithPrometheus returns an instance of the Say1Service decorated with prometheus summary metric
func NewSay1ServiceWithPrometheus(base Say1Service) Say1ServiceWithPrometheus {
	return Say1ServiceWithPrometheus{
		base:         base,
		instanceName: "say1service",
	}
}

// SayPod implements Say1Service
func (_d Say1ServiceWithPrometheus) SayPod(ctx context.Context, req SayPodReq) (pp1 *ent.Pod, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		say1serviceDurationSummaryVec.WithLabelValues(_d.instanceName, "SayPod", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.SayPod(ctx, req)
}

package service

// Code generated by gowrap. DO NOT EDIT.
// template: ../gowrap/templates/prometheus
// gowrap: http://github.com/fitan/gowrap

//go:generate gowrap gen -p hello/pkg/service -i HelloService -t ../gowrap/templates/prometheus -o service_with_prometheus.go -l ""

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// HelloServiceWithPrometheus implements HelloService interface with all methods wrapped
// with Prometheus metrics
type HelloServiceWithPrometheus struct {
	base         HelloService
	instanceName string
}

var helloserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "helloservice_duration_seconds",
		Help:       "helloservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewHelloServiceWithPrometheus returns an instance of the HelloService decorated with prometheus summary metric
func NewHelloServiceWithPrometheus(base HelloService, instanceName string) HelloServiceWithPrometheus {
	return HelloServiceWithPrometheus{
		base:         base,
		instanceName: instanceName,
	}
}

// Foo implements HelloService
func (_d HelloServiceWithPrometheus) Foo(ctx context.Context, s string) (rs string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		helloserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Foo", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Foo(ctx, s)
}

// Say implements HelloService
func (_d HelloServiceWithPrometheus) Say(ctx context.Context, req SayReq) (res SayRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		helloserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Say", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Say(ctx, req)
}

// Say1 implements HelloService
func (_d HelloServiceWithPrometheus) Say1(ctx context.Context, req SayReq) (res SayRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		helloserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Say1", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Say1(ctx, req)
}

// SayHello implements HelloService
func (_d HelloServiceWithPrometheus) SayHello(ctx context.Context, req SayReq) (res SayRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		helloserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "SayHello", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.SayHello(ctx, req)
}

// SayHello1 implements HelloService
func (_d HelloServiceWithPrometheus) SayHello1(ctx context.Context, s1 string, s2 string) (res SayRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		helloserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "SayHello1", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.SayHello1(ctx, s1, s2)
}

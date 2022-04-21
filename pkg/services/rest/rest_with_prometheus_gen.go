package rest

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"time"

	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// RestServiceWithPrometheus implements RestService interface with all methods wrapped
// with Prometheus metrics
type RestServiceWithPrometheus struct {
	base         RestService
	instanceName string
}

var restserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "services_rest_duration_seconds",
		Help:       "restservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewRestServiceWithPrometheus returns an instance of the RestService decorated with prometheus summary metric
func NewRestServiceWithPrometheus(base RestService) RestServiceWithPrometheus {
	return RestServiceWithPrometheus{
		base:         base,
		instanceName: "restservice",
	}
}

// CreateUser implements RestService
func (_d RestServiceWithPrometheus) CreateUser(ctx context.Context, req CreateUserReq) (res CreateUserRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateUser", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateUser(ctx, req)
}

// CreateUsers implements RestService
func (_d RestServiceWithPrometheus) CreateUsers(ctx context.Context, req CreateUsersReq) (res CreateUsersRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateUsers", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateUsers(ctx, req)
}

// DeleteUserById implements RestService
func (_d RestServiceWithPrometheus) DeleteUserById(ctx context.Context, req DeleteUserByIdReq) (res DeleteUserByIdRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteUserById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteUserById(ctx, req)
}

// DeleteUsers implements RestService
func (_d RestServiceWithPrometheus) DeleteUsers(ctx context.Context, req DeleteUsersReq) (res DeleteUsersRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteUsers", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteUsers(ctx, req)
}

// ReadUserById implements RestService
func (_d RestServiceWithPrometheus) ReadUserById(ctx context.Context, req ReadUserByIdReq) (res ReadUserByIdRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ReadUserById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ReadUserById(ctx, req)
}

// ReadUsers implements RestService
func (_d RestServiceWithPrometheus) ReadUsers(ctx context.Context, req ReadUsersReq) (res ReadUsersRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ReadUsers", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ReadUsers(ctx, req)
}

// UpdateUserById implements RestService
func (_d RestServiceWithPrometheus) UpdateUserById(ctx context.Context, req UpdateUserIdReq) (res UpdateUserIdRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateUserById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateUserById(ctx, req)
}

// UpdateUsers implements RestService
func (_d RestServiceWithPrometheus) UpdateUsers(ctx context.Context, req UpdateUsersReq) (res UpdateUsersRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		restserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateUsers", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateUsers(ctx, req)
}
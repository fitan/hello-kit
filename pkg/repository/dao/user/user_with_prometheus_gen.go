package user

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"hello/pkg/ent"
)

// UserServiceWithPrometheus implements UserService interface with all methods wrapped
// with Prometheus metrics
type UserServiceWithPrometheus struct {
	base         UserService
	instanceName string
}

var userserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "dao_user_duration_seconds",
		Help:       "userservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewUserServiceWithPrometheus returns an instance of the UserService decorated with prometheus summary metric
func NewUserServiceWithPrometheus(base UserService) UserServiceWithPrometheus {
	return UserServiceWithPrometheus{
		base:         base,
		instanceName: "userservice",
	}
}

// ByQueries implements UserService
func (_d UserServiceWithPrometheus) ByQueries(ctx context.Context, i interface{}) (res []ent.UserBaseGetRes, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueries", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueries(ctx, i)
}

// Create implements UserService
func (_d UserServiceWithPrometheus) Create(ctx context.Context, v ent.UserBaseCreateReq) (res *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, v)
}

// CreateMany implements UserService
func (_d UserServiceWithPrometheus) CreateMany(ctx context.Context, vs []ent.UserBaseCreateReq) (u1 ent.Users, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateMany(ctx, vs)
}

// CreatePodsSliceByUserId implements UserService
func (_d UserServiceWithPrometheus) CreatePodsSliceByUserId(ctx context.Context, id int, vs []ent.PodBaseCreateReq) (res *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreatePodsSliceByUserId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreatePodsSliceByUserId(ctx, id, vs)
}

// DeleteById implements UserService
func (_d UserServiceWithPrometheus) DeleteById(ctx context.Context, id int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteById(ctx, id)
}

// DeleteMany implements UserService
func (_d UserServiceWithPrometheus) DeleteMany(ctx context.Context, ids []int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteMany(ctx, ids)
}

// GetById implements UserService
func (_d UserServiceWithPrometheus) GetById(ctx context.Context, id int) (res ent.UserBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "GetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetById(ctx, id)
}

// GetPodsSliceByUserId implements UserService
func (_d UserServiceWithPrometheus) GetPodsSliceByUserId(ctx context.Context, id int, i interface{}) (res []ent.PodBaseGetRes, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "GetPodsSliceByUserId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetPodsSliceByUserId(ctx, id, i)
}

// UpdateById implements UserService
func (_d UserServiceWithPrometheus) UpdateById(ctx context.Context, id int, v ent.UserBaseUpdateReq) (up1 *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateById(ctx, id, v)
}

// UpdateMany implements UserService
func (_d UserServiceWithPrometheus) UpdateMany(ctx context.Context, vs []ent.UserBaseUpdateReq) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateMany(ctx, vs)
}

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

// ByQueriesAll implements UserService
func (_d UserServiceWithPrometheus) ByQueriesAll(ctx context.Context, i interface{}) (res []ent.UserBaseGetRes, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueriesAll(ctx, i)
}

// ByQueriesOne implements UserService
func (_d UserServiceWithPrometheus) ByQueriesOne(ctx context.Context, i interface{}) (res ent.UserBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueriesOne", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueriesOne(ctx, i)
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

// RawByQueriesAll implements UserService
func (_d UserServiceWithPrometheus) RawByQueriesAll(ctx context.Context, i interface{}) (res ent.Users, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawByQueriesAll(ctx, i)
}

// RawByQueriesOne implements UserService
func (_d UserServiceWithPrometheus) RawByQueriesOne(ctx context.Context, i interface{}) (res *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawByQueriesOne", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawByQueriesOne(ctx, i)
}

// RawCreate implements UserService
func (_d UserServiceWithPrometheus) RawCreate(ctx context.Context, v *ent.User) (res *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawCreate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawCreate(ctx, v)
}

// RawCreateMany implements UserService
func (_d UserServiceWithPrometheus) RawCreateMany(ctx context.Context, vs ent.Users) (u1 ent.Users, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawCreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawCreateMany(ctx, vs)
}

// RawGetById implements UserService
func (_d UserServiceWithPrometheus) RawGetById(ctx context.Context, id int) (res *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawGetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawGetById(ctx, id)
}

// RawUpdateById implements UserService
func (_d UserServiceWithPrometheus) RawUpdateById(ctx context.Context, id int, v *ent.User) (up1 *ent.User, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawUpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawUpdateById(ctx, id, v)
}

// RawUpdateMany implements UserService
func (_d UserServiceWithPrometheus) RawUpdateMany(ctx context.Context, vs ent.Users) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		userserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawUpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawUpdateMany(ctx, vs)
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

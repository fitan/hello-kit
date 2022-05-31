package service

// Code generated by gowrap. DO NOT EDIT.
// template:
// gowrap: http://github.com/fitan/gowrap

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"context"

	"hello/pkg/ent"
)

// ServiceServiceWithPrometheus implements ServiceService interface with all methods wrapped
// with Prometheus metrics
type ServiceServiceWithPrometheus struct {
	base         ServiceService
	instanceName string
}

var serviceserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "dao_service_duration_seconds",
		Help:       "serviceservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewServiceServiceWithPrometheus returns an instance of the ServiceService decorated with prometheus summary metric
func NewServiceServiceWithPrometheus(base ServiceService) ServiceServiceWithPrometheus {
	return ServiceServiceWithPrometheus{
		base:         base,
		instanceName: "serviceservice",
	}
}

// ByQueriesAll implements ServiceService
func (_d ServiceServiceWithPrometheus) ByQueriesAll(ctx context.Context, i interface{}) (res []ent.ServiceBaseGetRes, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueriesAll(ctx, i)
}

// ByQueriesOne implements ServiceService
func (_d ServiceServiceWithPrometheus) ByQueriesOne(ctx context.Context, i interface{}) (res ent.ServiceBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ByQueriesOne", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ByQueriesOne(ctx, i)
}

// Create implements ServiceService
func (_d ServiceServiceWithPrometheus) Create(ctx context.Context, v ent.ServiceBaseCreateReq) (res *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "Create", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.Create(ctx, v)
}

// CreateMany implements ServiceService
func (_d ServiceServiceWithPrometheus) CreateMany(ctx context.Context, vs []ent.ServiceBaseCreateReq) (s1 ent.Services, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateMany(ctx, vs)
}

// CreateProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) CreateProjectByServiceId(ctx context.Context, id int, v ent.ProjectBaseCreateReq) (res *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "CreateProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.CreateProjectByServiceId(ctx, id, v)
}

// DeleteById implements ServiceService
func (_d ServiceServiceWithPrometheus) DeleteById(ctx context.Context, id int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteById(ctx, id)
}

// DeleteMany implements ServiceService
func (_d ServiceServiceWithPrometheus) DeleteMany(ctx context.Context, ids []int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "DeleteMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.DeleteMany(ctx, ids)
}

// GetById implements ServiceService
func (_d ServiceServiceWithPrometheus) GetById(ctx context.Context, id int) (res ent.ServiceBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "GetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetById(ctx, id)
}

// GetProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) GetProjectByServiceId(ctx context.Context, id int) (res ent.ProjectBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "GetProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.GetProjectByServiceId(ctx, id)
}

// RawAddBindProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) RawAddBindProjectByServiceId(ctx context.Context, id int, addId int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawAddBindProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawAddBindProjectByServiceId(ctx, id, addId)
}

// RawByQueriesAll implements ServiceService
func (_d ServiceServiceWithPrometheus) RawByQueriesAll(ctx context.Context, i interface{}) (res ent.Services, count int, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawByQueriesAll(ctx, i)
}

// RawByQueriesOne implements ServiceService
func (_d ServiceServiceWithPrometheus) RawByQueriesOne(ctx context.Context, i interface{}) (res *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawByQueriesOne", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawByQueriesOne(ctx, i)
}

// RawCreate implements ServiceService
func (_d ServiceServiceWithPrometheus) RawCreate(ctx context.Context, v *ent.Service) (res *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawCreate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawCreate(ctx, v)
}

// RawCreateMany implements ServiceService
func (_d ServiceServiceWithPrometheus) RawCreateMany(ctx context.Context, vs ent.Services) (s1 ent.Services, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawCreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawCreateMany(ctx, vs)
}

// RawCreateProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) RawCreateProjectByServiceId(ctx context.Context, id int, v *ent.Project) (res *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawCreateProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawCreateProjectByServiceId(ctx, id, v)
}

// RawDeleteProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) RawDeleteProjectByServiceId(ctx context.Context, id int, deleteId int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawDeleteProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawDeleteProjectByServiceId(ctx, id, deleteId)
}

// RawGetById implements ServiceService
func (_d ServiceServiceWithPrometheus) RawGetById(ctx context.Context, id int) (res *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawGetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawGetById(ctx, id)
}

// RawGetProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) RawGetProjectByServiceId(ctx context.Context, id int) (res *ent.Project, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawGetProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawGetProjectByServiceId(ctx, id)
}

// RawRemoveBindProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) RawRemoveBindProjectByServiceId(ctx context.Context, id int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawRemoveBindProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawRemoveBindProjectByServiceId(ctx, id)
}

// RawUpdateBindProjectByServiceId implements ServiceService
func (_d ServiceServiceWithPrometheus) RawUpdateBindProjectByServiceId(ctx context.Context, id int, updateId int) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawUpdateBindProjectByServiceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawUpdateBindProjectByServiceId(ctx, id, updateId)
}

// RawUpdateById implements ServiceService
func (_d ServiceServiceWithPrometheus) RawUpdateById(ctx context.Context, id int, v *ent.Service) (sp1 *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawUpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawUpdateById(ctx, id, v)
}

// RawUpdateMany implements ServiceService
func (_d ServiceServiceWithPrometheus) RawUpdateMany(ctx context.Context, vs ent.Services) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "RawUpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.RawUpdateMany(ctx, vs)
}

// UpdateById implements ServiceService
func (_d ServiceServiceWithPrometheus) UpdateById(ctx context.Context, id int, v ent.ServiceBaseUpdateReq) (sp1 *ent.Service, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateById(ctx, id, v)
}

// UpdateMany implements ServiceService
func (_d ServiceServiceWithPrometheus) UpdateMany(ctx context.Context, vs []ent.ServiceBaseUpdateReq) (err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		serviceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "UpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.UpdateMany(ctx, vs)
}
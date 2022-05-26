package project

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

// ProjectServiceWithPrometheus implements ProjectService interface with all methods wrapped
// with Prometheus metrics
type ProjectServiceWithPrometheus struct {
	base         ProjectService
	instanceName string
}

var projectserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "services_project_duration_seconds",
		Help:       "projectservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewProjectServiceWithPrometheus returns an instance of the ProjectService decorated with prometheus summary metric
func NewProjectServiceWithPrometheus(base ProjectService) ProjectServiceWithPrometheus {
	return ProjectServiceWithPrometheus{
		base:         base,
		instanceName: "projectservice",
	}
}

// ProjectRestByQueriesAll implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestByQueriesAll(ctx context.Context, req ent.ProjectRestByQueriesAllReq) (res ent.ProjectRestByQueriesAllRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestByQueriesAll(ctx, req)
}

// ProjectRestCreate implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestCreate(ctx context.Context, req ent.ProjectRestCreateReq) (res *ent.Project, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestCreate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestCreate(ctx, req)
}

// ProjectRestCreateMany implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestCreateMany(ctx context.Context, req ent.ProjectRestCreateManyReq) (res ent.Projects, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestCreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestCreateMany(ctx, req)
}

// ProjectRestCreateServicesByProjectId implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestCreateServicesByProjectId(ctx context.Context, req ent.ProjectRestCreateServicesByProjectIdReq) (res *ent.Project, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestCreateServicesByProjectId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestCreateServicesByProjectId(ctx, req)
}

// ProjectRestDeleteById implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestDeleteById(ctx context.Context, req ent.ProjectRestDeleteByIdReq) (success bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestDeleteById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestDeleteById(ctx, req)
}

// ProjectRestDeleteMany implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestDeleteMany(ctx context.Context, req ent.ProjectRestDeleteManyReq) (success bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestDeleteMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestDeleteMany(ctx, req)
}

// ProjectRestGetById implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestGetById(ctx context.Context, req ent.ProjectRestGetByIdReq) (res ent.ProjectBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestGetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestGetById(ctx, req)
}

// ProjectRestGetServicesByProjectId implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestGetServicesByProjectId(ctx context.Context, req ent.ProjectRestGetServicesByProjectIdReq) (res ent.ProjectRestGetServicesByProjectIdRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestGetServicesByProjectId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestGetServicesByProjectId(ctx, req)
}

// ProjectRestUpdateById implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestUpdateById(ctx context.Context, req ent.ProjectRestUpdateByIdReq) (res *ent.Project, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestUpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestUpdateById(ctx, req)
}

// ProjectRestUpdateMany implements ProjectService
func (_d ProjectServiceWithPrometheus) ProjectRestUpdateMany(ctx context.Context, req ent.ProjectRestUpdateManyReq) (success bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		projectserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ProjectRestUpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ProjectRestUpdateMany(ctx, req)
}

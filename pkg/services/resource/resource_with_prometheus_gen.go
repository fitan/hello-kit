package resource

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

// ResourceServiceWithPrometheus implements ResourceService interface with all methods wrapped
// with Prometheus metrics
type ResourceServiceWithPrometheus struct {
	base         ResourceService
	instanceName string
}

var resourceserviceDurationSummaryVec = promauto.NewSummaryVec(
	prometheus.SummaryOpts{
		Name:       "services_resource_duration_seconds",
		Help:       "resourceservice runtime duration and result",
		MaxAge:     time.Minute,
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"instance_name", "method", "result"})

// NewResourceServiceWithPrometheus returns an instance of the ResourceService decorated with prometheus summary metric
func NewResourceServiceWithPrometheus(base ResourceService) ResourceServiceWithPrometheus {
	return ResourceServiceWithPrometheus{
		base:         base,
		instanceName: "resourceservice",
	}
}

// ResourceRestAddBindNextByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestAddBindNextByResourceId(ctx context.Context, req ent.ResourceRestAddBindNextByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestAddBindNextByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestAddBindNextByResourceId(ctx, req)
}

// ResourceRestAddBindPreByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestAddBindPreByResourceId(ctx context.Context, req ent.ResourceRestAddBindPreByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestAddBindPreByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestAddBindPreByResourceId(ctx, req)
}

// ResourceRestByQueriesAll implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestByQueriesAll(ctx context.Context, req ent.ResourceRestByQueriesAllReq) (res ent.ResourceRestByQueriesAllRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestByQueriesAll", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestByQueriesAll(ctx, req)
}

// ResourceRestCreate implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestCreate(ctx context.Context, req ent.ResourceRestCreateReq) (res *ent.Resource, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestCreate", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestCreate(ctx, req)
}

// ResourceRestCreateMany implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestCreateMany(ctx context.Context, req ent.ResourceRestCreateManyReq) (res ent.Resources, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestCreateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestCreateMany(ctx, req)
}

// ResourceRestCreatePreByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestCreatePreByResourceId(ctx context.Context, req ent.ResourceRestCreatePreByResourceIdReq) (res *ent.Resource, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestCreatePreByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestCreatePreByResourceId(ctx, req)
}

// ResourceRestCreateResourcesByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestCreateResourcesByResourceId(ctx context.Context, req ent.ResourceRestCreateResourcesByResourceIdReq) (res *ent.Resource, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestCreateResourcesByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestCreateResourcesByResourceId(ctx, req)
}

// ResourceRestDeleteById implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestDeleteById(ctx context.Context, req ent.ResourceRestDeleteByIdReq) (success bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestDeleteById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestDeleteById(ctx, req)
}

// ResourceRestDeleteMany implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestDeleteMany(ctx context.Context, req ent.ResourceRestDeleteManyReq) (success bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestDeleteMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestDeleteMany(ctx, req)
}

// ResourceRestDeleteNextByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestDeleteNextByResourceId(ctx context.Context, req ent.ResourceRestDeleteNextByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestDeleteNextByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestDeleteNextByResourceId(ctx, req)
}

// ResourceRestDeletePreByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestDeletePreByResourceId(ctx context.Context, req ent.ResourceRestDeletePreByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestDeletePreByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestDeletePreByResourceId(ctx, req)
}

// ResourceRestGetById implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestGetById(ctx context.Context, req ent.ResourceRestGetByIdReq) (res ent.ResourceBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestGetById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestGetById(ctx, req)
}

// ResourceRestGetNextByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestGetNextByResourceId(ctx context.Context, req ent.ResourceRestGetNextByResourceIdReq) (res ent.ResourceRestGetNextByResourceIdRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestGetNextByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestGetNextByResourceId(ctx, req)
}

// ResourceRestGetPreByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestGetPreByResourceId(ctx context.Context, req ent.ResourceRestGetPreByResourceIdReq) (res ent.ResourceBaseGetRes, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestGetPreByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestGetPreByResourceId(ctx, req)
}

// ResourceRestRemoveBindNextByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestRemoveBindNextByResourceId(ctx context.Context, req ent.ResourceRestRemoveBindNextByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestRemoveBindNextByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestRemoveBindNextByResourceId(ctx, req)
}

// ResourceRestRemoveBindPreByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestRemoveBindPreByResourceId(ctx context.Context, req ent.ResourceRestRemoveBindPreByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestRemoveBindPreByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestRemoveBindPreByResourceId(ctx, req)
}

// ResourceRestUpdateBindNextByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestUpdateBindNextByResourceId(ctx context.Context, req ent.ResourceRestUpdateBindNextByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestUpdateBindNextByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestUpdateBindNextByResourceId(ctx, req)
}

// ResourceRestUpdateBindPreByResourceId implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestUpdateBindPreByResourceId(ctx context.Context, req ent.ResourceRestUpdateBindPreByResourceIdReq) (res string, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestUpdateBindPreByResourceId", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestUpdateBindPreByResourceId(ctx, req)
}

// ResourceRestUpdateById implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestUpdateById(ctx context.Context, req ent.ResourceRestUpdateByIdReq) (res *ent.Resource, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestUpdateById", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestUpdateById(ctx, req)
}

// ResourceRestUpdateMany implements ResourceService
func (_d ResourceServiceWithPrometheus) ResourceRestUpdateMany(ctx context.Context, req ent.ResourceRestUpdateManyReq) (success bool, err error) {
	_since := time.Now()
	defer func() {
		result := "ok"
		if err != nil {
			result = "error"
		}

		resourceserviceDurationSummaryVec.WithLabelValues(_d.instanceName, "ResourceRestUpdateMany", result).Observe(time.Since(_since).Seconds())
	}()
	return _d.base.ResourceRestUpdateMany(ctx, req)
}
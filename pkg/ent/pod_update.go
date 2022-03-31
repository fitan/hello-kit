// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hello/pkg/ent/pod"
	"hello/pkg/ent/predicate"
	"hello/pkg/ent/spiderdevtblservicetree"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PodUpdate is the builder for updating Pod entities.
type PodUpdate struct {
	config
	hooks    []Hook
	mutation *PodMutation
}

// Where appends a list predicates to the PodUpdate builder.
func (pu *PodUpdate) Where(ps ...predicate.Pod) *PodUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetClusterName sets the "cluster_name" field.
func (pu *PodUpdate) SetClusterName(s string) *PodUpdate {
	pu.mutation.SetClusterName(s)
	return pu
}

// SetNillableClusterName sets the "cluster_name" field if the given value is not nil.
func (pu *PodUpdate) SetNillableClusterName(s *string) *PodUpdate {
	if s != nil {
		pu.SetClusterName(*s)
	}
	return pu
}

// ClearClusterName clears the value of the "cluster_name" field.
func (pu *PodUpdate) ClearClusterName() *PodUpdate {
	pu.mutation.ClearClusterName()
	return pu
}

// SetNamespace sets the "namespace" field.
func (pu *PodUpdate) SetNamespace(s string) *PodUpdate {
	pu.mutation.SetNamespace(s)
	return pu
}

// SetNillableNamespace sets the "namespace" field if the given value is not nil.
func (pu *PodUpdate) SetNillableNamespace(s *string) *PodUpdate {
	if s != nil {
		pu.SetNamespace(*s)
	}
	return pu
}

// ClearNamespace clears the value of the "namespace" field.
func (pu *PodUpdate) ClearNamespace() *PodUpdate {
	pu.mutation.ClearNamespace()
	return pu
}

// SetServiceName sets the "service_name" field.
func (pu *PodUpdate) SetServiceName(s string) *PodUpdate {
	pu.mutation.SetServiceName(s)
	return pu
}

// SetNillableServiceName sets the "service_name" field if the given value is not nil.
func (pu *PodUpdate) SetNillableServiceName(s *string) *PodUpdate {
	if s != nil {
		pu.SetServiceName(*s)
	}
	return pu
}

// ClearServiceName clears the value of the "service_name" field.
func (pu *PodUpdate) ClearServiceName() *PodUpdate {
	pu.mutation.ClearServiceName()
	return pu
}

// SetPodName sets the "pod_name" field.
func (pu *PodUpdate) SetPodName(s string) *PodUpdate {
	pu.mutation.SetPodName(s)
	return pu
}

// SetNillablePodName sets the "pod_name" field if the given value is not nil.
func (pu *PodUpdate) SetNillablePodName(s *string) *PodUpdate {
	if s != nil {
		pu.SetPodName(*s)
	}
	return pu
}

// ClearPodName clears the value of the "pod_name" field.
func (pu *PodUpdate) ClearPodName() *PodUpdate {
	pu.mutation.ClearPodName()
	return pu
}

// SetResourceVersion sets the "resource_version" field.
func (pu *PodUpdate) SetResourceVersion(s string) *PodUpdate {
	pu.mutation.SetResourceVersion(s)
	return pu
}

// SetNillableResourceVersion sets the "resource_version" field if the given value is not nil.
func (pu *PodUpdate) SetNillableResourceVersion(s *string) *PodUpdate {
	if s != nil {
		pu.SetResourceVersion(*s)
	}
	return pu
}

// ClearResourceVersion clears the value of the "resource_version" field.
func (pu *PodUpdate) ClearResourceVersion() *PodUpdate {
	pu.mutation.ClearResourceVersion()
	return pu
}

// SetPodIP sets the "pod_ip" field.
func (pu *PodUpdate) SetPodIP(s string) *PodUpdate {
	pu.mutation.SetPodIP(s)
	return pu
}

// SetNillablePodIP sets the "pod_ip" field if the given value is not nil.
func (pu *PodUpdate) SetNillablePodIP(s *string) *PodUpdate {
	if s != nil {
		pu.SetPodIP(*s)
	}
	return pu
}

// ClearPodIP clears the value of the "pod_ip" field.
func (pu *PodUpdate) ClearPodIP() *PodUpdate {
	pu.mutation.ClearPodIP()
	return pu
}

// SetHostIP sets the "host_ip" field.
func (pu *PodUpdate) SetHostIP(s string) *PodUpdate {
	pu.mutation.SetHostIP(s)
	return pu
}

// SetNillableHostIP sets the "host_ip" field if the given value is not nil.
func (pu *PodUpdate) SetNillableHostIP(s *string) *PodUpdate {
	if s != nil {
		pu.SetHostIP(*s)
	}
	return pu
}

// ClearHostIP clears the value of the "host_ip" field.
func (pu *PodUpdate) ClearHostIP() *PodUpdate {
	pu.mutation.ClearHostIP()
	return pu
}

// SetStartTime sets the "start_time" field.
func (pu *PodUpdate) SetStartTime(t time.Time) *PodUpdate {
	pu.mutation.SetStartTime(t)
	return pu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pu *PodUpdate) SetNillableStartTime(t *time.Time) *PodUpdate {
	if t != nil {
		pu.SetStartTime(*t)
	}
	return pu
}

// ClearStartTime clears the value of the "start_time" field.
func (pu *PodUpdate) ClearStartTime() *PodUpdate {
	pu.mutation.ClearStartTime()
	return pu
}

// SetPhase sets the "phase" field.
func (pu *PodUpdate) SetPhase(s string) *PodUpdate {
	pu.mutation.SetPhase(s)
	return pu
}

// SetNillablePhase sets the "phase" field if the given value is not nil.
func (pu *PodUpdate) SetNillablePhase(s *string) *PodUpdate {
	if s != nil {
		pu.SetPhase(*s)
	}
	return pu
}

// ClearPhase clears the value of the "phase" field.
func (pu *PodUpdate) ClearPhase() *PodUpdate {
	pu.mutation.ClearPhase()
	return pu
}

// SetReason sets the "reason" field.
func (pu *PodUpdate) SetReason(s string) *PodUpdate {
	pu.mutation.SetReason(s)
	return pu
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (pu *PodUpdate) SetNillableReason(s *string) *PodUpdate {
	if s != nil {
		pu.SetReason(*s)
	}
	return pu
}

// ClearReason clears the value of the "reason" field.
func (pu *PodUpdate) ClearReason() *PodUpdate {
	pu.mutation.ClearReason()
	return pu
}

// SetMessage sets the "message" field.
func (pu *PodUpdate) SetMessage(s string) *PodUpdate {
	pu.mutation.SetMessage(s)
	return pu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (pu *PodUpdate) SetNillableMessage(s *string) *PodUpdate {
	if s != nil {
		pu.SetMessage(*s)
	}
	return pu
}

// ClearMessage clears the value of the "message" field.
func (pu *PodUpdate) ClearMessage() *PodUpdate {
	pu.mutation.ClearMessage()
	return pu
}

// SetDetail sets the "detail" field.
func (pu *PodUpdate) SetDetail(s string) *PodUpdate {
	pu.mutation.SetDetail(s)
	return pu
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (pu *PodUpdate) SetNillableDetail(s *string) *PodUpdate {
	if s != nil {
		pu.SetDetail(*s)
	}
	return pu
}

// ClearDetail clears the value of the "detail" field.
func (pu *PodUpdate) ClearDetail() *PodUpdate {
	pu.mutation.ClearDetail()
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PodUpdate) SetCreatedAt(t time.Time) *PodUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PodUpdate) SetNillableCreatedAt(t *time.Time) *PodUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (pu *PodUpdate) ClearCreatedAt() *PodUpdate {
	pu.mutation.ClearCreatedAt()
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PodUpdate) SetUpdatedAt(t time.Time) *PodUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pu *PodUpdate) SetNillableUpdatedAt(t *time.Time) *PodUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PodUpdate) ClearUpdatedAt() *PodUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetServicetreeID sets the "servicetree" edge to the SpiderDevTblServicetree entity by ID.
func (pu *PodUpdate) SetServicetreeID(id int32) *PodUpdate {
	pu.mutation.SetServicetreeID(id)
	return pu
}

// SetNillableServicetreeID sets the "servicetree" edge to the SpiderDevTblServicetree entity by ID if the given value is not nil.
func (pu *PodUpdate) SetNillableServicetreeID(id *int32) *PodUpdate {
	if id != nil {
		pu = pu.SetServicetreeID(*id)
	}
	return pu
}

// SetServicetree sets the "servicetree" edge to the SpiderDevTblServicetree entity.
func (pu *PodUpdate) SetServicetree(s *SpiderDevTblServicetree) *PodUpdate {
	return pu.SetServicetreeID(s.ID)
}

// Mutation returns the PodMutation object of the builder.
func (pu *PodUpdate) Mutation() *PodMutation {
	return pu.mutation
}

// ClearServicetree clears the "servicetree" edge to the SpiderDevTblServicetree entity.
func (pu *PodUpdate) ClearServicetree() *PodUpdate {
	pu.mutation.ClearServicetree()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PodUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PodUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PodUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pod.Table,
			Columns: pod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: pod.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ClusterName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldClusterName,
		})
	}
	if pu.mutation.ClusterNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldClusterName,
		})
	}
	if value, ok := pu.mutation.Namespace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldNamespace,
		})
	}
	if pu.mutation.NamespaceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldNamespace,
		})
	}
	if value, ok := pu.mutation.ServiceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldServiceName,
		})
	}
	if pu.mutation.ServiceNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldServiceName,
		})
	}
	if value, ok := pu.mutation.PodName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPodName,
		})
	}
	if pu.mutation.PodNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldPodName,
		})
	}
	if value, ok := pu.mutation.ResourceVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldResourceVersion,
		})
	}
	if pu.mutation.ResourceVersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldResourceVersion,
		})
	}
	if value, ok := pu.mutation.PodIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPodIP,
		})
	}
	if pu.mutation.PodIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldPodIP,
		})
	}
	if value, ok := pu.mutation.HostIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldHostIP,
		})
	}
	if pu.mutation.HostIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldHostIP,
		})
	}
	if value, ok := pu.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldStartTime,
		})
	}
	if pu.mutation.StartTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pod.FieldStartTime,
		})
	}
	if value, ok := pu.mutation.Phase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPhase,
		})
	}
	if pu.mutation.PhaseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldPhase,
		})
	}
	if value, ok := pu.mutation.Reason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldReason,
		})
	}
	if pu.mutation.ReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldReason,
		})
	}
	if value, ok := pu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldMessage,
		})
	}
	if pu.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldMessage,
		})
	}
	if value, ok := pu.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldDetail,
		})
	}
	if pu.mutation.DetailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldDetail,
		})
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldCreatedAt,
		})
	}
	if pu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pod.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldUpdatedAt,
		})
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pod.FieldUpdatedAt,
		})
	}
	if pu.mutation.ServicetreeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pod.ServicetreeTable,
			Columns: []string{pod.ServicetreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: spiderdevtblservicetree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ServicetreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pod.ServicetreeTable,
			Columns: []string{pod.ServicetreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: spiderdevtblservicetree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PodUpdateOne is the builder for updating a single Pod entity.
type PodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PodMutation
}

// SetClusterName sets the "cluster_name" field.
func (puo *PodUpdateOne) SetClusterName(s string) *PodUpdateOne {
	puo.mutation.SetClusterName(s)
	return puo
}

// SetNillableClusterName sets the "cluster_name" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableClusterName(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetClusterName(*s)
	}
	return puo
}

// ClearClusterName clears the value of the "cluster_name" field.
func (puo *PodUpdateOne) ClearClusterName() *PodUpdateOne {
	puo.mutation.ClearClusterName()
	return puo
}

// SetNamespace sets the "namespace" field.
func (puo *PodUpdateOne) SetNamespace(s string) *PodUpdateOne {
	puo.mutation.SetNamespace(s)
	return puo
}

// SetNillableNamespace sets the "namespace" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableNamespace(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetNamespace(*s)
	}
	return puo
}

// ClearNamespace clears the value of the "namespace" field.
func (puo *PodUpdateOne) ClearNamespace() *PodUpdateOne {
	puo.mutation.ClearNamespace()
	return puo
}

// SetServiceName sets the "service_name" field.
func (puo *PodUpdateOne) SetServiceName(s string) *PodUpdateOne {
	puo.mutation.SetServiceName(s)
	return puo
}

// SetNillableServiceName sets the "service_name" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableServiceName(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetServiceName(*s)
	}
	return puo
}

// ClearServiceName clears the value of the "service_name" field.
func (puo *PodUpdateOne) ClearServiceName() *PodUpdateOne {
	puo.mutation.ClearServiceName()
	return puo
}

// SetPodName sets the "pod_name" field.
func (puo *PodUpdateOne) SetPodName(s string) *PodUpdateOne {
	puo.mutation.SetPodName(s)
	return puo
}

// SetNillablePodName sets the "pod_name" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillablePodName(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetPodName(*s)
	}
	return puo
}

// ClearPodName clears the value of the "pod_name" field.
func (puo *PodUpdateOne) ClearPodName() *PodUpdateOne {
	puo.mutation.ClearPodName()
	return puo
}

// SetResourceVersion sets the "resource_version" field.
func (puo *PodUpdateOne) SetResourceVersion(s string) *PodUpdateOne {
	puo.mutation.SetResourceVersion(s)
	return puo
}

// SetNillableResourceVersion sets the "resource_version" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableResourceVersion(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetResourceVersion(*s)
	}
	return puo
}

// ClearResourceVersion clears the value of the "resource_version" field.
func (puo *PodUpdateOne) ClearResourceVersion() *PodUpdateOne {
	puo.mutation.ClearResourceVersion()
	return puo
}

// SetPodIP sets the "pod_ip" field.
func (puo *PodUpdateOne) SetPodIP(s string) *PodUpdateOne {
	puo.mutation.SetPodIP(s)
	return puo
}

// SetNillablePodIP sets the "pod_ip" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillablePodIP(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetPodIP(*s)
	}
	return puo
}

// ClearPodIP clears the value of the "pod_ip" field.
func (puo *PodUpdateOne) ClearPodIP() *PodUpdateOne {
	puo.mutation.ClearPodIP()
	return puo
}

// SetHostIP sets the "host_ip" field.
func (puo *PodUpdateOne) SetHostIP(s string) *PodUpdateOne {
	puo.mutation.SetHostIP(s)
	return puo
}

// SetNillableHostIP sets the "host_ip" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableHostIP(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetHostIP(*s)
	}
	return puo
}

// ClearHostIP clears the value of the "host_ip" field.
func (puo *PodUpdateOne) ClearHostIP() *PodUpdateOne {
	puo.mutation.ClearHostIP()
	return puo
}

// SetStartTime sets the "start_time" field.
func (puo *PodUpdateOne) SetStartTime(t time.Time) *PodUpdateOne {
	puo.mutation.SetStartTime(t)
	return puo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableStartTime(t *time.Time) *PodUpdateOne {
	if t != nil {
		puo.SetStartTime(*t)
	}
	return puo
}

// ClearStartTime clears the value of the "start_time" field.
func (puo *PodUpdateOne) ClearStartTime() *PodUpdateOne {
	puo.mutation.ClearStartTime()
	return puo
}

// SetPhase sets the "phase" field.
func (puo *PodUpdateOne) SetPhase(s string) *PodUpdateOne {
	puo.mutation.SetPhase(s)
	return puo
}

// SetNillablePhase sets the "phase" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillablePhase(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetPhase(*s)
	}
	return puo
}

// ClearPhase clears the value of the "phase" field.
func (puo *PodUpdateOne) ClearPhase() *PodUpdateOne {
	puo.mutation.ClearPhase()
	return puo
}

// SetReason sets the "reason" field.
func (puo *PodUpdateOne) SetReason(s string) *PodUpdateOne {
	puo.mutation.SetReason(s)
	return puo
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableReason(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetReason(*s)
	}
	return puo
}

// ClearReason clears the value of the "reason" field.
func (puo *PodUpdateOne) ClearReason() *PodUpdateOne {
	puo.mutation.ClearReason()
	return puo
}

// SetMessage sets the "message" field.
func (puo *PodUpdateOne) SetMessage(s string) *PodUpdateOne {
	puo.mutation.SetMessage(s)
	return puo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableMessage(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetMessage(*s)
	}
	return puo
}

// ClearMessage clears the value of the "message" field.
func (puo *PodUpdateOne) ClearMessage() *PodUpdateOne {
	puo.mutation.ClearMessage()
	return puo
}

// SetDetail sets the "detail" field.
func (puo *PodUpdateOne) SetDetail(s string) *PodUpdateOne {
	puo.mutation.SetDetail(s)
	return puo
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableDetail(s *string) *PodUpdateOne {
	if s != nil {
		puo.SetDetail(*s)
	}
	return puo
}

// ClearDetail clears the value of the "detail" field.
func (puo *PodUpdateOne) ClearDetail() *PodUpdateOne {
	puo.mutation.ClearDetail()
	return puo
}

// SetCreatedAt sets the "created_at" field.
func (puo *PodUpdateOne) SetCreatedAt(t time.Time) *PodUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableCreatedAt(t *time.Time) *PodUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (puo *PodUpdateOne) ClearCreatedAt() *PodUpdateOne {
	puo.mutation.ClearCreatedAt()
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PodUpdateOne) SetUpdatedAt(t time.Time) *PodUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (puo *PodUpdateOne) SetNillableUpdatedAt(t *time.Time) *PodUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PodUpdateOne) ClearUpdatedAt() *PodUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetServicetreeID sets the "servicetree" edge to the SpiderDevTblServicetree entity by ID.
func (puo *PodUpdateOne) SetServicetreeID(id int32) *PodUpdateOne {
	puo.mutation.SetServicetreeID(id)
	return puo
}

// SetNillableServicetreeID sets the "servicetree" edge to the SpiderDevTblServicetree entity by ID if the given value is not nil.
func (puo *PodUpdateOne) SetNillableServicetreeID(id *int32) *PodUpdateOne {
	if id != nil {
		puo = puo.SetServicetreeID(*id)
	}
	return puo
}

// SetServicetree sets the "servicetree" edge to the SpiderDevTblServicetree entity.
func (puo *PodUpdateOne) SetServicetree(s *SpiderDevTblServicetree) *PodUpdateOne {
	return puo.SetServicetreeID(s.ID)
}

// Mutation returns the PodMutation object of the builder.
func (puo *PodUpdateOne) Mutation() *PodMutation {
	return puo.mutation
}

// ClearServicetree clears the "servicetree" edge to the SpiderDevTblServicetree entity.
func (puo *PodUpdateOne) ClearServicetree() *PodUpdateOne {
	puo.mutation.ClearServicetree()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PodUpdateOne) Select(field string, fields ...string) *PodUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pod entity.
func (puo *PodUpdateOne) Save(ctx context.Context) (*Pod, error) {
	var (
		err  error
		node *Pod
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PodUpdateOne) SaveX(ctx context.Context) *Pod {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PodUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PodUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PodUpdateOne) sqlSave(ctx context.Context) (_node *Pod, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pod.Table,
			Columns: pod.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: pod.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pod.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pod.FieldID)
		for _, f := range fields {
			if !pod.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pod.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.ClusterName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldClusterName,
		})
	}
	if puo.mutation.ClusterNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldClusterName,
		})
	}
	if value, ok := puo.mutation.Namespace(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldNamespace,
		})
	}
	if puo.mutation.NamespaceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldNamespace,
		})
	}
	if value, ok := puo.mutation.ServiceName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldServiceName,
		})
	}
	if puo.mutation.ServiceNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldServiceName,
		})
	}
	if value, ok := puo.mutation.PodName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPodName,
		})
	}
	if puo.mutation.PodNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldPodName,
		})
	}
	if value, ok := puo.mutation.ResourceVersion(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldResourceVersion,
		})
	}
	if puo.mutation.ResourceVersionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldResourceVersion,
		})
	}
	if value, ok := puo.mutation.PodIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPodIP,
		})
	}
	if puo.mutation.PodIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldPodIP,
		})
	}
	if value, ok := puo.mutation.HostIP(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldHostIP,
		})
	}
	if puo.mutation.HostIPCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldHostIP,
		})
	}
	if value, ok := puo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldStartTime,
		})
	}
	if puo.mutation.StartTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pod.FieldStartTime,
		})
	}
	if value, ok := puo.mutation.Phase(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPhase,
		})
	}
	if puo.mutation.PhaseCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldPhase,
		})
	}
	if value, ok := puo.mutation.Reason(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldReason,
		})
	}
	if puo.mutation.ReasonCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldReason,
		})
	}
	if value, ok := puo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldMessage,
		})
	}
	if puo.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldMessage,
		})
	}
	if value, ok := puo.mutation.Detail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldDetail,
		})
	}
	if puo.mutation.DetailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pod.FieldDetail,
		})
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldCreatedAt,
		})
	}
	if puo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pod.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldUpdatedAt,
		})
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: pod.FieldUpdatedAt,
		})
	}
	if puo.mutation.ServicetreeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pod.ServicetreeTable,
			Columns: []string{pod.ServicetreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: spiderdevtblservicetree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ServicetreeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   pod.ServicetreeTable,
			Columns: []string{pod.ServicetreeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: spiderdevtblservicetree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pod{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pod.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
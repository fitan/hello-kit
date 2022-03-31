// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hello/pkg/ent/pod"
	"hello/pkg/ent/spiderdevtblservicetree"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PodCreate is the builder for creating a Pod entity.
type PodCreate struct {
	config
	mutation *PodMutation
	hooks    []Hook
}

// SetClusterName sets the "cluster_name" field.
func (pc *PodCreate) SetClusterName(s string) *PodCreate {
	pc.mutation.SetClusterName(s)
	return pc
}

// SetNillableClusterName sets the "cluster_name" field if the given value is not nil.
func (pc *PodCreate) SetNillableClusterName(s *string) *PodCreate {
	if s != nil {
		pc.SetClusterName(*s)
	}
	return pc
}

// SetNamespace sets the "namespace" field.
func (pc *PodCreate) SetNamespace(s string) *PodCreate {
	pc.mutation.SetNamespace(s)
	return pc
}

// SetNillableNamespace sets the "namespace" field if the given value is not nil.
func (pc *PodCreate) SetNillableNamespace(s *string) *PodCreate {
	if s != nil {
		pc.SetNamespace(*s)
	}
	return pc
}

// SetServiceName sets the "service_name" field.
func (pc *PodCreate) SetServiceName(s string) *PodCreate {
	pc.mutation.SetServiceName(s)
	return pc
}

// SetNillableServiceName sets the "service_name" field if the given value is not nil.
func (pc *PodCreate) SetNillableServiceName(s *string) *PodCreate {
	if s != nil {
		pc.SetServiceName(*s)
	}
	return pc
}

// SetPodName sets the "pod_name" field.
func (pc *PodCreate) SetPodName(s string) *PodCreate {
	pc.mutation.SetPodName(s)
	return pc
}

// SetNillablePodName sets the "pod_name" field if the given value is not nil.
func (pc *PodCreate) SetNillablePodName(s *string) *PodCreate {
	if s != nil {
		pc.SetPodName(*s)
	}
	return pc
}

// SetResourceVersion sets the "resource_version" field.
func (pc *PodCreate) SetResourceVersion(s string) *PodCreate {
	pc.mutation.SetResourceVersion(s)
	return pc
}

// SetNillableResourceVersion sets the "resource_version" field if the given value is not nil.
func (pc *PodCreate) SetNillableResourceVersion(s *string) *PodCreate {
	if s != nil {
		pc.SetResourceVersion(*s)
	}
	return pc
}

// SetPodIP sets the "pod_ip" field.
func (pc *PodCreate) SetPodIP(s string) *PodCreate {
	pc.mutation.SetPodIP(s)
	return pc
}

// SetNillablePodIP sets the "pod_ip" field if the given value is not nil.
func (pc *PodCreate) SetNillablePodIP(s *string) *PodCreate {
	if s != nil {
		pc.SetPodIP(*s)
	}
	return pc
}

// SetHostIP sets the "host_ip" field.
func (pc *PodCreate) SetHostIP(s string) *PodCreate {
	pc.mutation.SetHostIP(s)
	return pc
}

// SetNillableHostIP sets the "host_ip" field if the given value is not nil.
func (pc *PodCreate) SetNillableHostIP(s *string) *PodCreate {
	if s != nil {
		pc.SetHostIP(*s)
	}
	return pc
}

// SetStartTime sets the "start_time" field.
func (pc *PodCreate) SetStartTime(t time.Time) *PodCreate {
	pc.mutation.SetStartTime(t)
	return pc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (pc *PodCreate) SetNillableStartTime(t *time.Time) *PodCreate {
	if t != nil {
		pc.SetStartTime(*t)
	}
	return pc
}

// SetPhase sets the "phase" field.
func (pc *PodCreate) SetPhase(s string) *PodCreate {
	pc.mutation.SetPhase(s)
	return pc
}

// SetNillablePhase sets the "phase" field if the given value is not nil.
func (pc *PodCreate) SetNillablePhase(s *string) *PodCreate {
	if s != nil {
		pc.SetPhase(*s)
	}
	return pc
}

// SetReason sets the "reason" field.
func (pc *PodCreate) SetReason(s string) *PodCreate {
	pc.mutation.SetReason(s)
	return pc
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (pc *PodCreate) SetNillableReason(s *string) *PodCreate {
	if s != nil {
		pc.SetReason(*s)
	}
	return pc
}

// SetMessage sets the "message" field.
func (pc *PodCreate) SetMessage(s string) *PodCreate {
	pc.mutation.SetMessage(s)
	return pc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (pc *PodCreate) SetNillableMessage(s *string) *PodCreate {
	if s != nil {
		pc.SetMessage(*s)
	}
	return pc
}

// SetDetail sets the "detail" field.
func (pc *PodCreate) SetDetail(s string) *PodCreate {
	pc.mutation.SetDetail(s)
	return pc
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (pc *PodCreate) SetNillableDetail(s *string) *PodCreate {
	if s != nil {
		pc.SetDetail(*s)
	}
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *PodCreate) SetCreatedAt(t time.Time) *PodCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *PodCreate) SetNillableCreatedAt(t *time.Time) *PodCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PodCreate) SetUpdatedAt(t time.Time) *PodCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PodCreate) SetNillableUpdatedAt(t *time.Time) *PodCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PodCreate) SetID(i int64) *PodCreate {
	pc.mutation.SetID(i)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PodCreate) SetNillableID(i *int64) *PodCreate {
	if i != nil {
		pc.SetID(*i)
	}
	return pc
}

// SetServicetreeID sets the "servicetree" edge to the SpiderDevTblServicetree entity by ID.
func (pc *PodCreate) SetServicetreeID(id int32) *PodCreate {
	pc.mutation.SetServicetreeID(id)
	return pc
}

// SetNillableServicetreeID sets the "servicetree" edge to the SpiderDevTblServicetree entity by ID if the given value is not nil.
func (pc *PodCreate) SetNillableServicetreeID(id *int32) *PodCreate {
	if id != nil {
		pc = pc.SetServicetreeID(*id)
	}
	return pc
}

// SetServicetree sets the "servicetree" edge to the SpiderDevTblServicetree entity.
func (pc *PodCreate) SetServicetree(s *SpiderDevTblServicetree) *PodCreate {
	return pc.SetServicetreeID(s.ID)
}

// Mutation returns the PodMutation object of the builder.
func (pc *PodCreate) Mutation() *PodMutation {
	return pc.mutation
}

// Save creates the Pod in the database.
func (pc *PodCreate) Save(ctx context.Context) (*Pod, error) {
	var (
		err  error
		node *Pod
	)
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PodCreate) SaveX(ctx context.Context) *Pod {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PodCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PodCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PodCreate) check() error {
	return nil
}

func (pc *PodCreate) sqlSave(ctx context.Context) (*Pod, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (pc *PodCreate) createSpec() (*Pod, *sqlgraph.CreateSpec) {
	var (
		_node = &Pod{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pod.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: pod.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.ClusterName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldClusterName,
		})
		_node.ClusterName = value
	}
	if value, ok := pc.mutation.Namespace(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldNamespace,
		})
		_node.Namespace = value
	}
	if value, ok := pc.mutation.ServiceName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldServiceName,
		})
		_node.ServiceName = value
	}
	if value, ok := pc.mutation.PodName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPodName,
		})
		_node.PodName = value
	}
	if value, ok := pc.mutation.ResourceVersion(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldResourceVersion,
		})
		_node.ResourceVersion = value
	}
	if value, ok := pc.mutation.PodIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPodIP,
		})
		_node.PodIP = value
	}
	if value, ok := pc.mutation.HostIP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldHostIP,
		})
		_node.HostIP = value
	}
	if value, ok := pc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := pc.mutation.Phase(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldPhase,
		})
		_node.Phase = value
	}
	if value, ok := pc.mutation.Reason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldReason,
		})
		_node.Reason = value
	}
	if value, ok := pc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := pc.mutation.Detail(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pod.FieldDetail,
		})
		_node.Detail = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pod.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.ServicetreeIDs(); len(nodes) > 0 {
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
		_node.aname = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PodCreateBulk is the builder for creating many Pod entities in bulk.
type PodCreateBulk struct {
	config
	builders []*PodCreate
}

// Save creates the Pod entities in the database.
func (pcb *PodCreateBulk) Save(ctx context.Context) ([]*Pod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pod, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PodMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PodCreateBulk) SaveX(ctx context.Context) []*Pod {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PodCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PodCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

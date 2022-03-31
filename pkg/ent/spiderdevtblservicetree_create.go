// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hello/pkg/ent/spiderdevtblservicetree"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpiderDevTblServicetreeCreate is the builder for creating a SpiderDevTblServicetree entity.
type SpiderDevTblServicetreeCreate struct {
	config
	mutation *SpiderDevTblServicetreeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetName(s string) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetName(s)
	return sdtsc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillableName(s *string) *SpiderDevTblServicetreeCreate {
	if s != nil {
		sdtsc.SetName(*s)
	}
	return sdtsc
}

// SetAname sets the "aname" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetAname(s string) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetAname(s)
	return sdtsc
}

// SetNillableAname sets the "aname" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillableAname(s *string) *SpiderDevTblServicetreeCreate {
	if s != nil {
		sdtsc.SetAname(*s)
	}
	return sdtsc
}

// SetPnode sets the "pnode" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetPnode(i int32) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetPnode(i)
	return sdtsc
}

// SetNillablePnode sets the "pnode" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillablePnode(i *int32) *SpiderDevTblServicetreeCreate {
	if i != nil {
		sdtsc.SetPnode(*i)
	}
	return sdtsc
}

// SetType sets the "type" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetType(i int32) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetType(i)
	return sdtsc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillableType(i *int32) *SpiderDevTblServicetreeCreate {
	if i != nil {
		sdtsc.SetType(*i)
	}
	return sdtsc
}

// SetKey sets the "key" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetKey(s string) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetKey(s)
	return sdtsc
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillableKey(s *string) *SpiderDevTblServicetreeCreate {
	if s != nil {
		sdtsc.SetKey(*s)
	}
	return sdtsc
}

// SetOrigin sets the "origin" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetOrigin(s string) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetOrigin(s)
	return sdtsc
}

// SetNillableOrigin sets the "origin" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillableOrigin(s *string) *SpiderDevTblServicetreeCreate {
	if s != nil {
		sdtsc.SetOrigin(*s)
	}
	return sdtsc
}

// SetID sets the "id" field.
func (sdtsc *SpiderDevTblServicetreeCreate) SetID(i int32) *SpiderDevTblServicetreeCreate {
	sdtsc.mutation.SetID(i)
	return sdtsc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sdtsc *SpiderDevTblServicetreeCreate) SetNillableID(i *int32) *SpiderDevTblServicetreeCreate {
	if i != nil {
		sdtsc.SetID(*i)
	}
	return sdtsc
}

// Mutation returns the SpiderDevTblServicetreeMutation object of the builder.
func (sdtsc *SpiderDevTblServicetreeCreate) Mutation() *SpiderDevTblServicetreeMutation {
	return sdtsc.mutation
}

// Save creates the SpiderDevTblServicetree in the database.
func (sdtsc *SpiderDevTblServicetreeCreate) Save(ctx context.Context) (*SpiderDevTblServicetree, error) {
	var (
		err  error
		node *SpiderDevTblServicetree
	)
	if len(sdtsc.hooks) == 0 {
		if err = sdtsc.check(); err != nil {
			return nil, err
		}
		node, err = sdtsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpiderDevTblServicetreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sdtsc.check(); err != nil {
				return nil, err
			}
			sdtsc.mutation = mutation
			if node, err = sdtsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sdtsc.hooks) - 1; i >= 0; i-- {
			if sdtsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdtsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdtsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sdtsc *SpiderDevTblServicetreeCreate) SaveX(ctx context.Context) *SpiderDevTblServicetree {
	v, err := sdtsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdtsc *SpiderDevTblServicetreeCreate) Exec(ctx context.Context) error {
	_, err := sdtsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtsc *SpiderDevTblServicetreeCreate) ExecX(ctx context.Context) {
	if err := sdtsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sdtsc *SpiderDevTblServicetreeCreate) check() error {
	return nil
}

func (sdtsc *SpiderDevTblServicetreeCreate) sqlSave(ctx context.Context) (*SpiderDevTblServicetree, error) {
	_node, _spec := sdtsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sdtsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (sdtsc *SpiderDevTblServicetreeCreate) createSpec() (*SpiderDevTblServicetree, *sqlgraph.CreateSpec) {
	var (
		_node = &SpiderDevTblServicetree{config: sdtsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: spiderdevtblservicetree.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: spiderdevtblservicetree.FieldID,
			},
		}
	)
	if id, ok := sdtsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sdtsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: spiderdevtblservicetree.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sdtsc.mutation.Aname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: spiderdevtblservicetree.FieldAname,
		})
		_node.Aname = value
	}
	if value, ok := sdtsc.mutation.Pnode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: spiderdevtblservicetree.FieldPnode,
		})
		_node.Pnode = value
	}
	if value, ok := sdtsc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: spiderdevtblservicetree.FieldType,
		})
		_node.Type = value
	}
	if value, ok := sdtsc.mutation.Key(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: spiderdevtblservicetree.FieldKey,
		})
		_node.Key = value
	}
	if value, ok := sdtsc.mutation.Origin(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: spiderdevtblservicetree.FieldOrigin,
		})
		_node.Origin = value
	}
	return _node, _spec
}

// SpiderDevTblServicetreeCreateBulk is the builder for creating many SpiderDevTblServicetree entities in bulk.
type SpiderDevTblServicetreeCreateBulk struct {
	config
	builders []*SpiderDevTblServicetreeCreate
}

// Save creates the SpiderDevTblServicetree entities in the database.
func (sdtscb *SpiderDevTblServicetreeCreateBulk) Save(ctx context.Context) ([]*SpiderDevTblServicetree, error) {
	specs := make([]*sqlgraph.CreateSpec, len(sdtscb.builders))
	nodes := make([]*SpiderDevTblServicetree, len(sdtscb.builders))
	mutators := make([]Mutator, len(sdtscb.builders))
	for i := range sdtscb.builders {
		func(i int, root context.Context) {
			builder := sdtscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SpiderDevTblServicetreeMutation)
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
					_, err = mutators[i+1].Mutate(root, sdtscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sdtscb.driver, spec); err != nil {
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
					nodes[i].ID = int32(id)
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
		if _, err := mutators[0].Mutate(ctx, sdtscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sdtscb *SpiderDevTblServicetreeCreateBulk) SaveX(ctx context.Context) []*SpiderDevTblServicetree {
	v, err := sdtscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sdtscb *SpiderDevTblServicetreeCreateBulk) Exec(ctx context.Context) error {
	_, err := sdtscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtscb *SpiderDevTblServicetreeCreateBulk) ExecX(ctx context.Context) {
	if err := sdtscb.Exec(ctx); err != nil {
		panic(err)
	}
}
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"hello/pkg/ent/predicate"
	"hello/pkg/ent/spiderdevtblservicetree"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SpiderDevTblServicetreeDelete is the builder for deleting a SpiderDevTblServicetree entity.
type SpiderDevTblServicetreeDelete struct {
	config
	hooks    []Hook
	mutation *SpiderDevTblServicetreeMutation
}

// Where appends a list predicates to the SpiderDevTblServicetreeDelete builder.
func (sdtsd *SpiderDevTblServicetreeDelete) Where(ps ...predicate.SpiderDevTblServicetree) *SpiderDevTblServicetreeDelete {
	sdtsd.mutation.Where(ps...)
	return sdtsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdtsd *SpiderDevTblServicetreeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdtsd.hooks) == 0 {
		affected, err = sdtsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SpiderDevTblServicetreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdtsd.mutation = mutation
			affected, err = sdtsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdtsd.hooks) - 1; i >= 0; i-- {
			if sdtsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdtsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdtsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtsd *SpiderDevTblServicetreeDelete) ExecX(ctx context.Context) int {
	n, err := sdtsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdtsd *SpiderDevTblServicetreeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: spiderdevtblservicetree.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: spiderdevtblservicetree.FieldID,
			},
		},
	}
	if ps := sdtsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sdtsd.driver, _spec)
}

// SpiderDevTblServicetreeDeleteOne is the builder for deleting a single SpiderDevTblServicetree entity.
type SpiderDevTblServicetreeDeleteOne struct {
	sdtsd *SpiderDevTblServicetreeDelete
}

// Exec executes the deletion query.
func (sdtsdo *SpiderDevTblServicetreeDeleteOne) Exec(ctx context.Context) error {
	n, err := sdtsdo.sdtsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spiderdevtblservicetree.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdtsdo *SpiderDevTblServicetreeDeleteOne) ExecX(ctx context.Context) {
	sdtsdo.sdtsd.ExecX(ctx)
}
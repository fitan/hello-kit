// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hello/pkg/ent/project"
	"hello/pkg/ent/service"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ServiceCreate is the builder for creating a Service entity.
type ServiceCreate struct {
	config
	mutation *ServiceMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (sc *ServiceCreate) SetCreateTime(t time.Time) *ServiceCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableCreateTime(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *ServiceCreate) SetUpdateTime(t time.Time) *ServiceCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *ServiceCreate) SetNillableUpdateTime(t *time.Time) *ServiceCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *ServiceCreate) SetName(s string) *ServiceCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetAname sets the "aname" field.
func (sc *ServiceCreate) SetAname(s string) *ServiceCreate {
	sc.mutation.SetAname(s)
	return sc
}

// SetComments sets the "comments" field.
func (sc *ServiceCreate) SetComments(s string) *ServiceCreate {
	sc.mutation.SetComments(s)
	return sc
}

// SetClasses sets the "classes" field.
func (sc *ServiceCreate) SetClasses(s service.Classes) *ServiceCreate {
	sc.mutation.SetClasses(s)
	return sc
}

// SetLang sets the "lang" field.
func (sc *ServiceCreate) SetLang(s string) *ServiceCreate {
	sc.mutation.SetLang(s)
	return sc
}

// SetGit sets the "git" field.
func (sc *ServiceCreate) SetGit(s string) *ServiceCreate {
	sc.mutation.SetGit(s)
	return sc
}

// SetProjectID sets the "project" edge to the Project entity by ID.
func (sc *ServiceCreate) SetProjectID(id int) *ServiceCreate {
	sc.mutation.SetProjectID(id)
	return sc
}

// SetNillableProjectID sets the "project" edge to the Project entity by ID if the given value is not nil.
func (sc *ServiceCreate) SetNillableProjectID(id *int) *ServiceCreate {
	if id != nil {
		sc = sc.SetProjectID(*id)
	}
	return sc
}

// SetProject sets the "project" edge to the Project entity.
func (sc *ServiceCreate) SetProject(p *Project) *ServiceCreate {
	return sc.SetProjectID(p.ID)
}

// Mutation returns the ServiceMutation object of the builder.
func (sc *ServiceCreate) Mutation() *ServiceMutation {
	return sc.mutation
}

// Save creates the Service in the database.
func (sc *ServiceCreate) Save(ctx context.Context) (*Service, error) {
	var (
		err  error
		node *Service
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ServiceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ServiceCreate) SaveX(ctx context.Context) *Service {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ServiceCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ServiceCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ServiceCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := service.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := service.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ServiceCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Service.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Service.update_time"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Service.name"`)}
	}
	if _, ok := sc.mutation.Aname(); !ok {
		return &ValidationError{Name: "aname", err: errors.New(`ent: missing required field "Service.aname"`)}
	}
	if _, ok := sc.mutation.Comments(); !ok {
		return &ValidationError{Name: "comments", err: errors.New(`ent: missing required field "Service.comments"`)}
	}
	if _, ok := sc.mutation.Classes(); !ok {
		return &ValidationError{Name: "classes", err: errors.New(`ent: missing required field "Service.classes"`)}
	}
	if v, ok := sc.mutation.Classes(); ok {
		if err := service.ClassesValidator(v); err != nil {
			return &ValidationError{Name: "classes", err: fmt.Errorf(`ent: validator failed for field "Service.classes": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Lang(); !ok {
		return &ValidationError{Name: "lang", err: errors.New(`ent: missing required field "Service.lang"`)}
	}
	if _, ok := sc.mutation.Git(); !ok {
		return &ValidationError{Name: "git", err: errors.New(`ent: missing required field "Service.git"`)}
	}
	return nil
}

func (sc *ServiceCreate) sqlSave(ctx context.Context) (*Service, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ServiceCreate) createSpec() (*Service, *sqlgraph.CreateSpec) {
	var (
		_node = &Service{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: service.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: service.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: service.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: service.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: service.FieldName,
		})
		_node.Name = value
	}
	if value, ok := sc.mutation.Aname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: service.FieldAname,
		})
		_node.Aname = value
	}
	if value, ok := sc.mutation.Comments(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: service.FieldComments,
		})
		_node.Comments = value
	}
	if value, ok := sc.mutation.Classes(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: service.FieldClasses,
		})
		_node.Classes = value
	}
	if value, ok := sc.mutation.Lang(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: service.FieldLang,
		})
		_node.Lang = value
	}
	if value, ok := sc.mutation.Git(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: service.FieldGit,
		})
		_node.Git = value
	}
	if nodes := sc.mutation.ProjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   service.ProjectTable,
			Columns: []string{service.ProjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: project.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.project_services = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ServiceCreateBulk is the builder for creating many Service entities in bulk.
type ServiceCreateBulk struct {
	config
	builders []*ServiceCreate
}

// Save creates the Service entities in the database.
func (scb *ServiceCreateBulk) Save(ctx context.Context) ([]*Service, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Service, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ServiceMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ServiceCreateBulk) SaveX(ctx context.Context) []*Service {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ServiceCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ServiceCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
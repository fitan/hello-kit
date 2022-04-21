// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"hello/pkg/ent/predicate"
	"hello/pkg/ent/project"
	"math"
	"reflect"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProjectQuery is the builder for querying Project entities.
type ProjectQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Project
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProjectQuery builder.
func (pq *ProjectQuery) Where(ps ...predicate.Project) *ProjectQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *ProjectQuery) Limit(limit int) *ProjectQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *ProjectQuery) Offset(offset int) *ProjectQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProjectQuery) Unique(unique bool) *ProjectQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *ProjectQuery) Order(o ...OrderFunc) *ProjectQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// First returns the first Project entity from the query.
// Returns a *NotFoundError when no Project was found.
func (pq *ProjectQuery) First(ctx context.Context) (*Project, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{project.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProjectQuery) FirstX(ctx context.Context) *Project {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Project ID from the query.
// Returns a *NotFoundError when no Project ID was found.
func (pq *ProjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{project.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProjectQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Project entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Project entity is found.
// Returns a *NotFoundError when no Project entities are found.
func (pq *ProjectQuery) Only(ctx context.Context) (*Project, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{project.Label}
	default:
		return nil, &NotSingularError{project.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProjectQuery) OnlyX(ctx context.Context) *Project {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Project ID in the query.
// Returns a *NotSingularError when more than one Project ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = &NotSingularError{project.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Projects.
func (pq *ProjectQuery) All(ctx context.Context) ([]*Project, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProjectQuery) AllX(ctx context.Context) []*Project {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Project IDs.
func (pq *ProjectQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(project.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProjectQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProjectQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProjectQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProjectQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProjectQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProjectQuery) Clone() *ProjectQuery {
	if pq == nil {
		return nil
	}
	return &ProjectQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		predicates: append([]predicate.Project{}, pq.predicates...),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Alias string `json:"alias,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Project.Query().
//		GroupBy(project.FieldAlias).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *ProjectQuery) GroupBy(field string, fields ...string) *ProjectGroupBy {
	group := &ProjectGroupBy{config: pq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Alias string `json:"alias,omitempty"`
//	}
//
//	client.Project.Query().
//		Select(project.FieldAlias).
//		Scan(ctx, &v)
//
func (pq *ProjectQuery) Select(fields ...string) *ProjectSelect {
	pq.fields = append(pq.fields, fields...)
	return &ProjectSelect{ProjectQuery: pq}
}

func (pq *ProjectQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !project.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProjectQuery) sqlAll(ctx context.Context) ([]*Project, error) {
	var (
		nodes = []*Project{}
		_spec = pq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Project{config: pq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pq *ProjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProjectQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *ProjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   project.Table,
			Columns: project.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: project.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for i := range fields {
			if fields[i] != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(project.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = project.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

func (pq *ProjectQuery) ByQueries(ctx context.Context, i interface{}) (res Projects, count int, err error) {
	queryList, countList := SetProjectFormQueries(i)
	countQ := pq.Clone()
	for _, v := range queryList {
		v.Query(pq)
	}
	for _, v := range countList {
		v.Query(countQ)
	}
	count, err = countQ.Count(ctx)
	if err != nil {
		return
	}
	res, err = pq.All(ctx)
	return
}

type ProjectTableFormer interface {
	Query(q *ProjectQuery)
	CountQuery() bool
}

type ProjectTablePagingForm struct {
	Limit *int `json:"_limit" form:"_limit" `
	Page  *int `json:"_page" form:"_page"`
}

func (f ProjectTablePagingForm) Query(q *ProjectQuery) {
	if f.Limit != nil && f.Page != nil {
		q.Limit(*f.Limit).Offset((*f.Page - 1) * *f.Limit)
	}
}

func (f ProjectTablePagingForm) CountQuery() bool {
	return false
}

type ProjectTableOrderForm struct {
	Order  *string `json:"order" form:"_order" binding:"omitempty,oneof=acs desc"`
	SortBy *string `json:"sortBy" form:"_sortBy"`
}

func (f ProjectTableOrderForm) Query(q *ProjectQuery) {
	if f.Order != nil && f.SortBy != nil {
		if *f.Order == "acs" {
			q.Order(Asc(*f.SortBy))
		}

		if *f.Order == "desc" {
			q.Order(Desc(*f.SortBy))
		}
	}
}
func (f ProjectTableOrderForm) CountQuery() bool {
	return false
}

func SetProjectFormQueries(o interface{}) ([]ProjectTableFormer, []ProjectTableFormer) {
	queryList := make([]ProjectTableFormer, 0)
	countList := make([]ProjectTableFormer, 0)
	v := reflect.ValueOf(o)
	former := reflect.TypeOf((*ProjectTableFormer)(nil)).Elem()
	ProjectFormDepValue(v, former, &queryList, &countList)
	return queryList, countList
}

func ProjectFormDepValue(v reflect.Value, former reflect.Type, queryList *[]ProjectTableFormer, countList *[]ProjectTableFormer) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.IsZero() {
			continue
		}
		if f.Type().Implements(former) {
			former := f.Interface().(ProjectTableFormer)
			*queryList = append(*queryList, former)
			if former.CountQuery() {
				*countList = append(*countList, former)
			}
			continue
		}
		if f.Type().Kind() == reflect.Struct {
			ProjectFormDepValue(f, former, queryList, countList)
		}
	}
}

type ProjectQueryOps struct {
}

type ProjectTableAliasEQForm struct {
	AliasEQ *string `form:"AliasEQ" json:"AliasEQ"`
}

func (f ProjectTableAliasEQForm) Query(q *ProjectQuery) {
	if f.AliasEQ != nil {
		q.Where(project.AliasEQ(*f.AliasEQ))
	}
}
func (f ProjectTableAliasEQForm) CountQuery() bool {
	return true
}

type ProjectTableAliasNEQForm struct {
	AliasNEQ *string `form:"AliasNEQ" json:"AliasNEQ"`
}

func (f ProjectTableAliasNEQForm) Query(q *ProjectQuery) {
	if f.AliasNEQ != nil {
		q.Where(project.AliasNEQ(*f.AliasNEQ))
	}
}
func (f ProjectTableAliasNEQForm) CountQuery() bool {
	return true
}

type ProjectTableAliasInForm struct {
	AliasIn *[]string `form:"AliasIn" json:"AliasIn"`
}

func (f ProjectTableAliasInForm) Query(q *ProjectQuery) {
	if f.AliasIn != nil {
		q.Where(project.AliasIn(*f.AliasIn...))
	}
}
func (f ProjectTableAliasInForm) CountQuery() bool {
	return true
}

type ProjectTableAliasNotInForm struct {
	AliasNotIn *[]string `form:"AliasNotIn" json:"AliasNotIn"`
}

func (f ProjectTableAliasNotInForm) Query(q *ProjectQuery) {
	if f.AliasNotIn != nil {
		q.Where(project.AliasNotIn(*f.AliasNotIn...))
	}
}
func (f ProjectTableAliasNotInForm) CountQuery() bool {
	return true
}

type ProjectTableAliasGTForm struct {
	AliasGT *string `form:"AliasGT" json:"AliasGT"`
}

func (f ProjectTableAliasGTForm) Query(q *ProjectQuery) {
	if f.AliasGT != nil {
		q.Where(project.AliasGT(*f.AliasGT))
	}
}
func (f ProjectTableAliasGTForm) CountQuery() bool {
	return true
}

type ProjectTableAliasGTEForm struct {
	AliasGTE *string `form:"AliasGTE" json:"AliasGTE"`
}

func (f ProjectTableAliasGTEForm) Query(q *ProjectQuery) {
	if f.AliasGTE != nil {
		q.Where(project.AliasGTE(*f.AliasGTE))
	}
}
func (f ProjectTableAliasGTEForm) CountQuery() bool {
	return true
}

type ProjectTableAliasLTForm struct {
	AliasLT *string `form:"AliasLT" json:"AliasLT"`
}

func (f ProjectTableAliasLTForm) Query(q *ProjectQuery) {
	if f.AliasLT != nil {
		q.Where(project.AliasLT(*f.AliasLT))
	}
}
func (f ProjectTableAliasLTForm) CountQuery() bool {
	return true
}

type ProjectTableAliasLTEForm struct {
	AliasLTE *string `form:"AliasLTE" json:"AliasLTE"`
}

func (f ProjectTableAliasLTEForm) Query(q *ProjectQuery) {
	if f.AliasLTE != nil {
		q.Where(project.AliasLTE(*f.AliasLTE))
	}
}
func (f ProjectTableAliasLTEForm) CountQuery() bool {
	return true
}

type ProjectTableAliasContainsForm struct {
	AliasContains *string `form:"AliasContains" json:"AliasContains"`
}

func (f ProjectTableAliasContainsForm) Query(q *ProjectQuery) {
	if f.AliasContains != nil {
		q.Where(project.AliasContains(*f.AliasContains))
	}
}
func (f ProjectTableAliasContainsForm) CountQuery() bool {
	return true
}

type ProjectTableAliasHasPrefixForm struct {
	AliasHasPrefix *string `form:"AliasHasPrefix" json:"AliasHasPrefix"`
}

func (f ProjectTableAliasHasPrefixForm) Query(q *ProjectQuery) {
	if f.AliasHasPrefix != nil {
		q.Where(project.AliasHasPrefix(*f.AliasHasPrefix))
	}
}
func (f ProjectTableAliasHasPrefixForm) CountQuery() bool {
	return true
}

type ProjectTableAliasHasSuffixForm struct {
	AliasHasSuffix *string `form:"AliasHasSuffix" json:"AliasHasSuffix"`
}

func (f ProjectTableAliasHasSuffixForm) Query(q *ProjectQuery) {
	if f.AliasHasSuffix != nil {
		q.Where(project.AliasHasSuffix(*f.AliasHasSuffix))
	}
}
func (f ProjectTableAliasHasSuffixForm) CountQuery() bool {
	return true
}

type ProjectTableAliasEqualFoldForm struct {
	AliasEqualFold *string `form:"AliasEqualFold" json:"AliasEqualFold"`
}

func (f ProjectTableAliasEqualFoldForm) Query(q *ProjectQuery) {
	if f.AliasEqualFold != nil {
		q.Where(project.AliasEqualFold(*f.AliasEqualFold))
	}
}
func (f ProjectTableAliasEqualFoldForm) CountQuery() bool {
	return true
}

type ProjectTableAliasContainsFoldForm struct {
	AliasContainsFold *string `form:"AliasContainsFold" json:"AliasContainsFold"`
}

func (f ProjectTableAliasContainsFoldForm) Query(q *ProjectQuery) {
	if f.AliasContainsFold != nil {
		q.Where(project.AliasContainsFold(*f.AliasContainsFold))
	}
}
func (f ProjectTableAliasContainsFoldForm) CountQuery() bool {
	return true
}

type ProjectTableNameEQForm struct {
	NameEQ *string `form:"NameEQ" json:"NameEQ"`
}

func (f ProjectTableNameEQForm) Query(q *ProjectQuery) {
	if f.NameEQ != nil {
		q.Where(project.NameEQ(*f.NameEQ))
	}
}
func (f ProjectTableNameEQForm) CountQuery() bool {
	return true
}

type ProjectTableNameNEQForm struct {
	NameNEQ *string `form:"NameNEQ" json:"NameNEQ"`
}

func (f ProjectTableNameNEQForm) Query(q *ProjectQuery) {
	if f.NameNEQ != nil {
		q.Where(project.NameNEQ(*f.NameNEQ))
	}
}
func (f ProjectTableNameNEQForm) CountQuery() bool {
	return true
}

type ProjectTableNameInForm struct {
	NameIn *[]string `form:"NameIn" json:"NameIn"`
}

func (f ProjectTableNameInForm) Query(q *ProjectQuery) {
	if f.NameIn != nil {
		q.Where(project.NameIn(*f.NameIn...))
	}
}
func (f ProjectTableNameInForm) CountQuery() bool {
	return true
}

type ProjectTableNameNotInForm struct {
	NameNotIn *[]string `form:"NameNotIn" json:"NameNotIn"`
}

func (f ProjectTableNameNotInForm) Query(q *ProjectQuery) {
	if f.NameNotIn != nil {
		q.Where(project.NameNotIn(*f.NameNotIn...))
	}
}
func (f ProjectTableNameNotInForm) CountQuery() bool {
	return true
}

type ProjectTableNameGTForm struct {
	NameGT *string `form:"NameGT" json:"NameGT"`
}

func (f ProjectTableNameGTForm) Query(q *ProjectQuery) {
	if f.NameGT != nil {
		q.Where(project.NameGT(*f.NameGT))
	}
}
func (f ProjectTableNameGTForm) CountQuery() bool {
	return true
}

type ProjectTableNameGTEForm struct {
	NameGTE *string `form:"NameGTE" json:"NameGTE"`
}

func (f ProjectTableNameGTEForm) Query(q *ProjectQuery) {
	if f.NameGTE != nil {
		q.Where(project.NameGTE(*f.NameGTE))
	}
}
func (f ProjectTableNameGTEForm) CountQuery() bool {
	return true
}

type ProjectTableNameLTForm struct {
	NameLT *string `form:"NameLT" json:"NameLT"`
}

func (f ProjectTableNameLTForm) Query(q *ProjectQuery) {
	if f.NameLT != nil {
		q.Where(project.NameLT(*f.NameLT))
	}
}
func (f ProjectTableNameLTForm) CountQuery() bool {
	return true
}

type ProjectTableNameLTEForm struct {
	NameLTE *string `form:"NameLTE" json:"NameLTE"`
}

func (f ProjectTableNameLTEForm) Query(q *ProjectQuery) {
	if f.NameLTE != nil {
		q.Where(project.NameLTE(*f.NameLTE))
	}
}
func (f ProjectTableNameLTEForm) CountQuery() bool {
	return true
}

type ProjectTableNameContainsForm struct {
	NameContains *string `form:"NameContains" json:"NameContains"`
}

func (f ProjectTableNameContainsForm) Query(q *ProjectQuery) {
	if f.NameContains != nil {
		q.Where(project.NameContains(*f.NameContains))
	}
}
func (f ProjectTableNameContainsForm) CountQuery() bool {
	return true
}

type ProjectTableNameHasPrefixForm struct {
	NameHasPrefix *string `form:"NameHasPrefix" json:"NameHasPrefix"`
}

func (f ProjectTableNameHasPrefixForm) Query(q *ProjectQuery) {
	if f.NameHasPrefix != nil {
		q.Where(project.NameHasPrefix(*f.NameHasPrefix))
	}
}
func (f ProjectTableNameHasPrefixForm) CountQuery() bool {
	return true
}

type ProjectTableNameHasSuffixForm struct {
	NameHasSuffix *string `form:"NameHasSuffix" json:"NameHasSuffix"`
}

func (f ProjectTableNameHasSuffixForm) Query(q *ProjectQuery) {
	if f.NameHasSuffix != nil {
		q.Where(project.NameHasSuffix(*f.NameHasSuffix))
	}
}
func (f ProjectTableNameHasSuffixForm) CountQuery() bool {
	return true
}

type ProjectTableNameEqualFoldForm struct {
	NameEqualFold *string `form:"NameEqualFold" json:"NameEqualFold"`
}

func (f ProjectTableNameEqualFoldForm) Query(q *ProjectQuery) {
	if f.NameEqualFold != nil {
		q.Where(project.NameEqualFold(*f.NameEqualFold))
	}
}
func (f ProjectTableNameEqualFoldForm) CountQuery() bool {
	return true
}

type ProjectTableNameContainsFoldForm struct {
	NameContainsFold *string `form:"NameContainsFold" json:"NameContainsFold"`
}

func (f ProjectTableNameContainsFoldForm) Query(q *ProjectQuery) {
	if f.NameContainsFold != nil {
		q.Where(project.NameContainsFold(*f.NameContainsFold))
	}
}
func (f ProjectTableNameContainsFoldForm) CountQuery() bool {
	return true
}

// ProjectGroupBy is the group-by builder for Project entities.
type ProjectGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProjectGroupBy) Aggregate(fns ...AggregateFunc) *ProjectGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *ProjectGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (pgb *ProjectGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := pgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProjectGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (pgb *ProjectGroupBy) StringsX(ctx context.Context) []string {
	v, err := pgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = pgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (pgb *ProjectGroupBy) StringX(ctx context.Context) string {
	v, err := pgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProjectGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (pgb *ProjectGroupBy) IntsX(ctx context.Context) []int {
	v, err := pgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = pgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (pgb *ProjectGroupBy) IntX(ctx context.Context) int {
	v, err := pgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProjectGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (pgb *ProjectGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := pgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = pgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (pgb *ProjectGroupBy) Float64X(ctx context.Context) float64 {
	v, err := pgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(pgb.fields) > 1 {
		return nil, errors.New("ent: ProjectGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := pgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (pgb *ProjectGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := pgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (pgb *ProjectGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = pgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (pgb *ProjectGroupBy) BoolX(ctx context.Context) bool {
	v, err := pgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pgb *ProjectGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !project.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *ProjectGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// ProjectSelect is the builder for selecting fields of Project entities.
type ProjectSelect struct {
	*ProjectQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProjectSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.ProjectQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ps *ProjectSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProjectSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ps *ProjectSelect) StringsX(ctx context.Context) []string {
	v, err := ps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ps *ProjectSelect) StringX(ctx context.Context) string {
	v, err := ps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProjectSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ps *ProjectSelect) IntsX(ctx context.Context) []int {
	v, err := ps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ps *ProjectSelect) IntX(ctx context.Context) int {
	v, err := ps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProjectSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ps *ProjectSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ps *ProjectSelect) Float64X(ctx context.Context) float64 {
	v, err := ps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ps.fields) > 1 {
		return nil, errors.New("ent: ProjectSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ps *ProjectSelect) BoolsX(ctx context.Context) []bool {
	v, err := ps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ps *ProjectSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{project.Label}
	default:
		err = fmt.Errorf("ent: ProjectSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ps *ProjectSelect) BoolX(ctx context.Context) bool {
	v, err := ps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ps *ProjectSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

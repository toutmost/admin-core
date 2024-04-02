// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/toutmost/admin-core/rpc/ent/department"
	"github.com/toutmost/admin-core/rpc/ent/predicate"
)

// DepartmentDelete is the builder for deleting a Department entity.
type DepartmentDelete struct {
	config
	hooks    []Hook
	mutation *DepartmentMutation
}

// Where appends a list predicates to the DepartmentDelete builder.
func (dd *DepartmentDelete) Where(ps ...predicate.Department) *DepartmentDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DepartmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DepartmentDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DepartmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(department.Table, sqlgraph.NewFieldSpec(department.FieldID, field.TypeUint64))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DepartmentDeleteOne is the builder for deleting a single Department entity.
type DepartmentDeleteOne struct {
	dd *DepartmentDelete
}

// Where appends a list predicates to the DepartmentDelete builder.
func (ddo *DepartmentDeleteOne) Where(ps ...predicate.Department) *DepartmentDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DepartmentDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{department.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DepartmentDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}

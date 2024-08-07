// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/compatibilityresult"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/predicate"
)

// CompatibilityResultDelete is the builder for deleting a CompatibilityResult entity.
type CompatibilityResultDelete struct {
	config
	hooks    []Hook
	mutation *CompatibilityResultMutation
}

// Where appends a list predicates to the CompatibilityResultDelete builder.
func (crd *CompatibilityResultDelete) Where(ps ...predicate.CompatibilityResult) *CompatibilityResultDelete {
	crd.mutation.Where(ps...)
	return crd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crd *CompatibilityResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crd.sqlExec, crd.mutation, crd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crd *CompatibilityResultDelete) ExecX(ctx context.Context) int {
	n, err := crd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crd *CompatibilityResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(compatibilityresult.Table, sqlgraph.NewFieldSpec(compatibilityresult.FieldID, field.TypeInt))
	if ps := crd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crd.mutation.done = true
	return affected, err
}

// CompatibilityResultDeleteOne is the builder for deleting a single CompatibilityResult entity.
type CompatibilityResultDeleteOne struct {
	crd *CompatibilityResultDelete
}

// Where appends a list predicates to the CompatibilityResultDelete builder.
func (crdo *CompatibilityResultDeleteOne) Where(ps ...predicate.CompatibilityResult) *CompatibilityResultDeleteOne {
	crdo.crd.mutation.Where(ps...)
	return crdo
}

// Exec executes the deletion query.
func (crdo *CompatibilityResultDeleteOne) Exec(ctx context.Context) error {
	n, err := crdo.crd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{compatibilityresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crdo *CompatibilityResultDeleteOne) ExecX(ctx context.Context) {
	if err := crdo.Exec(ctx); err != nil {
		panic(err)
	}
}

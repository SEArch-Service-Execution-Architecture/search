// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/clpombo/search/ent/predicate"
	"github.com/clpombo/search/ent/registeredcontract"
)

// RegisteredContractDelete is the builder for deleting a RegisteredContract entity.
type RegisteredContractDelete struct {
	config
	hooks    []Hook
	mutation *RegisteredContractMutation
}

// Where appends a list predicates to the RegisteredContractDelete builder.
func (rcd *RegisteredContractDelete) Where(ps ...predicate.RegisteredContract) *RegisteredContractDelete {
	rcd.mutation.Where(ps...)
	return rcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rcd *RegisteredContractDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, RegisteredContractMutation](ctx, rcd.sqlExec, rcd.mutation, rcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rcd *RegisteredContractDelete) ExecX(ctx context.Context) int {
	n, err := rcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rcd *RegisteredContractDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(registeredcontract.Table, sqlgraph.NewFieldSpec(registeredcontract.FieldID, field.TypeString))
	if ps := rcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rcd.mutation.done = true
	return affected, err
}

// RegisteredContractDeleteOne is the builder for deleting a single RegisteredContract entity.
type RegisteredContractDeleteOne struct {
	rcd *RegisteredContractDelete
}

// Where appends a list predicates to the RegisteredContractDelete builder.
func (rcdo *RegisteredContractDeleteOne) Where(ps ...predicate.RegisteredContract) *RegisteredContractDeleteOne {
	rcdo.rcd.mutation.Where(ps...)
	return rcdo
}

// Exec executes the deletion query.
func (rcdo *RegisteredContractDeleteOne) Exec(ctx context.Context) error {
	n, err := rcdo.rcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{registeredcontract.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rcdo *RegisteredContractDeleteOne) ExecX(ctx context.Context) {
	if err := rcdo.Exec(ctx); err != nil {
		panic(err)
	}
}

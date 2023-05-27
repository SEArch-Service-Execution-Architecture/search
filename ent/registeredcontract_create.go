// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/pmontepagano/search/ent/compatibilityresult"
	"github.com/pmontepagano/search/ent/registeredcontract"
	"github.com/pmontepagano/search/ent/registeredprovider"
)

// RegisteredContractCreate is the builder for creating a RegisteredContract entity.
type RegisteredContractCreate struct {
	config
	mutation *RegisteredContractMutation
	hooks    []Hook
}

// SetFormat sets the "format" field.
func (rcc *RegisteredContractCreate) SetFormat(i int) *RegisteredContractCreate {
	rcc.mutation.SetFormat(i)
	return rcc
}

// SetContract sets the "contract" field.
func (rcc *RegisteredContractCreate) SetContract(b []byte) *RegisteredContractCreate {
	rcc.mutation.SetContract(b)
	return rcc
}

// SetCreatedAt sets the "created_at" field.
func (rcc *RegisteredContractCreate) SetCreatedAt(t time.Time) *RegisteredContractCreate {
	rcc.mutation.SetCreatedAt(t)
	return rcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rcc *RegisteredContractCreate) SetNillableCreatedAt(t *time.Time) *RegisteredContractCreate {
	if t != nil {
		rcc.SetCreatedAt(*t)
	}
	return rcc
}

// SetID sets the "id" field.
func (rcc *RegisteredContractCreate) SetID(s string) *RegisteredContractCreate {
	rcc.mutation.SetID(s)
	return rcc
}

// AddProviderIDs adds the "providers" edge to the RegisteredProvider entity by IDs.
func (rcc *RegisteredContractCreate) AddProviderIDs(ids ...uuid.UUID) *RegisteredContractCreate {
	rcc.mutation.AddProviderIDs(ids...)
	return rcc
}

// AddProviders adds the "providers" edges to the RegisteredProvider entity.
func (rcc *RegisteredContractCreate) AddProviders(r ...*RegisteredProvider) *RegisteredContractCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rcc.AddProviderIDs(ids...)
}

// AddCompatibilityResultsAsRequirementIDs adds the "compatibility_results_as_requirement" edge to the CompatibilityResult entity by IDs.
func (rcc *RegisteredContractCreate) AddCompatibilityResultsAsRequirementIDs(ids ...int) *RegisteredContractCreate {
	rcc.mutation.AddCompatibilityResultsAsRequirementIDs(ids...)
	return rcc
}

// AddCompatibilityResultsAsRequirement adds the "compatibility_results_as_requirement" edges to the CompatibilityResult entity.
func (rcc *RegisteredContractCreate) AddCompatibilityResultsAsRequirement(c ...*CompatibilityResult) *RegisteredContractCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rcc.AddCompatibilityResultsAsRequirementIDs(ids...)
}

// AddCompatibilityResultsAsProviderIDs adds the "compatibility_results_as_provider" edge to the CompatibilityResult entity by IDs.
func (rcc *RegisteredContractCreate) AddCompatibilityResultsAsProviderIDs(ids ...int) *RegisteredContractCreate {
	rcc.mutation.AddCompatibilityResultsAsProviderIDs(ids...)
	return rcc
}

// AddCompatibilityResultsAsProvider adds the "compatibility_results_as_provider" edges to the CompatibilityResult entity.
func (rcc *RegisteredContractCreate) AddCompatibilityResultsAsProvider(c ...*CompatibilityResult) *RegisteredContractCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return rcc.AddCompatibilityResultsAsProviderIDs(ids...)
}

// Mutation returns the RegisteredContractMutation object of the builder.
func (rcc *RegisteredContractCreate) Mutation() *RegisteredContractMutation {
	return rcc.mutation
}

// Save creates the RegisteredContract in the database.
func (rcc *RegisteredContractCreate) Save(ctx context.Context) (*RegisteredContract, error) {
	rcc.defaults()
	return withHooks[*RegisteredContract, RegisteredContractMutation](ctx, rcc.sqlSave, rcc.mutation, rcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rcc *RegisteredContractCreate) SaveX(ctx context.Context) *RegisteredContract {
	v, err := rcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcc *RegisteredContractCreate) Exec(ctx context.Context) error {
	_, err := rcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcc *RegisteredContractCreate) ExecX(ctx context.Context) {
	if err := rcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rcc *RegisteredContractCreate) defaults() {
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		v := registeredcontract.DefaultCreatedAt()
		rcc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rcc *RegisteredContractCreate) check() error {
	if _, ok := rcc.mutation.Format(); !ok {
		return &ValidationError{Name: "format", err: errors.New(`ent: missing required field "RegisteredContract.format"`)}
	}
	if _, ok := rcc.mutation.Contract(); !ok {
		return &ValidationError{Name: "contract", err: errors.New(`ent: missing required field "RegisteredContract.contract"`)}
	}
	if v, ok := rcc.mutation.Contract(); ok {
		if err := registeredcontract.ContractValidator(v); err != nil {
			return &ValidationError{Name: "contract", err: fmt.Errorf(`ent: validator failed for field "RegisteredContract.contract": %w`, err)}
		}
	}
	if _, ok := rcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RegisteredContract.created_at"`)}
	}
	if v, ok := rcc.mutation.ID(); ok {
		if err := registeredcontract.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "RegisteredContract.id": %w`, err)}
		}
	}
	return nil
}

func (rcc *RegisteredContractCreate) sqlSave(ctx context.Context) (*RegisteredContract, error) {
	if err := rcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RegisteredContract.ID type: %T", _spec.ID.Value)
		}
	}
	rcc.mutation.id = &_node.ID
	rcc.mutation.done = true
	return _node, nil
}

func (rcc *RegisteredContractCreate) createSpec() (*RegisteredContract, *sqlgraph.CreateSpec) {
	var (
		_node = &RegisteredContract{config: rcc.config}
		_spec = sqlgraph.NewCreateSpec(registeredcontract.Table, sqlgraph.NewFieldSpec(registeredcontract.FieldID, field.TypeString))
	)
	if id, ok := rcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rcc.mutation.Format(); ok {
		_spec.SetField(registeredcontract.FieldFormat, field.TypeInt, value)
		_node.Format = value
	}
	if value, ok := rcc.mutation.Contract(); ok {
		_spec.SetField(registeredcontract.FieldContract, field.TypeBytes, value)
		_node.Contract = value
	}
	if value, ok := rcc.mutation.CreatedAt(); ok {
		_spec.SetField(registeredcontract.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := rcc.mutation.ProvidersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   registeredcontract.ProvidersTable,
			Columns: []string{registeredcontract.ProvidersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registeredprovider.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.CompatibilityResultsAsRequirementIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   registeredcontract.CompatibilityResultsAsRequirementTable,
			Columns: []string{registeredcontract.CompatibilityResultsAsRequirementColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(compatibilityresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rcc.mutation.CompatibilityResultsAsProviderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   registeredcontract.CompatibilityResultsAsProviderTable,
			Columns: []string{registeredcontract.CompatibilityResultsAsProviderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(compatibilityresult.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RegisteredContractCreateBulk is the builder for creating many RegisteredContract entities in bulk.
type RegisteredContractCreateBulk struct {
	config
	builders []*RegisteredContractCreate
}

// Save creates the RegisteredContract entities in the database.
func (rccb *RegisteredContractCreateBulk) Save(ctx context.Context) ([]*RegisteredContract, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rccb.builders))
	nodes := make([]*RegisteredContract, len(rccb.builders))
	mutators := make([]Mutator, len(rccb.builders))
	for i := range rccb.builders {
		func(i int, root context.Context) {
			builder := rccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RegisteredContractMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rccb *RegisteredContractCreateBulk) SaveX(ctx context.Context) []*RegisteredContract {
	v, err := rccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rccb *RegisteredContractCreateBulk) Exec(ctx context.Context) error {
	_, err := rccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rccb *RegisteredContractCreateBulk) ExecX(ctx context.Context) {
	if err := rccb.Exec(ctx); err != nil {
		panic(err)
	}
}

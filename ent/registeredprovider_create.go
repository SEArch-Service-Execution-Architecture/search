// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/pmontepagano/search/ent/registeredcontract"
	"github.com/pmontepagano/search/ent/registeredprovider"
)

// RegisteredProviderCreate is the builder for creating a RegisteredProvider entity.
type RegisteredProviderCreate struct {
	config
	mutation *RegisteredProviderMutation
	hooks    []Hook
}

// SetURL sets the "url" field.
func (rpc *RegisteredProviderCreate) SetURL(u *url.URL) *RegisteredProviderCreate {
	rpc.mutation.SetURL(u)
	return rpc
}

// SetContractID sets the "contract_id" field.
func (rpc *RegisteredProviderCreate) SetContractID(s string) *RegisteredProviderCreate {
	rpc.mutation.SetContractID(s)
	return rpc
}

// SetCreatedAt sets the "created_at" field.
func (rpc *RegisteredProviderCreate) SetCreatedAt(t time.Time) *RegisteredProviderCreate {
	rpc.mutation.SetCreatedAt(t)
	return rpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rpc *RegisteredProviderCreate) SetNillableCreatedAt(t *time.Time) *RegisteredProviderCreate {
	if t != nil {
		rpc.SetCreatedAt(*t)
	}
	return rpc
}

// SetUpdatedAt sets the "updated_at" field.
func (rpc *RegisteredProviderCreate) SetUpdatedAt(t time.Time) *RegisteredProviderCreate {
	rpc.mutation.SetUpdatedAt(t)
	return rpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rpc *RegisteredProviderCreate) SetNillableUpdatedAt(t *time.Time) *RegisteredProviderCreate {
	if t != nil {
		rpc.SetUpdatedAt(*t)
	}
	return rpc
}

// SetID sets the "id" field.
func (rpc *RegisteredProviderCreate) SetID(u uuid.UUID) *RegisteredProviderCreate {
	rpc.mutation.SetID(u)
	return rpc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rpc *RegisteredProviderCreate) SetNillableID(u *uuid.UUID) *RegisteredProviderCreate {
	if u != nil {
		rpc.SetID(*u)
	}
	return rpc
}

// SetContract sets the "contract" edge to the RegisteredContract entity.
func (rpc *RegisteredProviderCreate) SetContract(r *RegisteredContract) *RegisteredProviderCreate {
	return rpc.SetContractID(r.ID)
}

// Mutation returns the RegisteredProviderMutation object of the builder.
func (rpc *RegisteredProviderCreate) Mutation() *RegisteredProviderMutation {
	return rpc.mutation
}

// Save creates the RegisteredProvider in the database.
func (rpc *RegisteredProviderCreate) Save(ctx context.Context) (*RegisteredProvider, error) {
	rpc.defaults()
	return withHooks(ctx, rpc.sqlSave, rpc.mutation, rpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rpc *RegisteredProviderCreate) SaveX(ctx context.Context) *RegisteredProvider {
	v, err := rpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpc *RegisteredProviderCreate) Exec(ctx context.Context) error {
	_, err := rpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpc *RegisteredProviderCreate) ExecX(ctx context.Context) {
	if err := rpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rpc *RegisteredProviderCreate) defaults() {
	if _, ok := rpc.mutation.CreatedAt(); !ok {
		v := registeredprovider.DefaultCreatedAt()
		rpc.mutation.SetCreatedAt(v)
	}
	if _, ok := rpc.mutation.UpdatedAt(); !ok {
		v := registeredprovider.DefaultUpdatedAt()
		rpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rpc.mutation.ID(); !ok {
		v := registeredprovider.DefaultID()
		rpc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpc *RegisteredProviderCreate) check() error {
	if _, ok := rpc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "RegisteredProvider.url"`)}
	}
	if _, ok := rpc.mutation.ContractID(); !ok {
		return &ValidationError{Name: "contract_id", err: errors.New(`ent: missing required field "RegisteredProvider.contract_id"`)}
	}
	if v, ok := rpc.mutation.ContractID(); ok {
		if err := registeredprovider.ContractIDValidator(v); err != nil {
			return &ValidationError{Name: "contract_id", err: fmt.Errorf(`ent: validator failed for field "RegisteredProvider.contract_id": %w`, err)}
		}
	}
	if _, ok := rpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "RegisteredProvider.created_at"`)}
	}
	if _, ok := rpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "RegisteredProvider.updated_at"`)}
	}
	if _, ok := rpc.mutation.ContractID(); !ok {
		return &ValidationError{Name: "contract", err: errors.New(`ent: missing required edge "RegisteredProvider.contract"`)}
	}
	return nil
}

func (rpc *RegisteredProviderCreate) sqlSave(ctx context.Context) (*RegisteredProvider, error) {
	if err := rpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rpc.mutation.id = &_node.ID
	rpc.mutation.done = true
	return _node, nil
}

func (rpc *RegisteredProviderCreate) createSpec() (*RegisteredProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &RegisteredProvider{config: rpc.config}
		_spec = sqlgraph.NewCreateSpec(registeredprovider.Table, sqlgraph.NewFieldSpec(registeredprovider.FieldID, field.TypeUUID))
	)
	if id, ok := rpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rpc.mutation.URL(); ok {
		_spec.SetField(registeredprovider.FieldURL, field.TypeJSON, value)
		_node.URL = value
	}
	if value, ok := rpc.mutation.CreatedAt(); ok {
		_spec.SetField(registeredprovider.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := rpc.mutation.UpdatedAt(); ok {
		_spec.SetField(registeredprovider.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := rpc.mutation.ContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   registeredprovider.ContractTable,
			Columns: []string{registeredprovider.ContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registeredcontract.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ContractID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RegisteredProviderCreateBulk is the builder for creating many RegisteredProvider entities in bulk.
type RegisteredProviderCreateBulk struct {
	config
	err      error
	builders []*RegisteredProviderCreate
}

// Save creates the RegisteredProvider entities in the database.
func (rpcb *RegisteredProviderCreateBulk) Save(ctx context.Context) ([]*RegisteredProvider, error) {
	if rpcb.err != nil {
		return nil, rpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rpcb.builders))
	nodes := make([]*RegisteredProvider, len(rpcb.builders))
	mutators := make([]Mutator, len(rpcb.builders))
	for i := range rpcb.builders {
		func(i int, root context.Context) {
			builder := rpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RegisteredProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, rpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rpcb *RegisteredProviderCreateBulk) SaveX(ctx context.Context) []*RegisteredProvider {
	v, err := rpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rpcb *RegisteredProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := rpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpcb *RegisteredProviderCreateBulk) ExecX(ctx context.Context) {
	if err := rpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

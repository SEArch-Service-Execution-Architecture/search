// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/clpombo/search/ent/compatibilityresult"
	"github.com/clpombo/search/ent/registeredcontract"
	"github.com/clpombo/search/ent/schema"
)

// CompatibilityResultCreate is the builder for creating a CompatibilityResult entity.
type CompatibilityResultCreate struct {
	config
	mutation *CompatibilityResultMutation
	hooks    []Hook
}

// SetRequirementContractID sets the "requirement_contract_id" field.
func (crc *CompatibilityResultCreate) SetRequirementContractID(s string) *CompatibilityResultCreate {
	crc.mutation.SetRequirementContractID(s)
	return crc
}

// SetProviderContractID sets the "provider_contract_id" field.
func (crc *CompatibilityResultCreate) SetProviderContractID(s string) *CompatibilityResultCreate {
	crc.mutation.SetProviderContractID(s)
	return crc
}

// SetParticipantNameReq sets the "participant_name_req" field.
func (crc *CompatibilityResultCreate) SetParticipantNameReq(s string) *CompatibilityResultCreate {
	crc.mutation.SetParticipantNameReq(s)
	return crc
}

// SetParticipantNameProv sets the "participant_name_prov" field.
func (crc *CompatibilityResultCreate) SetParticipantNameProv(s string) *CompatibilityResultCreate {
	crc.mutation.SetParticipantNameProv(s)
	return crc
}

// SetResult sets the "result" field.
func (crc *CompatibilityResultCreate) SetResult(b bool) *CompatibilityResultCreate {
	crc.mutation.SetResult(b)
	return crc
}

// SetMapping sets the "mapping" field.
func (crc *CompatibilityResultCreate) SetMapping(snm schema.ParticipantNameMapping) *CompatibilityResultCreate {
	crc.mutation.SetMapping(snm)
	return crc
}

// SetCreatedAt sets the "created_at" field.
func (crc *CompatibilityResultCreate) SetCreatedAt(t time.Time) *CompatibilityResultCreate {
	crc.mutation.SetCreatedAt(t)
	return crc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (crc *CompatibilityResultCreate) SetNillableCreatedAt(t *time.Time) *CompatibilityResultCreate {
	if t != nil {
		crc.SetCreatedAt(*t)
	}
	return crc
}

// SetUpdatedAt sets the "updated_at" field.
func (crc *CompatibilityResultCreate) SetUpdatedAt(t time.Time) *CompatibilityResultCreate {
	crc.mutation.SetUpdatedAt(t)
	return crc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (crc *CompatibilityResultCreate) SetNillableUpdatedAt(t *time.Time) *CompatibilityResultCreate {
	if t != nil {
		crc.SetUpdatedAt(*t)
	}
	return crc
}

// SetRequirementContract sets the "requirement_contract" edge to the RegisteredContract entity.
func (crc *CompatibilityResultCreate) SetRequirementContract(r *RegisteredContract) *CompatibilityResultCreate {
	return crc.SetRequirementContractID(r.ID)
}

// SetProviderContract sets the "provider_contract" edge to the RegisteredContract entity.
func (crc *CompatibilityResultCreate) SetProviderContract(r *RegisteredContract) *CompatibilityResultCreate {
	return crc.SetProviderContractID(r.ID)
}

// Mutation returns the CompatibilityResultMutation object of the builder.
func (crc *CompatibilityResultCreate) Mutation() *CompatibilityResultMutation {
	return crc.mutation
}

// Save creates the CompatibilityResult in the database.
func (crc *CompatibilityResultCreate) Save(ctx context.Context) (*CompatibilityResult, error) {
	crc.defaults()
	return withHooks[*CompatibilityResult, CompatibilityResultMutation](ctx, crc.sqlSave, crc.mutation, crc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (crc *CompatibilityResultCreate) SaveX(ctx context.Context) *CompatibilityResult {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crc *CompatibilityResultCreate) Exec(ctx context.Context) error {
	_, err := crc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crc *CompatibilityResultCreate) ExecX(ctx context.Context) {
	if err := crc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crc *CompatibilityResultCreate) defaults() {
	if _, ok := crc.mutation.CreatedAt(); !ok {
		v := compatibilityresult.DefaultCreatedAt()
		crc.mutation.SetCreatedAt(v)
	}
	if _, ok := crc.mutation.UpdatedAt(); !ok {
		v := compatibilityresult.DefaultUpdatedAt()
		crc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crc *CompatibilityResultCreate) check() error {
	if _, ok := crc.mutation.RequirementContractID(); !ok {
		return &ValidationError{Name: "requirement_contract_id", err: errors.New(`ent: missing required field "CompatibilityResult.requirement_contract_id"`)}
	}
	if v, ok := crc.mutation.RequirementContractID(); ok {
		if err := compatibilityresult.RequirementContractIDValidator(v); err != nil {
			return &ValidationError{Name: "requirement_contract_id", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.requirement_contract_id": %w`, err)}
		}
	}
	if _, ok := crc.mutation.ProviderContractID(); !ok {
		return &ValidationError{Name: "provider_contract_id", err: errors.New(`ent: missing required field "CompatibilityResult.provider_contract_id"`)}
	}
	if v, ok := crc.mutation.ProviderContractID(); ok {
		if err := compatibilityresult.ProviderContractIDValidator(v); err != nil {
			return &ValidationError{Name: "provider_contract_id", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.provider_contract_id": %w`, err)}
		}
	}
	if _, ok := crc.mutation.ParticipantNameReq(); !ok {
		return &ValidationError{Name: "participant_name_req", err: errors.New(`ent: missing required field "CompatibilityResult.participant_name_req"`)}
	}
	if v, ok := crc.mutation.ParticipantNameReq(); ok {
		if err := compatibilityresult.ParticipantNameReqValidator(v); err != nil {
			return &ValidationError{Name: "participant_name_req", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.participant_name_req": %w`, err)}
		}
	}
	if _, ok := crc.mutation.ParticipantNameProv(); !ok {
		return &ValidationError{Name: "participant_name_prov", err: errors.New(`ent: missing required field "CompatibilityResult.participant_name_prov"`)}
	}
	if v, ok := crc.mutation.ParticipantNameProv(); ok {
		if err := compatibilityresult.ParticipantNameProvValidator(v); err != nil {
			return &ValidationError{Name: "participant_name_prov", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.participant_name_prov": %w`, err)}
		}
	}
	if _, ok := crc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "CompatibilityResult.result"`)}
	}
	if _, ok := crc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CompatibilityResult.created_at"`)}
	}
	if _, ok := crc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CompatibilityResult.updated_at"`)}
	}
	if _, ok := crc.mutation.RequirementContractID(); !ok {
		return &ValidationError{Name: "requirement_contract", err: errors.New(`ent: missing required edge "CompatibilityResult.requirement_contract"`)}
	}
	if _, ok := crc.mutation.ProviderContractID(); !ok {
		return &ValidationError{Name: "provider_contract", err: errors.New(`ent: missing required edge "CompatibilityResult.provider_contract"`)}
	}
	return nil
}

func (crc *CompatibilityResultCreate) sqlSave(ctx context.Context) (*CompatibilityResult, error) {
	if err := crc.check(); err != nil {
		return nil, err
	}
	_node, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	crc.mutation.id = &_node.ID
	crc.mutation.done = true
	return _node, nil
}

func (crc *CompatibilityResultCreate) createSpec() (*CompatibilityResult, *sqlgraph.CreateSpec) {
	var (
		_node = &CompatibilityResult{config: crc.config}
		_spec = sqlgraph.NewCreateSpec(compatibilityresult.Table, sqlgraph.NewFieldSpec(compatibilityresult.FieldID, field.TypeInt))
	)
	if value, ok := crc.mutation.ParticipantNameReq(); ok {
		_spec.SetField(compatibilityresult.FieldParticipantNameReq, field.TypeString, value)
		_node.ParticipantNameReq = value
	}
	if value, ok := crc.mutation.ParticipantNameProv(); ok {
		_spec.SetField(compatibilityresult.FieldParticipantNameProv, field.TypeString, value)
		_node.ParticipantNameProv = value
	}
	if value, ok := crc.mutation.Result(); ok {
		_spec.SetField(compatibilityresult.FieldResult, field.TypeBool, value)
		_node.Result = value
	}
	if value, ok := crc.mutation.Mapping(); ok {
		_spec.SetField(compatibilityresult.FieldMapping, field.TypeJSON, value)
		_node.Mapping = value
	}
	if value, ok := crc.mutation.CreatedAt(); ok {
		_spec.SetField(compatibilityresult.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := crc.mutation.UpdatedAt(); ok {
		_spec.SetField(compatibilityresult.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := crc.mutation.RequirementContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   compatibilityresult.RequirementContractTable,
			Columns: []string{compatibilityresult.RequirementContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registeredcontract.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RequirementContractID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := crc.mutation.ProviderContractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   compatibilityresult.ProviderContractTable,
			Columns: []string{compatibilityresult.ProviderContractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(registeredcontract.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProviderContractID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompatibilityResultCreateBulk is the builder for creating many CompatibilityResult entities in bulk.
type CompatibilityResultCreateBulk struct {
	config
	builders []*CompatibilityResultCreate
}

// Save creates the CompatibilityResult entities in the database.
func (crcb *CompatibilityResultCreateBulk) Save(ctx context.Context) ([]*CompatibilityResult, error) {
	specs := make([]*sqlgraph.CreateSpec, len(crcb.builders))
	nodes := make([]*CompatibilityResult, len(crcb.builders))
	mutators := make([]Mutator, len(crcb.builders))
	for i := range crcb.builders {
		func(i int, root context.Context) {
			builder := crcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompatibilityResultMutation)
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
					_, err = mutators[i+1].Mutate(root, crcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, crcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crcb *CompatibilityResultCreateBulk) SaveX(ctx context.Context) []*CompatibilityResult {
	v, err := crcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crcb *CompatibilityResultCreateBulk) Exec(ctx context.Context) error {
	_, err := crcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcb *CompatibilityResultCreateBulk) ExecX(ctx context.Context) {
	if err := crcb.Exec(ctx); err != nil {
		panic(err)
	}
}

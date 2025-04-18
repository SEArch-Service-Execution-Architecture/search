// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/compatibilityresult"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/predicate"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/registeredcontract"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/schema"
)

// CompatibilityResultUpdate is the builder for updating CompatibilityResult entities.
type CompatibilityResultUpdate struct {
	config
	hooks    []Hook
	mutation *CompatibilityResultMutation
}

// Where appends a list predicates to the CompatibilityResultUpdate builder.
func (cru *CompatibilityResultUpdate) Where(ps ...predicate.CompatibilityResult) *CompatibilityResultUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetRequirementContractID sets the "requirement_contract_id" field.
func (cru *CompatibilityResultUpdate) SetRequirementContractID(s string) *CompatibilityResultUpdate {
	cru.mutation.SetRequirementContractID(s)
	return cru
}

// SetProviderContractID sets the "provider_contract_id" field.
func (cru *CompatibilityResultUpdate) SetProviderContractID(s string) *CompatibilityResultUpdate {
	cru.mutation.SetProviderContractID(s)
	return cru
}

// SetResult sets the "result" field.
func (cru *CompatibilityResultUpdate) SetResult(b bool) *CompatibilityResultUpdate {
	cru.mutation.SetResult(b)
	return cru
}

// SetMapping sets the "mapping" field.
func (cru *CompatibilityResultUpdate) SetMapping(snm schema.ParticipantNameMapping) *CompatibilityResultUpdate {
	cru.mutation.SetMapping(snm)
	return cru
}

// ClearMapping clears the value of the "mapping" field.
func (cru *CompatibilityResultUpdate) ClearMapping() *CompatibilityResultUpdate {
	cru.mutation.ClearMapping()
	return cru
}

// SetUpdatedAt sets the "updated_at" field.
func (cru *CompatibilityResultUpdate) SetUpdatedAt(t time.Time) *CompatibilityResultUpdate {
	cru.mutation.SetUpdatedAt(t)
	return cru
}

// SetRequirementContract sets the "requirement_contract" edge to the RegisteredContract entity.
func (cru *CompatibilityResultUpdate) SetRequirementContract(r *RegisteredContract) *CompatibilityResultUpdate {
	return cru.SetRequirementContractID(r.ID)
}

// SetProviderContract sets the "provider_contract" edge to the RegisteredContract entity.
func (cru *CompatibilityResultUpdate) SetProviderContract(r *RegisteredContract) *CompatibilityResultUpdate {
	return cru.SetProviderContractID(r.ID)
}

// Mutation returns the CompatibilityResultMutation object of the builder.
func (cru *CompatibilityResultUpdate) Mutation() *CompatibilityResultMutation {
	return cru.mutation
}

// ClearRequirementContract clears the "requirement_contract" edge to the RegisteredContract entity.
func (cru *CompatibilityResultUpdate) ClearRequirementContract() *CompatibilityResultUpdate {
	cru.mutation.ClearRequirementContract()
	return cru
}

// ClearProviderContract clears the "provider_contract" edge to the RegisteredContract entity.
func (cru *CompatibilityResultUpdate) ClearProviderContract() *CompatibilityResultUpdate {
	cru.mutation.ClearProviderContract()
	return cru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *CompatibilityResultUpdate) Save(ctx context.Context) (int, error) {
	cru.defaults()
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *CompatibilityResultUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *CompatibilityResultUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *CompatibilityResultUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cru *CompatibilityResultUpdate) defaults() {
	if _, ok := cru.mutation.UpdatedAt(); !ok {
		v := compatibilityresult.UpdateDefaultUpdatedAt()
		cru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *CompatibilityResultUpdate) check() error {
	if v, ok := cru.mutation.RequirementContractID(); ok {
		if err := compatibilityresult.RequirementContractIDValidator(v); err != nil {
			return &ValidationError{Name: "requirement_contract_id", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.requirement_contract_id": %w`, err)}
		}
	}
	if v, ok := cru.mutation.ProviderContractID(); ok {
		if err := compatibilityresult.ProviderContractIDValidator(v); err != nil {
			return &ValidationError{Name: "provider_contract_id", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.provider_contract_id": %w`, err)}
		}
	}
	if _, ok := cru.mutation.RequirementContractID(); cru.mutation.RequirementContractCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CompatibilityResult.requirement_contract"`)
	}
	if _, ok := cru.mutation.ProviderContractID(); cru.mutation.ProviderContractCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CompatibilityResult.provider_contract"`)
	}
	return nil
}

func (cru *CompatibilityResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(compatibilityresult.Table, compatibilityresult.Columns, sqlgraph.NewFieldSpec(compatibilityresult.FieldID, field.TypeInt))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.Result(); ok {
		_spec.SetField(compatibilityresult.FieldResult, field.TypeBool, value)
	}
	if value, ok := cru.mutation.Mapping(); ok {
		_spec.SetField(compatibilityresult.FieldMapping, field.TypeJSON, value)
	}
	if cru.mutation.MappingCleared() {
		_spec.ClearField(compatibilityresult.FieldMapping, field.TypeJSON)
	}
	if value, ok := cru.mutation.UpdatedAt(); ok {
		_spec.SetField(compatibilityresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if cru.mutation.RequirementContractCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.RequirementContractIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cru.mutation.ProviderContractCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.ProviderContractIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{compatibilityresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// CompatibilityResultUpdateOne is the builder for updating a single CompatibilityResult entity.
type CompatibilityResultUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CompatibilityResultMutation
}

// SetRequirementContractID sets the "requirement_contract_id" field.
func (cruo *CompatibilityResultUpdateOne) SetRequirementContractID(s string) *CompatibilityResultUpdateOne {
	cruo.mutation.SetRequirementContractID(s)
	return cruo
}

// SetProviderContractID sets the "provider_contract_id" field.
func (cruo *CompatibilityResultUpdateOne) SetProviderContractID(s string) *CompatibilityResultUpdateOne {
	cruo.mutation.SetProviderContractID(s)
	return cruo
}

// SetResult sets the "result" field.
func (cruo *CompatibilityResultUpdateOne) SetResult(b bool) *CompatibilityResultUpdateOne {
	cruo.mutation.SetResult(b)
	return cruo
}

// SetMapping sets the "mapping" field.
func (cruo *CompatibilityResultUpdateOne) SetMapping(snm schema.ParticipantNameMapping) *CompatibilityResultUpdateOne {
	cruo.mutation.SetMapping(snm)
	return cruo
}

// ClearMapping clears the value of the "mapping" field.
func (cruo *CompatibilityResultUpdateOne) ClearMapping() *CompatibilityResultUpdateOne {
	cruo.mutation.ClearMapping()
	return cruo
}

// SetUpdatedAt sets the "updated_at" field.
func (cruo *CompatibilityResultUpdateOne) SetUpdatedAt(t time.Time) *CompatibilityResultUpdateOne {
	cruo.mutation.SetUpdatedAt(t)
	return cruo
}

// SetRequirementContract sets the "requirement_contract" edge to the RegisteredContract entity.
func (cruo *CompatibilityResultUpdateOne) SetRequirementContract(r *RegisteredContract) *CompatibilityResultUpdateOne {
	return cruo.SetRequirementContractID(r.ID)
}

// SetProviderContract sets the "provider_contract" edge to the RegisteredContract entity.
func (cruo *CompatibilityResultUpdateOne) SetProviderContract(r *RegisteredContract) *CompatibilityResultUpdateOne {
	return cruo.SetProviderContractID(r.ID)
}

// Mutation returns the CompatibilityResultMutation object of the builder.
func (cruo *CompatibilityResultUpdateOne) Mutation() *CompatibilityResultMutation {
	return cruo.mutation
}

// ClearRequirementContract clears the "requirement_contract" edge to the RegisteredContract entity.
func (cruo *CompatibilityResultUpdateOne) ClearRequirementContract() *CompatibilityResultUpdateOne {
	cruo.mutation.ClearRequirementContract()
	return cruo
}

// ClearProviderContract clears the "provider_contract" edge to the RegisteredContract entity.
func (cruo *CompatibilityResultUpdateOne) ClearProviderContract() *CompatibilityResultUpdateOne {
	cruo.mutation.ClearProviderContract()
	return cruo
}

// Where appends a list predicates to the CompatibilityResultUpdate builder.
func (cruo *CompatibilityResultUpdateOne) Where(ps ...predicate.CompatibilityResult) *CompatibilityResultUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *CompatibilityResultUpdateOne) Select(field string, fields ...string) *CompatibilityResultUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated CompatibilityResult entity.
func (cruo *CompatibilityResultUpdateOne) Save(ctx context.Context) (*CompatibilityResult, error) {
	cruo.defaults()
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *CompatibilityResultUpdateOne) SaveX(ctx context.Context) *CompatibilityResult {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *CompatibilityResultUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *CompatibilityResultUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cruo *CompatibilityResultUpdateOne) defaults() {
	if _, ok := cruo.mutation.UpdatedAt(); !ok {
		v := compatibilityresult.UpdateDefaultUpdatedAt()
		cruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *CompatibilityResultUpdateOne) check() error {
	if v, ok := cruo.mutation.RequirementContractID(); ok {
		if err := compatibilityresult.RequirementContractIDValidator(v); err != nil {
			return &ValidationError{Name: "requirement_contract_id", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.requirement_contract_id": %w`, err)}
		}
	}
	if v, ok := cruo.mutation.ProviderContractID(); ok {
		if err := compatibilityresult.ProviderContractIDValidator(v); err != nil {
			return &ValidationError{Name: "provider_contract_id", err: fmt.Errorf(`ent: validator failed for field "CompatibilityResult.provider_contract_id": %w`, err)}
		}
	}
	if _, ok := cruo.mutation.RequirementContractID(); cruo.mutation.RequirementContractCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CompatibilityResult.requirement_contract"`)
	}
	if _, ok := cruo.mutation.ProviderContractID(); cruo.mutation.ProviderContractCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "CompatibilityResult.provider_contract"`)
	}
	return nil
}

func (cruo *CompatibilityResultUpdateOne) sqlSave(ctx context.Context) (_node *CompatibilityResult, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(compatibilityresult.Table, compatibilityresult.Columns, sqlgraph.NewFieldSpec(compatibilityresult.FieldID, field.TypeInt))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CompatibilityResult.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, compatibilityresult.FieldID)
		for _, f := range fields {
			if !compatibilityresult.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != compatibilityresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.Result(); ok {
		_spec.SetField(compatibilityresult.FieldResult, field.TypeBool, value)
	}
	if value, ok := cruo.mutation.Mapping(); ok {
		_spec.SetField(compatibilityresult.FieldMapping, field.TypeJSON, value)
	}
	if cruo.mutation.MappingCleared() {
		_spec.ClearField(compatibilityresult.FieldMapping, field.TypeJSON)
	}
	if value, ok := cruo.mutation.UpdatedAt(); ok {
		_spec.SetField(compatibilityresult.FieldUpdatedAt, field.TypeTime, value)
	}
	if cruo.mutation.RequirementContractCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.RequirementContractIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cruo.mutation.ProviderContractCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.ProviderContractIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CompatibilityResult{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{compatibilityresult.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package registeredprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/SEArch-Service-Execution-Architecture/search/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLTE(FieldID, id))
}

// ContractID applies equality check predicate on the "contract_id" field. It's identical to ContractIDEQ.
func ContractID(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldContractID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// ContractIDEQ applies the EQ predicate on the "contract_id" field.
func ContractIDEQ(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldContractID, v))
}

// ContractIDNEQ applies the NEQ predicate on the "contract_id" field.
func ContractIDNEQ(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNEQ(FieldContractID, v))
}

// ContractIDIn applies the In predicate on the "contract_id" field.
func ContractIDIn(vs ...string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldIn(FieldContractID, vs...))
}

// ContractIDNotIn applies the NotIn predicate on the "contract_id" field.
func ContractIDNotIn(vs ...string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNotIn(FieldContractID, vs...))
}

// ContractIDGT applies the GT predicate on the "contract_id" field.
func ContractIDGT(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGT(FieldContractID, v))
}

// ContractIDGTE applies the GTE predicate on the "contract_id" field.
func ContractIDGTE(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGTE(FieldContractID, v))
}

// ContractIDLT applies the LT predicate on the "contract_id" field.
func ContractIDLT(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLT(FieldContractID, v))
}

// ContractIDLTE applies the LTE predicate on the "contract_id" field.
func ContractIDLTE(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLTE(FieldContractID, v))
}

// ContractIDContains applies the Contains predicate on the "contract_id" field.
func ContractIDContains(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldContains(FieldContractID, v))
}

// ContractIDHasPrefix applies the HasPrefix predicate on the "contract_id" field.
func ContractIDHasPrefix(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldHasPrefix(FieldContractID, v))
}

// ContractIDHasSuffix applies the HasSuffix predicate on the "contract_id" field.
func ContractIDHasSuffix(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldHasSuffix(FieldContractID, v))
}

// ContractIDEqualFold applies the EqualFold predicate on the "contract_id" field.
func ContractIDEqualFold(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEqualFold(FieldContractID, v))
}

// ContractIDContainsFold applies the ContainsFold predicate on the "contract_id" field.
func ContractIDContainsFold(v string) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldContainsFold(FieldContractID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasContract applies the HasEdge predicate on the "contract" edge.
func HasContract() predicate.RegisteredProvider {
	return predicate.RegisteredProvider(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ContractTable, ContractColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContractWith applies the HasEdge predicate on the "contract" edge with a given conditions (other predicates).
func HasContractWith(preds ...predicate.RegisteredContract) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(func(s *sql.Selector) {
		step := newContractStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RegisteredProvider) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RegisteredProvider) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.RegisteredProvider) predicate.RegisteredProvider {
	return predicate.RegisteredProvider(sql.NotPredicates(p))
}

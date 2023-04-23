// Code generated by ent, DO NOT EDIT.

package registeredprovider

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the registeredprovider type in the database.
	Label = "registered_provider"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldParticipantName holds the string denoting the participant_name field in the database.
	FieldParticipantName = "participant_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeContract holds the string denoting the contract edge name in mutations.
	EdgeContract = "contract"
	// Table holds the table name of the registeredprovider in the database.
	Table = "registered_providers"
	// ContractTable is the table that holds the contract relation/edge.
	ContractTable = "registered_providers"
	// ContractInverseTable is the table name for the RegisteredContract entity.
	// It exists in this package in order to avoid circular dependency with the "registeredcontract" package.
	ContractInverseTable = "registered_contracts"
	// ContractColumn is the table column denoting the contract relation/edge.
	ContractColumn = "registered_contract_providers"
)

// Columns holds all SQL columns for registeredprovider fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldParticipantName,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "registered_providers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"registered_contract_providers",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// ParticipantNameValidator is a validator for the "participant_name" field. It is called by the builders before save.
	ParticipantNameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the RegisteredProvider queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByParticipantName orders the results by the participant_name field.
func ByParticipantName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParticipantName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByContractField orders the results by contract field.
func ByContractField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newContractStep(), sql.OrderByField(field, opts...))
	}
}
func newContractStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ContractInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ContractTable, ContractColumn),
	)
}

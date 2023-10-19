// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/pmontepagano/search/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/pmontepagano/search/ent/compatibilityresult"
	"github.com/pmontepagano/search/ent/registeredcontract"
	"github.com/pmontepagano/search/ent/registeredprovider"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CompatibilityResult is the client for interacting with the CompatibilityResult builders.
	CompatibilityResult *CompatibilityResultClient
	// RegisteredContract is the client for interacting with the RegisteredContract builders.
	RegisteredContract *RegisteredContractClient
	// RegisteredProvider is the client for interacting with the RegisteredProvider builders.
	RegisteredProvider *RegisteredProviderClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CompatibilityResult = NewCompatibilityResultClient(c.config)
	c.RegisteredContract = NewRegisteredContractClient(c.config)
	c.RegisteredProvider = NewRegisteredProviderClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		CompatibilityResult: NewCompatibilityResultClient(cfg),
		RegisteredContract:  NewRegisteredContractClient(cfg),
		RegisteredProvider:  NewRegisteredProviderClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                 ctx,
		config:              cfg,
		CompatibilityResult: NewCompatibilityResultClient(cfg),
		RegisteredContract:  NewRegisteredContractClient(cfg),
		RegisteredProvider:  NewRegisteredProviderClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CompatibilityResult.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CompatibilityResult.Use(hooks...)
	c.RegisteredContract.Use(hooks...)
	c.RegisteredProvider.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.CompatibilityResult.Intercept(interceptors...)
	c.RegisteredContract.Intercept(interceptors...)
	c.RegisteredProvider.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CompatibilityResultMutation:
		return c.CompatibilityResult.mutate(ctx, m)
	case *RegisteredContractMutation:
		return c.RegisteredContract.mutate(ctx, m)
	case *RegisteredProviderMutation:
		return c.RegisteredProvider.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CompatibilityResultClient is a client for the CompatibilityResult schema.
type CompatibilityResultClient struct {
	config
}

// NewCompatibilityResultClient returns a client for the CompatibilityResult from the given config.
func NewCompatibilityResultClient(c config) *CompatibilityResultClient {
	return &CompatibilityResultClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `compatibilityresult.Hooks(f(g(h())))`.
func (c *CompatibilityResultClient) Use(hooks ...Hook) {
	c.hooks.CompatibilityResult = append(c.hooks.CompatibilityResult, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `compatibilityresult.Intercept(f(g(h())))`.
func (c *CompatibilityResultClient) Intercept(interceptors ...Interceptor) {
	c.inters.CompatibilityResult = append(c.inters.CompatibilityResult, interceptors...)
}

// Create returns a builder for creating a CompatibilityResult entity.
func (c *CompatibilityResultClient) Create() *CompatibilityResultCreate {
	mutation := newCompatibilityResultMutation(c.config, OpCreate)
	return &CompatibilityResultCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CompatibilityResult entities.
func (c *CompatibilityResultClient) CreateBulk(builders ...*CompatibilityResultCreate) *CompatibilityResultCreateBulk {
	return &CompatibilityResultCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CompatibilityResultClient) MapCreateBulk(slice any, setFunc func(*CompatibilityResultCreate, int)) *CompatibilityResultCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CompatibilityResultCreateBulk{err: fmt.Errorf("calling to CompatibilityResultClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CompatibilityResultCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CompatibilityResultCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CompatibilityResult.
func (c *CompatibilityResultClient) Update() *CompatibilityResultUpdate {
	mutation := newCompatibilityResultMutation(c.config, OpUpdate)
	return &CompatibilityResultUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CompatibilityResultClient) UpdateOne(cr *CompatibilityResult) *CompatibilityResultUpdateOne {
	mutation := newCompatibilityResultMutation(c.config, OpUpdateOne, withCompatibilityResult(cr))
	return &CompatibilityResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CompatibilityResultClient) UpdateOneID(id int) *CompatibilityResultUpdateOne {
	mutation := newCompatibilityResultMutation(c.config, OpUpdateOne, withCompatibilityResultID(id))
	return &CompatibilityResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CompatibilityResult.
func (c *CompatibilityResultClient) Delete() *CompatibilityResultDelete {
	mutation := newCompatibilityResultMutation(c.config, OpDelete)
	return &CompatibilityResultDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CompatibilityResultClient) DeleteOne(cr *CompatibilityResult) *CompatibilityResultDeleteOne {
	return c.DeleteOneID(cr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CompatibilityResultClient) DeleteOneID(id int) *CompatibilityResultDeleteOne {
	builder := c.Delete().Where(compatibilityresult.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CompatibilityResultDeleteOne{builder}
}

// Query returns a query builder for CompatibilityResult.
func (c *CompatibilityResultClient) Query() *CompatibilityResultQuery {
	return &CompatibilityResultQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCompatibilityResult},
		inters: c.Interceptors(),
	}
}

// Get returns a CompatibilityResult entity by its id.
func (c *CompatibilityResultClient) Get(ctx context.Context, id int) (*CompatibilityResult, error) {
	return c.Query().Where(compatibilityresult.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CompatibilityResultClient) GetX(ctx context.Context, id int) *CompatibilityResult {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRequirementContract queries the requirement_contract edge of a CompatibilityResult.
func (c *CompatibilityResultClient) QueryRequirementContract(cr *CompatibilityResult) *RegisteredContractQuery {
	query := (&RegisteredContractClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(compatibilityresult.Table, compatibilityresult.FieldID, id),
			sqlgraph.To(registeredcontract.Table, registeredcontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, compatibilityresult.RequirementContractTable, compatibilityresult.RequirementContractColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProviderContract queries the provider_contract edge of a CompatibilityResult.
func (c *CompatibilityResultClient) QueryProviderContract(cr *CompatibilityResult) *RegisteredContractQuery {
	query := (&RegisteredContractClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(compatibilityresult.Table, compatibilityresult.FieldID, id),
			sqlgraph.To(registeredcontract.Table, registeredcontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, compatibilityresult.ProviderContractTable, compatibilityresult.ProviderContractColumn),
		)
		fromV = sqlgraph.Neighbors(cr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CompatibilityResultClient) Hooks() []Hook {
	return c.hooks.CompatibilityResult
}

// Interceptors returns the client interceptors.
func (c *CompatibilityResultClient) Interceptors() []Interceptor {
	return c.inters.CompatibilityResult
}

func (c *CompatibilityResultClient) mutate(ctx context.Context, m *CompatibilityResultMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CompatibilityResultCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CompatibilityResultUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CompatibilityResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CompatibilityResultDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown CompatibilityResult mutation op: %q", m.Op())
	}
}

// RegisteredContractClient is a client for the RegisteredContract schema.
type RegisteredContractClient struct {
	config
}

// NewRegisteredContractClient returns a client for the RegisteredContract from the given config.
func NewRegisteredContractClient(c config) *RegisteredContractClient {
	return &RegisteredContractClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `registeredcontract.Hooks(f(g(h())))`.
func (c *RegisteredContractClient) Use(hooks ...Hook) {
	c.hooks.RegisteredContract = append(c.hooks.RegisteredContract, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `registeredcontract.Intercept(f(g(h())))`.
func (c *RegisteredContractClient) Intercept(interceptors ...Interceptor) {
	c.inters.RegisteredContract = append(c.inters.RegisteredContract, interceptors...)
}

// Create returns a builder for creating a RegisteredContract entity.
func (c *RegisteredContractClient) Create() *RegisteredContractCreate {
	mutation := newRegisteredContractMutation(c.config, OpCreate)
	return &RegisteredContractCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RegisteredContract entities.
func (c *RegisteredContractClient) CreateBulk(builders ...*RegisteredContractCreate) *RegisteredContractCreateBulk {
	return &RegisteredContractCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RegisteredContractClient) MapCreateBulk(slice any, setFunc func(*RegisteredContractCreate, int)) *RegisteredContractCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RegisteredContractCreateBulk{err: fmt.Errorf("calling to RegisteredContractClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RegisteredContractCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RegisteredContractCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RegisteredContract.
func (c *RegisteredContractClient) Update() *RegisteredContractUpdate {
	mutation := newRegisteredContractMutation(c.config, OpUpdate)
	return &RegisteredContractUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RegisteredContractClient) UpdateOne(rc *RegisteredContract) *RegisteredContractUpdateOne {
	mutation := newRegisteredContractMutation(c.config, OpUpdateOne, withRegisteredContract(rc))
	return &RegisteredContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RegisteredContractClient) UpdateOneID(id string) *RegisteredContractUpdateOne {
	mutation := newRegisteredContractMutation(c.config, OpUpdateOne, withRegisteredContractID(id))
	return &RegisteredContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RegisteredContract.
func (c *RegisteredContractClient) Delete() *RegisteredContractDelete {
	mutation := newRegisteredContractMutation(c.config, OpDelete)
	return &RegisteredContractDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RegisteredContractClient) DeleteOne(rc *RegisteredContract) *RegisteredContractDeleteOne {
	return c.DeleteOneID(rc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RegisteredContractClient) DeleteOneID(id string) *RegisteredContractDeleteOne {
	builder := c.Delete().Where(registeredcontract.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RegisteredContractDeleteOne{builder}
}

// Query returns a query builder for RegisteredContract.
func (c *RegisteredContractClient) Query() *RegisteredContractQuery {
	return &RegisteredContractQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRegisteredContract},
		inters: c.Interceptors(),
	}
}

// Get returns a RegisteredContract entity by its id.
func (c *RegisteredContractClient) Get(ctx context.Context, id string) (*RegisteredContract, error) {
	return c.Query().Where(registeredcontract.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RegisteredContractClient) GetX(ctx context.Context, id string) *RegisteredContract {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProviders queries the providers edge of a RegisteredContract.
func (c *RegisteredContractClient) QueryProviders(rc *RegisteredContract) *RegisteredProviderQuery {
	query := (&RegisteredProviderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(registeredcontract.Table, registeredcontract.FieldID, id),
			sqlgraph.To(registeredprovider.Table, registeredprovider.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, registeredcontract.ProvidersTable, registeredcontract.ProvidersColumn),
		)
		fromV = sqlgraph.Neighbors(rc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCompatibilityResultsAsRequirement queries the compatibility_results_as_requirement edge of a RegisteredContract.
func (c *RegisteredContractClient) QueryCompatibilityResultsAsRequirement(rc *RegisteredContract) *CompatibilityResultQuery {
	query := (&CompatibilityResultClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(registeredcontract.Table, registeredcontract.FieldID, id),
			sqlgraph.To(compatibilityresult.Table, compatibilityresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, registeredcontract.CompatibilityResultsAsRequirementTable, registeredcontract.CompatibilityResultsAsRequirementColumn),
		)
		fromV = sqlgraph.Neighbors(rc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCompatibilityResultsAsProvider queries the compatibility_results_as_provider edge of a RegisteredContract.
func (c *RegisteredContractClient) QueryCompatibilityResultsAsProvider(rc *RegisteredContract) *CompatibilityResultQuery {
	query := (&CompatibilityResultClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rc.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(registeredcontract.Table, registeredcontract.FieldID, id),
			sqlgraph.To(compatibilityresult.Table, compatibilityresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, registeredcontract.CompatibilityResultsAsProviderTable, registeredcontract.CompatibilityResultsAsProviderColumn),
		)
		fromV = sqlgraph.Neighbors(rc.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RegisteredContractClient) Hooks() []Hook {
	return c.hooks.RegisteredContract
}

// Interceptors returns the client interceptors.
func (c *RegisteredContractClient) Interceptors() []Interceptor {
	return c.inters.RegisteredContract
}

func (c *RegisteredContractClient) mutate(ctx context.Context, m *RegisteredContractMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RegisteredContractCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RegisteredContractUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RegisteredContractUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RegisteredContractDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RegisteredContract mutation op: %q", m.Op())
	}
}

// RegisteredProviderClient is a client for the RegisteredProvider schema.
type RegisteredProviderClient struct {
	config
}

// NewRegisteredProviderClient returns a client for the RegisteredProvider from the given config.
func NewRegisteredProviderClient(c config) *RegisteredProviderClient {
	return &RegisteredProviderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `registeredprovider.Hooks(f(g(h())))`.
func (c *RegisteredProviderClient) Use(hooks ...Hook) {
	c.hooks.RegisteredProvider = append(c.hooks.RegisteredProvider, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `registeredprovider.Intercept(f(g(h())))`.
func (c *RegisteredProviderClient) Intercept(interceptors ...Interceptor) {
	c.inters.RegisteredProvider = append(c.inters.RegisteredProvider, interceptors...)
}

// Create returns a builder for creating a RegisteredProvider entity.
func (c *RegisteredProviderClient) Create() *RegisteredProviderCreate {
	mutation := newRegisteredProviderMutation(c.config, OpCreate)
	return &RegisteredProviderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RegisteredProvider entities.
func (c *RegisteredProviderClient) CreateBulk(builders ...*RegisteredProviderCreate) *RegisteredProviderCreateBulk {
	return &RegisteredProviderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RegisteredProviderClient) MapCreateBulk(slice any, setFunc func(*RegisteredProviderCreate, int)) *RegisteredProviderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RegisteredProviderCreateBulk{err: fmt.Errorf("calling to RegisteredProviderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RegisteredProviderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RegisteredProviderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RegisteredProvider.
func (c *RegisteredProviderClient) Update() *RegisteredProviderUpdate {
	mutation := newRegisteredProviderMutation(c.config, OpUpdate)
	return &RegisteredProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RegisteredProviderClient) UpdateOne(rp *RegisteredProvider) *RegisteredProviderUpdateOne {
	mutation := newRegisteredProviderMutation(c.config, OpUpdateOne, withRegisteredProvider(rp))
	return &RegisteredProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RegisteredProviderClient) UpdateOneID(id uuid.UUID) *RegisteredProviderUpdateOne {
	mutation := newRegisteredProviderMutation(c.config, OpUpdateOne, withRegisteredProviderID(id))
	return &RegisteredProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RegisteredProvider.
func (c *RegisteredProviderClient) Delete() *RegisteredProviderDelete {
	mutation := newRegisteredProviderMutation(c.config, OpDelete)
	return &RegisteredProviderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RegisteredProviderClient) DeleteOne(rp *RegisteredProvider) *RegisteredProviderDeleteOne {
	return c.DeleteOneID(rp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RegisteredProviderClient) DeleteOneID(id uuid.UUID) *RegisteredProviderDeleteOne {
	builder := c.Delete().Where(registeredprovider.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RegisteredProviderDeleteOne{builder}
}

// Query returns a query builder for RegisteredProvider.
func (c *RegisteredProviderClient) Query() *RegisteredProviderQuery {
	return &RegisteredProviderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRegisteredProvider},
		inters: c.Interceptors(),
	}
}

// Get returns a RegisteredProvider entity by its id.
func (c *RegisteredProviderClient) Get(ctx context.Context, id uuid.UUID) (*RegisteredProvider, error) {
	return c.Query().Where(registeredprovider.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RegisteredProviderClient) GetX(ctx context.Context, id uuid.UUID) *RegisteredProvider {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryContract queries the contract edge of a RegisteredProvider.
func (c *RegisteredProviderClient) QueryContract(rp *RegisteredProvider) *RegisteredContractQuery {
	query := (&RegisteredContractClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(registeredprovider.Table, registeredprovider.FieldID, id),
			sqlgraph.To(registeredcontract.Table, registeredcontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, registeredprovider.ContractTable, registeredprovider.ContractColumn),
		)
		fromV = sqlgraph.Neighbors(rp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RegisteredProviderClient) Hooks() []Hook {
	return c.hooks.RegisteredProvider
}

// Interceptors returns the client interceptors.
func (c *RegisteredProviderClient) Interceptors() []Interceptor {
	return c.inters.RegisteredProvider
}

func (c *RegisteredProviderClient) mutate(ctx context.Context, m *RegisteredProviderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RegisteredProviderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RegisteredProviderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RegisteredProviderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RegisteredProviderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RegisteredProvider mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		CompatibilityResult, RegisteredContract, RegisteredProvider []ent.Hook
	}
	inters struct {
		CompatibilityResult, RegisteredContract, RegisteredProvider []ent.Interceptor
	}
)

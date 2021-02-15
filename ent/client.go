// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/vicanso/forest/ent/migrate"

	"github.com/vicanso/forest/ent/configuration"
	"github.com/vicanso/forest/ent/user"
	"github.com/vicanso/forest/ent/userlogin"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Configuration is the client for interacting with the Configuration builders.
	Configuration *ConfigurationClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserLogin is the client for interacting with the UserLogin builders.
	UserLogin *UserLoginClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Configuration = NewConfigurationClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserLogin = NewUserLoginClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Configuration: NewConfigurationClient(cfg),
		User:          NewUserClient(cfg),
		UserLogin:     NewUserLoginClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:        cfg,
		Configuration: NewConfigurationClient(cfg),
		User:          NewUserClient(cfg),
		UserLogin:     NewUserLoginClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Configuration.
//		Query().
//		Count(ctx)
//
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
	c.Configuration.Use(hooks...)
	c.User.Use(hooks...)
	c.UserLogin.Use(hooks...)
}

// ConfigurationClient is a client for the Configuration schema.
type ConfigurationClient struct {
	config
}

// NewConfigurationClient returns a client for the Configuration from the given config.
func NewConfigurationClient(c config) *ConfigurationClient {
	return &ConfigurationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `configuration.Hooks(f(g(h())))`.
func (c *ConfigurationClient) Use(hooks ...Hook) {
	c.hooks.Configuration = append(c.hooks.Configuration, hooks...)
}

// Create returns a create builder for Configuration.
func (c *ConfigurationClient) Create() *ConfigurationCreate {
	mutation := newConfigurationMutation(c.config, OpCreate)
	return &ConfigurationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Configuration entities.
func (c *ConfigurationClient) CreateBulk(builders ...*ConfigurationCreate) *ConfigurationCreateBulk {
	return &ConfigurationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Configuration.
func (c *ConfigurationClient) Update() *ConfigurationUpdate {
	mutation := newConfigurationMutation(c.config, OpUpdate)
	return &ConfigurationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConfigurationClient) UpdateOne(co *Configuration) *ConfigurationUpdateOne {
	mutation := newConfigurationMutation(c.config, OpUpdateOne, withConfiguration(co))
	return &ConfigurationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConfigurationClient) UpdateOneID(id int) *ConfigurationUpdateOne {
	mutation := newConfigurationMutation(c.config, OpUpdateOne, withConfigurationID(id))
	return &ConfigurationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Configuration.
func (c *ConfigurationClient) Delete() *ConfigurationDelete {
	mutation := newConfigurationMutation(c.config, OpDelete)
	return &ConfigurationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ConfigurationClient) DeleteOne(co *Configuration) *ConfigurationDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ConfigurationClient) DeleteOneID(id int) *ConfigurationDeleteOne {
	builder := c.Delete().Where(configuration.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConfigurationDeleteOne{builder}
}

// Query returns a query builder for Configuration.
func (c *ConfigurationClient) Query() *ConfigurationQuery {
	return &ConfigurationQuery{config: c.config}
}

// Get returns a Configuration entity by its id.
func (c *ConfigurationClient) Get(ctx context.Context, id int) (*Configuration, error) {
	return c.Query().Where(configuration.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConfigurationClient) GetX(ctx context.Context, id int) *Configuration {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ConfigurationClient) Hooks() []Hook {
	return c.hooks.Configuration
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// UserLoginClient is a client for the UserLogin schema.
type UserLoginClient struct {
	config
}

// NewUserLoginClient returns a client for the UserLogin from the given config.
func NewUserLoginClient(c config) *UserLoginClient {
	return &UserLoginClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userlogin.Hooks(f(g(h())))`.
func (c *UserLoginClient) Use(hooks ...Hook) {
	c.hooks.UserLogin = append(c.hooks.UserLogin, hooks...)
}

// Create returns a create builder for UserLogin.
func (c *UserLoginClient) Create() *UserLoginCreate {
	mutation := newUserLoginMutation(c.config, OpCreate)
	return &UserLoginCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserLogin entities.
func (c *UserLoginClient) CreateBulk(builders ...*UserLoginCreate) *UserLoginCreateBulk {
	return &UserLoginCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserLogin.
func (c *UserLoginClient) Update() *UserLoginUpdate {
	mutation := newUserLoginMutation(c.config, OpUpdate)
	return &UserLoginUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserLoginClient) UpdateOne(ul *UserLogin) *UserLoginUpdateOne {
	mutation := newUserLoginMutation(c.config, OpUpdateOne, withUserLogin(ul))
	return &UserLoginUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserLoginClient) UpdateOneID(id int) *UserLoginUpdateOne {
	mutation := newUserLoginMutation(c.config, OpUpdateOne, withUserLoginID(id))
	return &UserLoginUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserLogin.
func (c *UserLoginClient) Delete() *UserLoginDelete {
	mutation := newUserLoginMutation(c.config, OpDelete)
	return &UserLoginDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserLoginClient) DeleteOne(ul *UserLogin) *UserLoginDeleteOne {
	return c.DeleteOneID(ul.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserLoginClient) DeleteOneID(id int) *UserLoginDeleteOne {
	builder := c.Delete().Where(userlogin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserLoginDeleteOne{builder}
}

// Query returns a query builder for UserLogin.
func (c *UserLoginClient) Query() *UserLoginQuery {
	return &UserLoginQuery{config: c.config}
}

// Get returns a UserLogin entity by its id.
func (c *UserLoginClient) Get(ctx context.Context, id int) (*UserLogin, error) {
	return c.Query().Where(userlogin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserLoginClient) GetX(ctx context.Context, id int) *UserLogin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserLoginClient) Hooks() []Hook {
	return c.hooks.UserLogin
}

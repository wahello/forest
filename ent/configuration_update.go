// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/forest/ent/configuration"
	"github.com/vicanso/forest/ent/predicate"
	"github.com/vicanso/forest/ent/schema"
)

// ConfigurationUpdate is the builder for updating Configuration entities.
type ConfigurationUpdate struct {
	config
	hooks    []Hook
	mutation *ConfigurationMutation
}

// Where adds a new predicate for the ConfigurationUpdate builder.
func (cu *ConfigurationUpdate) Where(ps ...predicate.Configuration) *ConfigurationUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ConfigurationUpdate) SetStatus(s schema.Status) *ConfigurationUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(s)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ConfigurationUpdate) SetNillableStatus(s *schema.Status) *ConfigurationUpdate {
	if s != nil {
		cu.SetStatus(*s)
	}
	return cu
}

// AddStatus adds s to the "status" field.
func (cu *ConfigurationUpdate) AddStatus(s schema.Status) *ConfigurationUpdate {
	cu.mutation.AddStatus(s)
	return cu
}

// SetName sets the "name" field.
func (cu *ConfigurationUpdate) SetName(s string) *ConfigurationUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetCategory sets the "category" field.
func (cu *ConfigurationUpdate) SetCategory(c configuration.Category) *ConfigurationUpdate {
	cu.mutation.SetCategory(c)
	return cu
}

// SetOwner sets the "owner" field.
func (cu *ConfigurationUpdate) SetOwner(s string) *ConfigurationUpdate {
	cu.mutation.SetOwner(s)
	return cu
}

// SetData sets the "data" field.
func (cu *ConfigurationUpdate) SetData(s string) *ConfigurationUpdate {
	cu.mutation.SetData(s)
	return cu
}

// SetStartedAt sets the "started_at" field.
func (cu *ConfigurationUpdate) SetStartedAt(t time.Time) *ConfigurationUpdate {
	cu.mutation.SetStartedAt(t)
	return cu
}

// SetEndedAt sets the "ended_at" field.
func (cu *ConfigurationUpdate) SetEndedAt(t time.Time) *ConfigurationUpdate {
	cu.mutation.SetEndedAt(t)
	return cu
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cu *ConfigurationUpdate) Mutation() *ConfigurationMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfigurationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfigurationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfigurationUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfigurationUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfigurationUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConfigurationUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := configuration.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConfigurationUpdate) check() error {
	if v, ok := cu.mutation.Status(); ok {
		if err := configuration.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := configuration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Category(); ok {
		if err := configuration.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf("ent: validator failed for field \"category\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Owner(); ok {
		if err := configuration.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Data(); ok {
		if err := configuration.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf("ent: validator failed for field \"data\": %w", err)}
		}
	}
	return nil
}

func (cu *ConfigurationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   configuration.Table,
			Columns: configuration.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configuration.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configuration.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: configuration.FieldStatus,
		})
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: configuration.FieldStatus,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configuration.FieldName,
		})
	}
	if value, ok := cu.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: configuration.FieldCategory,
		})
	}
	if value, ok := cu.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configuration.FieldOwner,
		})
	}
	if value, ok := cu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configuration.FieldData,
		})
	}
	if value, ok := cu.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configuration.FieldStartedAt,
		})
	}
	if value, ok := cu.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configuration.FieldEndedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ConfigurationUpdateOne is the builder for updating a single Configuration entity.
type ConfigurationUpdateOne struct {
	config
	hooks    []Hook
	mutation *ConfigurationMutation
}

// SetStatus sets the "status" field.
func (cuo *ConfigurationUpdateOne) SetStatus(s schema.Status) *ConfigurationUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(s)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ConfigurationUpdateOne) SetNillableStatus(s *schema.Status) *ConfigurationUpdateOne {
	if s != nil {
		cuo.SetStatus(*s)
	}
	return cuo
}

// AddStatus adds s to the "status" field.
func (cuo *ConfigurationUpdateOne) AddStatus(s schema.Status) *ConfigurationUpdateOne {
	cuo.mutation.AddStatus(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ConfigurationUpdateOne) SetName(s string) *ConfigurationUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetCategory sets the "category" field.
func (cuo *ConfigurationUpdateOne) SetCategory(c configuration.Category) *ConfigurationUpdateOne {
	cuo.mutation.SetCategory(c)
	return cuo
}

// SetOwner sets the "owner" field.
func (cuo *ConfigurationUpdateOne) SetOwner(s string) *ConfigurationUpdateOne {
	cuo.mutation.SetOwner(s)
	return cuo
}

// SetData sets the "data" field.
func (cuo *ConfigurationUpdateOne) SetData(s string) *ConfigurationUpdateOne {
	cuo.mutation.SetData(s)
	return cuo
}

// SetStartedAt sets the "started_at" field.
func (cuo *ConfigurationUpdateOne) SetStartedAt(t time.Time) *ConfigurationUpdateOne {
	cuo.mutation.SetStartedAt(t)
	return cuo
}

// SetEndedAt sets the "ended_at" field.
func (cuo *ConfigurationUpdateOne) SetEndedAt(t time.Time) *ConfigurationUpdateOne {
	cuo.mutation.SetEndedAt(t)
	return cuo
}

// Mutation returns the ConfigurationMutation object of the builder.
func (cuo *ConfigurationUpdateOne) Mutation() *ConfigurationMutation {
	return cuo.mutation
}

// Save executes the query and returns the updated Configuration entity.
func (cuo *ConfigurationUpdateOne) Save(ctx context.Context) (*Configuration, error) {
	var (
		err  error
		node *Configuration
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfigurationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) SaveX(ctx context.Context) *Configuration {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfigurationUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfigurationUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConfigurationUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := configuration.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConfigurationUpdateOne) check() error {
	if v, ok := cuo.mutation.Status(); ok {
		if err := configuration.StatusValidator(int8(v)); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := configuration.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Category(); ok {
		if err := configuration.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf("ent: validator failed for field \"category\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Owner(); ok {
		if err := configuration.OwnerValidator(v); err != nil {
			return &ValidationError{Name: "owner", err: fmt.Errorf("ent: validator failed for field \"owner\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Data(); ok {
		if err := configuration.DataValidator(v); err != nil {
			return &ValidationError{Name: "data", err: fmt.Errorf("ent: validator failed for field \"data\": %w", err)}
		}
	}
	return nil
}

func (cuo *ConfigurationUpdateOne) sqlSave(ctx context.Context) (_node *Configuration, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   configuration.Table,
			Columns: configuration.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configuration.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Configuration.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configuration.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: configuration.FieldStatus,
		})
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: configuration.FieldStatus,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configuration.FieldName,
		})
	}
	if value, ok := cuo.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: configuration.FieldCategory,
		})
	}
	if value, ok := cuo.mutation.Owner(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configuration.FieldOwner,
		})
	}
	if value, ok := cuo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configuration.FieldData,
		})
	}
	if value, ok := cuo.mutation.StartedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configuration.FieldStartedAt,
		})
	}
	if value, ok := cuo.mutation.EndedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configuration.FieldEndedAt,
		})
	}
	_node = &Configuration{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configuration.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

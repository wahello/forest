// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/forest/ent/userlogin"
)

// UserLoginCreate is the builder for creating a UserLogin entity.
type UserLoginCreate struct {
	config
	mutation *UserLoginMutation
	hooks    []Hook
}

// SetCreatedAt sets the created_at field.
func (ulc *UserLoginCreate) SetCreatedAt(t time.Time) *UserLoginCreate {
	ulc.mutation.SetCreatedAt(t)
	return ulc
}

// SetNillableCreatedAt sets the created_at field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableCreatedAt(t *time.Time) *UserLoginCreate {
	if t != nil {
		ulc.SetCreatedAt(*t)
	}
	return ulc
}

// SetUpdatedAt sets the updated_at field.
func (ulc *UserLoginCreate) SetUpdatedAt(t time.Time) *UserLoginCreate {
	ulc.mutation.SetUpdatedAt(t)
	return ulc
}

// SetNillableUpdatedAt sets the updated_at field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableUpdatedAt(t *time.Time) *UserLoginCreate {
	if t != nil {
		ulc.SetUpdatedAt(*t)
	}
	return ulc
}

// SetAccount sets the account field.
func (ulc *UserLoginCreate) SetAccount(s string) *UserLoginCreate {
	ulc.mutation.SetAccount(s)
	return ulc
}

// SetUserAgent sets the user_agent field.
func (ulc *UserLoginCreate) SetUserAgent(s string) *UserLoginCreate {
	ulc.mutation.SetUserAgent(s)
	return ulc
}

// SetNillableUserAgent sets the user_agent field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableUserAgent(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetUserAgent(*s)
	}
	return ulc
}

// SetIP sets the ip field.
func (ulc *UserLoginCreate) SetIP(s string) *UserLoginCreate {
	ulc.mutation.SetIP(s)
	return ulc
}

// SetNillableIP sets the ip field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableIP(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetIP(*s)
	}
	return ulc
}

// SetTrackID sets the track_id field.
func (ulc *UserLoginCreate) SetTrackID(s string) *UserLoginCreate {
	ulc.mutation.SetTrackID(s)
	return ulc
}

// SetNillableTrackID sets the track_id field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableTrackID(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetTrackID(*s)
	}
	return ulc
}

// SetSessionID sets the session_id field.
func (ulc *UserLoginCreate) SetSessionID(s string) *UserLoginCreate {
	ulc.mutation.SetSessionID(s)
	return ulc
}

// SetNillableSessionID sets the session_id field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableSessionID(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetSessionID(*s)
	}
	return ulc
}

// SetXForwardedFor sets the x_forwarded_for field.
func (ulc *UserLoginCreate) SetXForwardedFor(s string) *UserLoginCreate {
	ulc.mutation.SetXForwardedFor(s)
	return ulc
}

// SetNillableXForwardedFor sets the x_forwarded_for field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableXForwardedFor(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetXForwardedFor(*s)
	}
	return ulc
}

// SetCountry sets the country field.
func (ulc *UserLoginCreate) SetCountry(s string) *UserLoginCreate {
	ulc.mutation.SetCountry(s)
	return ulc
}

// SetNillableCountry sets the country field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableCountry(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetCountry(*s)
	}
	return ulc
}

// SetProvince sets the province field.
func (ulc *UserLoginCreate) SetProvince(s string) *UserLoginCreate {
	ulc.mutation.SetProvince(s)
	return ulc
}

// SetNillableProvince sets the province field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableProvince(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetProvince(*s)
	}
	return ulc
}

// SetCity sets the city field.
func (ulc *UserLoginCreate) SetCity(s string) *UserLoginCreate {
	ulc.mutation.SetCity(s)
	return ulc
}

// SetNillableCity sets the city field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableCity(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetCity(*s)
	}
	return ulc
}

// SetIsp sets the isp field.
func (ulc *UserLoginCreate) SetIsp(s string) *UserLoginCreate {
	ulc.mutation.SetIsp(s)
	return ulc
}

// SetNillableIsp sets the isp field if the given value is not nil.
func (ulc *UserLoginCreate) SetNillableIsp(s *string) *UserLoginCreate {
	if s != nil {
		ulc.SetIsp(*s)
	}
	return ulc
}

// Mutation returns the UserLoginMutation object of the builder.
func (ulc *UserLoginCreate) Mutation() *UserLoginMutation {
	return ulc.mutation
}

// Save creates the UserLogin in the database.
func (ulc *UserLoginCreate) Save(ctx context.Context) (*UserLogin, error) {
	var (
		err  error
		node *UserLogin
	)
	ulc.defaults()
	if len(ulc.hooks) == 0 {
		if err = ulc.check(); err != nil {
			return nil, err
		}
		node, err = ulc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserLoginMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ulc.check(); err != nil {
				return nil, err
			}
			ulc.mutation = mutation
			node, err = ulc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ulc.hooks) - 1; i >= 0; i-- {
			mut = ulc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ulc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ulc *UserLoginCreate) SaveX(ctx context.Context) *UserLogin {
	v, err := ulc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ulc *UserLoginCreate) defaults() {
	if _, ok := ulc.mutation.CreatedAt(); !ok {
		v := userlogin.DefaultCreatedAt()
		ulc.mutation.SetCreatedAt(v)
	}
	if _, ok := ulc.mutation.UpdatedAt(); !ok {
		v := userlogin.DefaultUpdatedAt()
		ulc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ulc *UserLoginCreate) check() error {
	if _, ok := ulc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := ulc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := ulc.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New("ent: missing required field \"account\"")}
	}
	if v, ok := ulc.mutation.Account(); ok {
		if err := userlogin.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf("ent: validator failed for field \"account\": %w", err)}
		}
	}
	return nil
}

func (ulc *UserLoginCreate) sqlSave(ctx context.Context) (*UserLogin, error) {
	_node, _spec := ulc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ulc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ulc *UserLoginCreate) createSpec() (*UserLogin, *sqlgraph.CreateSpec) {
	var (
		_node = &UserLogin{config: ulc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userlogin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userlogin.FieldID,
			},
		}
	)
	if value, ok := ulc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlogin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ulc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userlogin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ulc.mutation.Account(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldAccount,
		})
		_node.Account = value
	}
	if value, ok := ulc.mutation.UserAgent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldUserAgent,
		})
		_node.UserAgent = value
	}
	if value, ok := ulc.mutation.IP(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldIP,
		})
		_node.IP = value
	}
	if value, ok := ulc.mutation.TrackID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldTrackID,
		})
		_node.TrackID = value
	}
	if value, ok := ulc.mutation.SessionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldSessionID,
		})
		_node.SessionID = value
	}
	if value, ok := ulc.mutation.XForwardedFor(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldXForwardedFor,
		})
		_node.XForwardedFor = value
	}
	if value, ok := ulc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := ulc.mutation.Province(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldProvince,
		})
		_node.Province = value
	}
	if value, ok := ulc.mutation.City(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldCity,
		})
		_node.City = value
	}
	if value, ok := ulc.mutation.Isp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userlogin.FieldIsp,
		})
		_node.Isp = value
	}
	return _node, _spec
}

// UserLoginCreateBulk is the builder for creating a bulk of UserLogin entities.
type UserLoginCreateBulk struct {
	config
	builders []*UserLoginCreate
}

// Save creates the UserLogin entities in the database.
func (ulcb *UserLoginCreateBulk) Save(ctx context.Context) ([]*UserLogin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ulcb.builders))
	nodes := make([]*UserLogin, len(ulcb.builders))
	mutators := make([]Mutator, len(ulcb.builders))
	for i := range ulcb.builders {
		func(i int, root context.Context) {
			builder := ulcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserLoginMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ulcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ulcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ulcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ulcb *UserLoginCreateBulk) SaveX(ctx context.Context) []*UserLogin {
	v, err := ulcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

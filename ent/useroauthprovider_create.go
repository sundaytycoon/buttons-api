// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/sundaytycoon/buttons-api/ent/user"
	"github.com/sundaytycoon/buttons-api/ent/useroauthprovider"
)

// UserOAuthProviderCreate is the builder for creating a UserOAuthProvider entity.
type UserOAuthProviderCreate struct {
	config
	mutation *UserOAuthProviderMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uopc *UserOAuthProviderCreate) SetUserID(s string) *UserOAuthProviderCreate {
	uopc.mutation.SetUserID(s)
	return uopc
}

// SetCreatedAt sets the "created_at" field.
func (uopc *UserOAuthProviderCreate) SetCreatedAt(t time.Time) *UserOAuthProviderCreate {
	uopc.mutation.SetCreatedAt(t)
	return uopc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uopc *UserOAuthProviderCreate) SetNillableCreatedAt(t *time.Time) *UserOAuthProviderCreate {
	if t != nil {
		uopc.SetCreatedAt(*t)
	}
	return uopc
}

// SetCreatedBy sets the "created_by" field.
func (uopc *UserOAuthProviderCreate) SetCreatedBy(s string) *UserOAuthProviderCreate {
	uopc.mutation.SetCreatedBy(s)
	return uopc
}

// SetUpdatedAt sets the "updated_at" field.
func (uopc *UserOAuthProviderCreate) SetUpdatedAt(t time.Time) *UserOAuthProviderCreate {
	uopc.mutation.SetUpdatedAt(t)
	return uopc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uopc *UserOAuthProviderCreate) SetNillableUpdatedAt(t *time.Time) *UserOAuthProviderCreate {
	if t != nil {
		uopc.SetUpdatedAt(*t)
	}
	return uopc
}

// SetUpdatedBy sets the "updated_by" field.
func (uopc *UserOAuthProviderCreate) SetUpdatedBy(s string) *UserOAuthProviderCreate {
	uopc.mutation.SetUpdatedBy(s)
	return uopc
}

// SetStatus sets the "status" field.
func (uopc *UserOAuthProviderCreate) SetStatus(u useroauthprovider.Status) *UserOAuthProviderCreate {
	uopc.mutation.SetStatus(u)
	return uopc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (uopc *UserOAuthProviderCreate) SetNillableStatus(u *useroauthprovider.Status) *UserOAuthProviderCreate {
	if u != nil {
		uopc.SetStatus(*u)
	}
	return uopc
}

// SetProvider sets the "provider" field.
func (uopc *UserOAuthProviderCreate) SetProvider(u useroauthprovider.Provider) *UserOAuthProviderCreate {
	uopc.mutation.SetProvider(u)
	return uopc
}

// SetExpiry sets the "expiry" field.
func (uopc *UserOAuthProviderCreate) SetExpiry(t time.Time) *UserOAuthProviderCreate {
	uopc.mutation.SetExpiry(t)
	return uopc
}

// SetAccessToken sets the "access_token" field.
func (uopc *UserOAuthProviderCreate) SetAccessToken(s string) *UserOAuthProviderCreate {
	uopc.mutation.SetAccessToken(s)
	return uopc
}

// SetRefreshToken sets the "refresh_token" field.
func (uopc *UserOAuthProviderCreate) SetRefreshToken(s string) *UserOAuthProviderCreate {
	uopc.mutation.SetRefreshToken(s)
	return uopc
}

// SetID sets the "id" field.
func (uopc *UserOAuthProviderCreate) SetID(s string) *UserOAuthProviderCreate {
	uopc.mutation.SetID(s)
	return uopc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uopc *UserOAuthProviderCreate) SetUserID(id string) *UserOAuthProviderCreate {
	uopc.mutation.SetUserID(id)
	return uopc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (uopc *UserOAuthProviderCreate) SetNillableUserID(id *string) *UserOAuthProviderCreate {
	if id != nil {
		uopc = uopc.SetUserID(*id)
	}
	return uopc
}

// SetUser sets the "user" edge to the User entity.
func (uopc *UserOAuthProviderCreate) SetUser(u *User) *UserOAuthProviderCreate {
	return uopc.SetUserID(u.ID)
}

// Mutation returns the UserOAuthProviderMutation object of the builder.
func (uopc *UserOAuthProviderCreate) Mutation() *UserOAuthProviderMutation {
	return uopc.mutation
}

// Save creates the UserOAuthProvider in the database.
func (uopc *UserOAuthProviderCreate) Save(ctx context.Context) (*UserOAuthProvider, error) {
	var (
		err  error
		node *UserOAuthProvider
	)
	uopc.defaults()
	if len(uopc.hooks) == 0 {
		if err = uopc.check(); err != nil {
			return nil, err
		}
		node, err = uopc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserOAuthProviderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uopc.check(); err != nil {
				return nil, err
			}
			uopc.mutation = mutation
			if node, err = uopc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uopc.hooks) - 1; i >= 0; i-- {
			if uopc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uopc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uopc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uopc *UserOAuthProviderCreate) SaveX(ctx context.Context) *UserOAuthProvider {
	v, err := uopc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uopc *UserOAuthProviderCreate) Exec(ctx context.Context) error {
	_, err := uopc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uopc *UserOAuthProviderCreate) ExecX(ctx context.Context) {
	if err := uopc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uopc *UserOAuthProviderCreate) defaults() {
	if _, ok := uopc.mutation.CreatedAt(); !ok {
		v := useroauthprovider.DefaultCreatedAt()
		uopc.mutation.SetCreatedAt(v)
	}
	if _, ok := uopc.mutation.UpdatedAt(); !ok {
		v := useroauthprovider.DefaultUpdatedAt()
		uopc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uopc.mutation.Status(); !ok {
		v := useroauthprovider.DefaultStatus
		uopc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uopc *UserOAuthProviderCreate) check() error {
	if _, ok := uopc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserOAuthProvider.user_id"`)}
	}
	if _, ok := uopc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserOAuthProvider.created_at"`)}
	}
	if _, ok := uopc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "UserOAuthProvider.created_by"`)}
	}
	if _, ok := uopc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserOAuthProvider.updated_at"`)}
	}
	if _, ok := uopc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "UserOAuthProvider.updated_by"`)}
	}
	if _, ok := uopc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "UserOAuthProvider.status"`)}
	}
	if v, ok := uopc.mutation.Status(); ok {
		if err := useroauthprovider.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "UserOAuthProvider.status": %w`, err)}
		}
	}
	if _, ok := uopc.mutation.Provider(); !ok {
		return &ValidationError{Name: "provider", err: errors.New(`ent: missing required field "UserOAuthProvider.provider"`)}
	}
	if v, ok := uopc.mutation.Provider(); ok {
		if err := useroauthprovider.ProviderValidator(v); err != nil {
			return &ValidationError{Name: "provider", err: fmt.Errorf(`ent: validator failed for field "UserOAuthProvider.provider": %w`, err)}
		}
	}
	if _, ok := uopc.mutation.Expiry(); !ok {
		return &ValidationError{Name: "expiry", err: errors.New(`ent: missing required field "UserOAuthProvider.expiry"`)}
	}
	if _, ok := uopc.mutation.AccessToken(); !ok {
		return &ValidationError{Name: "access_token", err: errors.New(`ent: missing required field "UserOAuthProvider.access_token"`)}
	}
	if _, ok := uopc.mutation.RefreshToken(); !ok {
		return &ValidationError{Name: "refresh_token", err: errors.New(`ent: missing required field "UserOAuthProvider.refresh_token"`)}
	}
	return nil
}

func (uopc *UserOAuthProviderCreate) sqlSave(ctx context.Context) (*UserOAuthProvider, error) {
	_node, _spec := uopc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uopc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected UserOAuthProvider.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (uopc *UserOAuthProviderCreate) createSpec() (*UserOAuthProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &UserOAuthProvider{config: uopc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: useroauthprovider.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: useroauthprovider.FieldID,
			},
		}
	)
	if id, ok := uopc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uopc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := uopc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := uopc.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := uopc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := uopc.mutation.UpdatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldUpdatedBy,
		})
		_node.UpdatedBy = value
	}
	if value, ok := uopc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: useroauthprovider.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := uopc.mutation.Provider(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: useroauthprovider.FieldProvider,
		})
		_node.Provider = value
	}
	if value, ok := uopc.mutation.Expiry(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: useroauthprovider.FieldExpiry,
		})
		_node.Expiry = value
	}
	if value, ok := uopc.mutation.AccessToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldAccessToken,
		})
		_node.AccessToken = value
	}
	if value, ok := uopc.mutation.RefreshToken(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: useroauthprovider.FieldRefreshToken,
		})
		_node.RefreshToken = value
	}
	if nodes := uopc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   useroauthprovider.UserTable,
			Columns: []string{useroauthprovider.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_oauth_providers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserOAuthProviderCreateBulk is the builder for creating many UserOAuthProvider entities in bulk.
type UserOAuthProviderCreateBulk struct {
	config
	builders []*UserOAuthProviderCreate
}

// Save creates the UserOAuthProvider entities in the database.
func (uopcb *UserOAuthProviderCreateBulk) Save(ctx context.Context) ([]*UserOAuthProvider, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uopcb.builders))
	nodes := make([]*UserOAuthProvider, len(uopcb.builders))
	mutators := make([]Mutator, len(uopcb.builders))
	for i := range uopcb.builders {
		func(i int, root context.Context) {
			builder := uopcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserOAuthProviderMutation)
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
					_, err = mutators[i+1].Mutate(root, uopcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uopcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, uopcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uopcb *UserOAuthProviderCreateBulk) SaveX(ctx context.Context) []*UserOAuthProvider {
	v, err := uopcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uopcb *UserOAuthProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := uopcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uopcb *UserOAuthProviderCreateBulk) ExecX(ctx context.Context) {
	if err := uopcb.Exec(ctx); err != nil {
		panic(err)
	}
}

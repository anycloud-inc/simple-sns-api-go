// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"simple_sns_api/ent/message"
	"simple_sns_api/ent/room"
	"simple_sns_api/ent/roomuser"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RoomCreate is the builder for creating a Room entity.
type RoomCreate struct {
	config
	mutation *RoomMutation
	hooks    []Hook
}

// SetUserIds sets the "userIds" field.
func (rc *RoomCreate) SetUserIds(s string) *RoomCreate {
	rc.mutation.SetUserIds(s)
	return rc
}

// SetCreatedAt sets the "created_at" field.
func (rc *RoomCreate) SetCreatedAt(t time.Time) *RoomCreate {
	rc.mutation.SetCreatedAt(t)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableCreatedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetCreatedAt(*t)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RoomCreate) SetUpdatedAt(t time.Time) *RoomCreate {
	rc.mutation.SetUpdatedAt(t)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RoomCreate) SetNillableUpdatedAt(t *time.Time) *RoomCreate {
	if t != nil {
		rc.SetUpdatedAt(*t)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RoomCreate) SetID(u uuid.UUID) *RoomCreate {
	rc.mutation.SetID(u)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *RoomCreate) SetNillableID(u *uuid.UUID) *RoomCreate {
	if u != nil {
		rc.SetID(*u)
	}
	return rc
}

// AddRoomUserIDs adds the "roomUsers" edge to the RoomUser entity by IDs.
func (rc *RoomCreate) AddRoomUserIDs(ids ...int) *RoomCreate {
	rc.mutation.AddRoomUserIDs(ids...)
	return rc
}

// AddRoomUsers adds the "roomUsers" edges to the RoomUser entity.
func (rc *RoomCreate) AddRoomUsers(r ...*RoomUser) *RoomCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return rc.AddRoomUserIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (rc *RoomCreate) AddMessageIDs(ids ...int) *RoomCreate {
	rc.mutation.AddMessageIDs(ids...)
	return rc
}

// AddMessages adds the "messages" edges to the Message entity.
func (rc *RoomCreate) AddMessages(m ...*Message) *RoomCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return rc.AddMessageIDs(ids...)
}

// Mutation returns the RoomMutation object of the builder.
func (rc *RoomCreate) Mutation() *RoomMutation {
	return rc.mutation
}

// Save creates the Room in the database.
func (rc *RoomCreate) Save(ctx context.Context) (*Room, error) {
	var (
		err  error
		node *Room
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Room)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RoomMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RoomCreate) SaveX(ctx context.Context) *Room {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RoomCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RoomCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RoomCreate) defaults() {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		v := room.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		v := room.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := room.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RoomCreate) check() error {
	if _, ok := rc.mutation.UserIds(); !ok {
		return &ValidationError{Name: "userIds", err: errors.New(`ent: missing required field "Room.userIds"`)}
	}
	if v, ok := rc.mutation.UserIds(); ok {
		if err := room.UserIdsValidator(v); err != nil {
			return &ValidationError{Name: "userIds", err: fmt.Errorf(`ent: validator failed for field "Room.userIds": %w`, err)}
		}
	}
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Room.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Room.updated_at"`)}
	}
	return nil
}

func (rc *RoomCreate) sqlSave(ctx context.Context) (*Room, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
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
	return _node, nil
}

func (rc *RoomCreate) createSpec() (*Room, *sqlgraph.CreateSpec) {
	var (
		_node = &Room{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: room.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: room.FieldID,
			},
		}
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rc.mutation.UserIds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: room.FieldUserIds,
		})
		_node.UserIds = value
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: room.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: room.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := rc.mutation.RoomUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.RoomUsersTable,
			Columns: []string{room.RoomUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: roomuser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   room.MessagesTable,
			Columns: []string{room.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RoomCreateBulk is the builder for creating many Room entities in bulk.
type RoomCreateBulk struct {
	config
	builders []*RoomCreate
}

// Save creates the Room entities in the database.
func (rcb *RoomCreateBulk) Save(ctx context.Context) ([]*Room, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Room, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoomMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RoomCreateBulk) SaveX(ctx context.Context) []*Room {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RoomCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RoomCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

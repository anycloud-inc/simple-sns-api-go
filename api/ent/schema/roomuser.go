package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RoomUser holds the schema definition for the RoomUser entity.
type RoomUser struct {
	ent.Schema
}

// Fields of the RoomUser.
func (RoomUser) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the RoomUser.
func (RoomUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("roomUsers").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("roomUsers").
			Unique().
			Required(),
	}
}

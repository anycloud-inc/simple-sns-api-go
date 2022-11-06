package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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
		// association fields
		field.UUID("room_id", uuid.UUID{}),
		field.Int("user_id"),
	}
}

// Edges of the RoomUser.
func (RoomUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Field("room_id").
			Ref("roomUsers").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Field("user_id").
			Ref("roomUsers").
			Unique().
			Required(),
	}
}

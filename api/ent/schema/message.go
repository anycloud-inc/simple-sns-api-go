package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Default("").NotEmpty(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		// association fields
		field.UUID("room_id", uuid.UUID{}),
		field.Int("user_id"),
		field.Int("post_id").Optional(),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Field("room_id").
			Ref("messages").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Field("user_id").
			Ref("messages").
			Unique().
			Required(),
		edge.From("post", Post.Type).
			Field("post_id").
			Ref("messages").
			Unique(),
	}
}

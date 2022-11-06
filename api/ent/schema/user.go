package schema

import (
	"errors"
	"net/mail"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().MaxLen(100),
		field.String("email").NotEmpty().MaxLen(100).Validate(emailValidator),
		field.String("password").NotEmpty().MaxLen(100),
		field.String("iconImageUrl").Default(""),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func emailValidator(s string) error {
	_, err := mail.ParseAddress(s)
	if err != nil {
		return errors.New("invalid email")
	}
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("roomUsers", RoomUser.Type),
		edge.To("messages", Message.Type),
	}
}

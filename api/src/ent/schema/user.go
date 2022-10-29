package schema

import (
	"errors"
	"net/mail"

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
		field.String("name").NotEmpty().MaxLen((100)),
		field.String("email").NotEmpty().MaxLen((100)).Validate(emailValidator),
		field.String("password").NotEmpty().MaxLen((100)),
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
	}
}

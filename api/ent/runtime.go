// Code generated by ent, DO NOT EDIT.

package ent

import (
	"simple_sns_api/ent/message"
	"simple_sns_api/ent/post"
	"simple_sns_api/ent/room"
	"simple_sns_api/ent/roomuser"
	"simple_sns_api/ent/schema"
	"simple_sns_api/ent/user"
	"time"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescContent is the schema descriptor for content field.
	messageDescContent := messageFields[0].Descriptor()
	// message.DefaultContent holds the default value on creation for the content field.
	message.DefaultContent = messageDescContent.Default.(string)
	// message.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	message.ContentValidator = messageDescContent.Validators[0].(func(string) error)
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[1].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	// messageDescUpdatedAt is the schema descriptor for updated_at field.
	messageDescUpdatedAt := messageFields[2].Descriptor()
	// message.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	message.DefaultUpdatedAt = messageDescUpdatedAt.Default.(func() time.Time)
	// message.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	message.UpdateDefaultUpdatedAt = messageDescUpdatedAt.UpdateDefault.(func() time.Time)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescBody is the schema descriptor for body field.
	postDescBody := postFields[0].Descriptor()
	// post.DefaultBody holds the default value on creation for the body field.
	post.DefaultBody = postDescBody.Default.(string)
	// post.BodyValidator is a validator for the "body" field. It is called by the builders before save.
	post.BodyValidator = func() func(string) error {
		validators := postDescBody.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(body string) error {
			for _, fn := range fns {
				if err := fn(body); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[1].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postFields[2].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescUserIds is the schema descriptor for userIds field.
	roomDescUserIds := roomFields[1].Descriptor()
	// room.UserIdsValidator is a validator for the "userIds" field. It is called by the builders before save.
	room.UserIdsValidator = roomDescUserIds.Validators[0].(func(string) error)
	// roomDescCreatedAt is the schema descriptor for created_at field.
	roomDescCreatedAt := roomFields[2].Descriptor()
	// room.DefaultCreatedAt holds the default value on creation for the created_at field.
	room.DefaultCreatedAt = roomDescCreatedAt.Default.(func() time.Time)
	// roomDescUpdatedAt is the schema descriptor for updated_at field.
	roomDescUpdatedAt := roomFields[3].Descriptor()
	// room.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	room.DefaultUpdatedAt = roomDescUpdatedAt.Default.(func() time.Time)
	// room.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	room.UpdateDefaultUpdatedAt = roomDescUpdatedAt.UpdateDefault.(func() time.Time)
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() uuid.UUID)
	roomuserFields := schema.RoomUser{}.Fields()
	_ = roomuserFields
	// roomuserDescCreatedAt is the schema descriptor for created_at field.
	roomuserDescCreatedAt := roomuserFields[0].Descriptor()
	// roomuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	roomuser.DefaultCreatedAt = roomuserDescCreatedAt.Default.(func() time.Time)
	// roomuserDescUpdatedAt is the schema descriptor for updated_at field.
	roomuserDescUpdatedAt := roomuserFields[1].Descriptor()
	// roomuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	roomuser.DefaultUpdatedAt = roomuserDescUpdatedAt.Default.(func() time.Time)
	// roomuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	roomuser.UpdateDefaultUpdatedAt = roomuserDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = func() func(string) error {
		validators := userDescPassword.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(password string) error {
			for _, fn := range fns {
				if err := fn(password); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescIconImageUrl is the schema descriptor for iconImageUrl field.
	userDescIconImageUrl := userFields[3].Descriptor()
	// user.DefaultIconImageUrl holds the default value on creation for the iconImageUrl field.
	user.DefaultIconImageUrl = userDescIconImageUrl.Default.(string)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}

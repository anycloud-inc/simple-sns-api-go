package message

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/ent/post"
	"simple_sns_api/ent/roomuser"
	"simple_sns_api/src/db"
	"simple_sns_api/src/modules/room"

	"github.com/google/uuid"
)

type MessageService struct{}

type CreateParams struct {
	UserId  int
	RoomId  uuid.UUID
	Content string
}

func (s MessageService) Create(ctx context.Context, params CreateParams) (*ent.Message, error) {
	_, err := db.Client.RoomUser.Query().
		Where(roomuser.UserID(params.UserId)).
		Where(roomuser.RoomID(params.RoomId)).
		First(ctx)

	if err != nil {
		return nil, err
	}

	return db.Client.Message.Create().
		SetRoomID(params.RoomId).
		SetUserID(params.UserId).
		SetContent(params.Content).
		Save(ctx)
}

type CreateViaPostParams struct {
	UserId  int
	PostId  int
	Content string
}

func (s MessageService) CreateViaPost(ctx context.Context, params CreateViaPostParams) (*ent.Message, error) {
	post, err := db.Client.Post.Query().Where(post.ID(params.PostId)).First(ctx)
	if err != nil {
		return nil, err
	}
	room, err := room.RoomService{}.FindOrCreate(ctx, post.UserID, params.UserId)
	return s.Create(ctx, CreateParams{UserId: params.UserId, RoomId: room.ID, Content: params.Content})
}

package message

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/ent/message"
	"simple_sns_api/ent/post"
	"simple_sns_api/ent/roomuser"
	"simple_sns_api/src/db"
	"simple_sns_api/src/lib/pagination"
	"simple_sns_api/src/modules/room"

	"github.com/google/uuid"
)

type MessageService struct{}

func (s MessageService) Find(ctx context.Context, roomId string, pagination pagination.Params) ([]*ent.Message, error) {
	println(roomId)
	roomUuid, err := uuid.Parse(roomId)
	if err != nil {
		return nil, err
	}
	query := db.Client.Message.Query().WithUser().WithPost()
	if pagination.Cursor != 0 {
		query = query.Where(message.IDLT(pagination.Cursor))
	}
	if pagination.Size == 0 {
		pagination.Size = 20
	}
	messages, err := query.
		Where(message.RoomIDEQ(roomUuid)).
		Limit(pagination.Size).
		Order(ent.Desc(message.FieldID)).
		All(ctx)
	return messages, err
}

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

package message

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/ent/message"
	"simple_sns_api/ent/roomuser"
	"simple_sns_api/src/db"

	"github.com/google/uuid"
)

type MessageService struct{}

func (s MessageService) FindLatestMessages(ctx context.Context, roomIds []uuid.UUID) ([]*ent.Message, error) {
	var result []struct {
		RoomID uuid.UUID `json:"room_id"`
		Max    int
	}
	err := db.Client.Message.Query().
		Where(message.RoomIDIn(roomIds...)).
		GroupBy(message.FieldRoomID).
		Aggregate(ent.Max(message.FieldID)).
		Scan(ctx, &result)

	if err != nil {
		return nil, err
	}

	latestMessageIds := make([]int, len(result))
	for i, e := range result {
		latestMessageIds[i] = e.Max
	}

	if err != nil {
		return nil, err
	}

	return db.Client.Message.Query().
		Where(message.IDIn(latestMessageIds...)).
		Order(ent.Desc(message.FieldCreatedAt)).
		All(ctx)
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
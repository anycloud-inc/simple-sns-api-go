package room

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/ent/room"
	entRoomUser "simple_sns_api/ent/roomuser"
	"simple_sns_api/src/db"
	"simple_sns_api/src/modules/message"
	"sort"

	"github.com/google/uuid"
)

type RoomService struct{}

func (s RoomService) find(ctx context.Context, userId int) ([]*ent.Room, error) {
	roomUsers, err := db.Client.RoomUser.Query().
		Where(entRoomUser.UserID(userId)).All(ctx)
	roomIds := make([]uuid.UUID, len(roomUsers))
	for i, roomUser := range roomUsers {
		roomIds[i] = roomUser.RoomID
	}

	if err != nil {
		return nil, err
	}

	latestMessages, err := message.MessageService{}.FindLatestMessages(ctx, roomIds)

	if err != nil {
		return nil, err
	}

	for i := range latestMessages {
		roomIds[i] = latestMessages[i].RoomID
	}

	rooms, err := db.Client.Room.Query().
		WithRoomUsers().
		Where(room.IDIn(roomIds...)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	for _, room := range rooms {
		for _, latestMessage := range latestMessages {
			if room.ID == latestMessage.RoomID {
				room.Edges.Messages = []*ent.Message{latestMessage}
			}
		}
	}

	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i].Edges.Messages[0].CreatedAt.After(rooms[j].Edges.Messages[0].CreatedAt)
	})

	return rooms, nil
}

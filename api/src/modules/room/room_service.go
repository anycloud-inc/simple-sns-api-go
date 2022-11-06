package room

import (
	"context"
	"fmt"
	"simple_sns_api/ent"
	"simple_sns_api/ent/message"
	"simple_sns_api/ent/room"
	entRoomUser "simple_sns_api/ent/roomuser"
	"simple_sns_api/src/db"

	"sort"
	"strings"

	"github.com/google/uuid"
)

type RoomService struct{}

func (s RoomService) Find(ctx context.Context, userId int) ([]*ent.Room, error) {
	roomUsers, err := db.Client.RoomUser.Query().
		Where(entRoomUser.UserID(userId)).All(ctx)
	roomIds := make([]uuid.UUID, len(roomUsers))
	for i, roomUser := range roomUsers {
		roomIds[i] = roomUser.RoomID
	}

	if err != nil {
		return nil, err
	}

	latestMessages, err := findLatestMessages(ctx, roomIds)

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

func (s RoomService) FindOne(ctx context.Context, roomId uuid.UUID) (*ent.Room, error) {
	return db.Client.Room.Query().Where(room.ID(roomId)).WithRoomUsers().First(ctx)
}

func (s RoomService) FindOrCreate(ctx context.Context, userIds ...int) (*ent.Room, error) {
	usersId := makeUsersId(userIds)
	room, _ := db.Client.Room.Query().Where(room.UsersId(usersId)).WithRoomUsers().First(ctx)
	if room != nil {
		return room, nil
	}
	return createRoom(ctx, userIds)
}

func createRoom(ctx context.Context, userIds []int) (*ent.Room, error) {
	usersId := makeUsersId(userIds)
	room, err := db.Client.Room.Create().SetUsersId(usersId).Save(ctx)

	bulk := make([]*ent.RoomUserCreate, len(userIds))
	for i, userId := range userIds {
		bulk[i] = db.Client.RoomUser.Create().SetUserID(userId).SetRoomID(room.ID)
	}
	roomUsers, err := db.Client.RoomUser.CreateBulk(bulk...).Save(ctx)
	room.Edges.RoomUsers = roomUsers

	return room, err
}

func makeUsersId(userIds []int) string {
	sort.Ints(userIds)
	userIdStrings := make([]string, len(userIds))
	for i, userId := range userIds {
		userIdStrings[i] = fmt.Sprint(userId)
	}
	return strings.Join(userIdStrings, "-")
}

func findLatestMessages(ctx context.Context, roomIds []uuid.UUID) ([]*ent.Message, error) {
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

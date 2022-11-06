package room

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/src/db"
)

type RoomService struct{}

func (s RoomService) find(ctx context.Context) ([]*ent.Room, error) {
	return db.Client.Room.Query().WithRoomUsers().WithMessages().All(ctx)
}

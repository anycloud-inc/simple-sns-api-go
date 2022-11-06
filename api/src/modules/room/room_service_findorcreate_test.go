package room

import (
	"simple_sns_api/src/db"
	"testing"
)

func TestFindOrCreate(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user1, user2 := createAccountsX(ctx)
	// まだルームがない場合
	room, err := RoomService{}.FindOrCreate(ctx, user1.ID, user2.ID)
	if err != nil {
		t.Error(err)
	}
	if len(room.Edges.RoomUsers) != 2 {
		t.Error("Room User Should Be Created.")
	}

	//  すでにルームがある場合
	room2, err := RoomService{}.FindOrCreate(ctx, user1.ID, user2.ID)
	if err != nil {
		t.Error(err)
	}
	if room2.ID != room.ID {
		t.Error("Room Should Be Found.")
	}
}

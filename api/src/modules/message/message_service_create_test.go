package message

import (
	"fmt"
	"simple_sns_api/src/db"
	"testing"
)

func TestCreate(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	// prepare data
	user1, user2 := createAccountsX(ctx)
	user3, _ := createAccountsX(ctx)
	room := db.Client.Room.Create().SetUsersId(fmt.Sprintf("%d-%d", user1.ID, user2.ID)).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room).SetUser(user1).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room).SetUser(user2).SaveX(ctx)

	// test
	message, err := MessageService{}.Create(ctx, CreateParams{UserId: user1.ID, RoomId: room.ID, Content: "HOGE"})
	if err != nil {
		t.Error(err)
	}
	if message.Content != "HOGE" {
		t.Error("Message Content Should Be Created.")
	}

	message, err = MessageService{}.Create(ctx, CreateParams{UserId: user3.ID, RoomId: room.ID, Content: "HOGE"})
	if err == nil {
		t.Error("Should Be Error When User Is Not In Room.")
	}
}

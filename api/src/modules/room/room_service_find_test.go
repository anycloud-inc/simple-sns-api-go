package room

import (
	"context"
	"fmt"
	"simple_sns_api/ent"
	"simple_sns_api/src/db"
	"simple_sns_api/src/modules/account"
	"testing"
)

func TestFind(t *testing.T) {
	ctx := db.CreateTestConnection(t)

	// prepare data
	user1, user2 := createAccountsX(ctx)
	room := db.Client.Room.Create().SetUserIds(fmt.Sprintf("%d-%d", user1.ID, user2.ID)).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room).SetUser(user1).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room).SetUser(user2).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room).SetContent("HOGE").SetUser(user1).SaveX(ctx)

	// test
	rooms, err := RoomService{}.find(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 1 {
		t.Error("Invalid Rooms Size.")
	}
	if rooms[0].Edges.Messages[0].Content != "HOGE" {
		t.Error("Invalid Message Content.")
	}
	if len(rooms[0].Edges.RoomUsers) != 2 {
		t.Error("Invalid Room User Size.")
	}
}

func createAccountsX(ctx context.Context) (*ent.User, *ent.User) {
	user1, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}
	user2, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Fuga",
		Email:    "fuga@example.com",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}
	return user1, user2
}

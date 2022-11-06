package room

import (
	"context"
	"fmt"
	"math/rand"
	"simple_sns_api/ent"
	"simple_sns_api/src/db"
	"simple_sns_api/src/modules/account"
	"strconv"
	"testing"
	"time"
)

func TestFind(t *testing.T) {
	ctx := db.CreateTestConnection(t)

	// prepare data
	user1, user2 := createAccountsX(ctx)
	user3, user4 := createAccountsX(ctx)
	room1 := db.Client.Room.Create().SetUsersId(fmt.Sprintf("%d-%d", user1.ID, user2.ID)).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room1).SetUser(user1).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room1).SetUser(user2).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room1).SetUser(user1).
		SetContent("HOGE1").SetCreatedAt(time.Now()).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room1).SetUser(user1).
		SetContent("HOGE2").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)

	room2 := db.Client.Room.Create().SetUsersId(fmt.Sprintf("%d-%d", user1.ID, user3.ID)).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room2).SetUser(user1).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room2).SetUser(user3).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room2).SetUser(user1).
		SetContent("FUGA").SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)

	room3 := db.Client.Room.Create().SetUsersId(fmt.Sprintf("%d-%d", user1.ID, user4.ID)).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room3).SetUser(user1).SaveX(ctx)
	db.Client.RoomUser.Create().SetRoom(room3).SetUser(user4).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room3).SetUser(user1).
		SetContent("PIYO").SetCreatedAt(time.Now().Add(-time.Hour)).SaveX(ctx)

	// test
	rooms, err := RoomService{}.Find(ctx, user1.ID)
	if err != nil {
		t.Error(err)
	}
	if len(rooms) != 3 {
		t.Error("Invalid Rooms Size.")
	}
	if rooms[0].ID != room2.ID {
		t.Error("Room With Latest Message Should Be First.")
	}
	if rooms[0].Edges.Messages[0].Content != "FUGA" {
		t.Error("Message Should Be Loaded")
	}
	if len(rooms[0].Edges.RoomUsers) != 2 {
		t.Error("Room User Should Be Loaded.")
	}
	if rooms[1].ID != room1.ID {
		t.Error("Invalid Room ID.")
	}
	if len(rooms[1].Edges.Messages) != 1 {
		t.Error("Only Latest Message Should Be Loaded.")
	}
}

func createAccountsX(ctx context.Context) (*ent.User, *ent.User) {
	rand.Seed(time.Now().UnixNano())
	random := strconv.Itoa(rand.Intn(100000))
	user1, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge" + random,
		Email:    "hoge@example.com" + random,
		Password: "password",
	})
	if err != nil {
		panic(err)
	}
	user2, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Fuga" + random,
		Email:    "fuga@example.com" + random,
		Password: "password",
	})
	if err != nil {
		panic(err)
	}
	return user1, user2
}

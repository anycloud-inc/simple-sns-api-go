package message

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

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

	"github.com/google/uuid"
)

func TestFindLasteMessages(t *testing.T) {
	ctx := db.CreateTestConnection(t)

	// prepare data
	user1, user2 := createAccountsX(ctx)
	user3, user4 := createAccountsX(ctx)
	room1 := db.Client.Room.Create().SetUserIds(fmt.Sprintf("%d-%d", user1.ID, user2.ID)).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room1).SetUser(user1).
		SetContent("HOGE1").SetCreatedAt(time.Now()).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room1).SetUser(user1).
		SetContent("HOGE2").SetCreatedAt(time.Now().Add(time.Minute)).SaveX(ctx)

	room2 := db.Client.Room.Create().SetUserIds(fmt.Sprintf("%d-%d", user1.ID, user3.ID)).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room2).SetUser(user1).
		SetContent("FUGA").SetCreatedAt(time.Now().Add(time.Hour)).SaveX(ctx)

	room3 := db.Client.Room.Create().SetUserIds(fmt.Sprintf("%d-%d", user1.ID, user4.ID)).SaveX(ctx)
	db.Client.Message.Create().SetRoom(room3).SetUser(user1).
		SetContent("PIYO").SetCreatedAt(time.Now().Add(-time.Hour)).SaveX(ctx)

	// test
	roomIds := []uuid.UUID{room1.ID, room2.ID, room3.ID}
	latestMessages, err := MessageService{}.FindLatestMessages(ctx, roomIds)

	if err != nil {
		t.Error(err)
	}

	if len(latestMessages) != 3 {
		t.Errorf("Invalid length: %d", len(latestMessages))
	}

	if latestMessages[0].Content != "FUGA" {
		t.Errorf("Messages should be sorted by created_at desc")
	}

	if latestMessages[1].Content != "HOGE2" {
		t.Errorf("Messages should be sorted by created_at desc")
	}

	if latestMessages[2].Content != "PIYO" {
		t.Errorf("Messages should be sorted by created_at desc")
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

package message

import (
	"simple_sns_api/src/db"
	"simple_sns_api/src/lib/pagination"
	"simple_sns_api/src/modules/account"
	"simple_sns_api/src/modules/room"
	"testing"
)

func TestFind(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "UserName",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})
	if err != nil {
		t.Error(err)
	}

	room, err := room.RoomService{}.FindOrCreate(ctx, user.ID)
	if err != nil {
		t.Error(err)
	}
	_, err = MessageService{}.Create(ctx, CreateParams{UserId: user.ID, RoomId: room.ID, Content: "Hoge"})
	if err != nil {
		t.Error(err)
	}
	_, err = MessageService{}.Create(ctx, CreateParams{UserId: user.ID, RoomId: room.ID, Content: "Fuga"})
	if err != nil {
		t.Error(err)
	}
	messages, err := MessageService{}.Find(ctx, room.ID.String(), pagination.Params{Size: 1})
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 1 {
		t.Error("Invalid size.")
	}
	if messages[0].Content != "Fuga" {
		t.Error("Invalid body.")
	}
	if messages[0].Edges.User.Name != "UserName" {
		t.Error("Invalid User Name.")
	}
	messages, err = MessageService{}.Find(ctx, room.ID.String(), pagination.Params{Cursor: messages[0].ID, Size: 1})
	if err != nil {
		t.Error(err)
	}
	if len(messages) != 1 {
		t.Error("Invalid size.")
	}
	if messages[0].Content != "Hoge" {
		t.Error("Invalid body.")
	}
	if messages[0].Edges.User.Name != "UserName" {
		t.Error("Invalid User Name.")
	}
}

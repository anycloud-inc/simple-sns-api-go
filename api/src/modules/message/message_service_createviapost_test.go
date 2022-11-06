package message

import (
	"simple_sns_api/ent/room"
	"simple_sns_api/src/db"
	"simple_sns_api/src/modules/post"
	"testing"
)

func TestCreateViaPost(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	// prepare data
	user1, user2 := createAccountsX(ctx)
	post, err := post.PostService{}.Create(ctx, post.CreateParams{UserId: user2.ID, Body: "Hoge"})

	// test
	message, err := MessageService{}.CreateViaPost(ctx, CreateViaPostParams{PostId: post.ID, UserId: user1.ID, Content: "Hello"})
	if err != nil {
		t.Error(err)
	}
	if message.Content != "Hello" {
		t.Error("Message Content Should Be Created.")
	}

	room := db.Client.Room.Query().Where(room.ID(message.RoomID)).FirstX(ctx)
	if room == nil {
		t.Error("Room Should Be Created.")
	}

}

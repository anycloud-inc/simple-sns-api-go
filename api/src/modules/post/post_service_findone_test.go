package post

import (
	"simple_sns_api/src/db"
	"simple_sns_api/src/modules/account"
	"testing"
)

func TestFindOne(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})

	post, _ := PostService{}.Create(ctx, CreateParams{UserId: user.ID, Body: "PostBody"})
	post, err = PostService{}.FindOne(ctx, post.ID)
	if err != nil {
		t.Error(err)
	}
	if post.Body != "PostBody" {
		t.Error("Invalid body.")
	}
	if post.Edges.User.Name != "Hoge" {
		t.Error("Invalid User Name.")
	}
}

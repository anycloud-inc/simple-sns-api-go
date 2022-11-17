package post

import (
	"simple_sns_api/src/db"
	"simple_sns_api/src/lib/pagination"
	"simple_sns_api/src/modules/account"
	"testing"
)

func TestFind(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "UserName",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})

	PostService{}.Create(ctx, CreateParams{UserId: user.ID, Body: "Hoge"})
	PostService{}.Create(ctx, CreateParams{UserId: user.ID, Body: "Fuga"})
	posts, err := PostService{}.Find(ctx, pagination.Params{Size: 1})
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 {
		t.Error("Invalid size.")
	}
	if posts[0].Body != "Fuga" {
		t.Error("Invalid body.")
	}
	if posts[0].Edges.User.Name != "UserName" {
		t.Error("Invalid User Name.")
	}
	posts, err = PostService{}.Find(ctx, pagination.Params{Cursor: posts[0].ID, Size: 1})
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 {
		t.Error("Invalid size.")
	}
	if posts[0].Body != "Hoge" {
		t.Error("Invalid body.")
	}
	if posts[0].Edges.User.Name != "UserName" {
		t.Error("Invalid User Name.")
	}
}

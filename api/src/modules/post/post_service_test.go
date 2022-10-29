package post

import (
	"context"
	"simple_sns_api/src/db"
	"simple_sns_api/src/ent"
	"simple_sns_api/src/modules/account"
	"testing"
)

func TestFind(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _ := createAccount(ctx)

	PostService{}.Create(ctx, CreateParams{UserId: user.ID, Body: "Hoge"})
	PostService{}.Create(ctx, CreateParams{UserId: user.ID, Body: "Fuga"})
	posts, err := PostService{}.find(ctx, PaginationParams{Size: 1})
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 {
		t.Error("Invalid size.")
	}
	if posts[0].Body != "Fuga" {
		t.Error("Invalid body.")
	}
	posts, err = PostService{}.find(ctx, PaginationParams{Cursor: posts[0].ID, Size: 1})
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 {
		t.Error("Invalid size.")
	}
	if posts[0].Body != "Hoge" {
		t.Error("Invalid body.")
	}
}

func TestCreateSuccess(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, err := createAccount(ctx)

	if err != nil {
		t.Error(err)
	}
	post, err := PostService{}.Create(ctx, CreateParams{
		UserId: user.ID,
		Body:   "Hoge",
	})
	if post.Body != "Hoge" {
		t.Error("Invalid body.")
	}
}

func TestCreateFailure(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, err := createAccount(ctx)

	if err != nil {
		t.Error(err)
	}
	// Bodyが空なので失敗
	post, err := PostService{}.Create(ctx, CreateParams{
		UserId: user.ID,
		Body:   "",
	})
	if post != nil {
		t.Error("Post should be nil")
	}
	if err == nil {
		t.Error("Err should exist")
	}
	// UserIdが不正なので失敗
	post, err = PostService{}.Create(ctx, CreateParams{
		UserId: -100,
		Body:   "Hoge",
	})
	if post != nil {
		t.Error("Post should be nil")
	}
	if err == nil {
		t.Error("Err should exist")
	}
}

func createAccount(ctx context.Context) (*ent.User, error) {
	user, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})
	return user, err
}

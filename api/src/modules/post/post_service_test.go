package post

import (
	"context"
	"simple_sns_api/src/db"
	"simple_sns_api/src/ent"
	"simple_sns_api/src/modules/account"
	"testing"
)

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
}

func createAccount(ctx context.Context) (*ent.User, error) {
	user, _, err := account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})
	return user, err
}

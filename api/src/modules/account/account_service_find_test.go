package account

import (
	"simple_sns_api/src/db"
	"testing"
)

func TestFind(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _, _ := AccountService{}.Register(ctx, RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})

	user, err := AccountService{}.Find(ctx, user.ID)

	if err != nil {
		t.Error(err)
	}
	if user.Name != "Hoge" {
		t.Error("Invalid name.")
	}
}

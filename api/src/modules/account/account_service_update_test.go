package account

import (
	"simple_sns_api/src/db"
	"testing"
)

func TestUpdate(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _, _ := AccountService{}.Register(ctx, RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})

	user, err := AccountService{}.Update(ctx, user.ID, UpdateParams{Name: "Fuga", Email: "fuga@example.com"})

	if err != nil {
		t.Error(err)
	}
	if user.Name != "Fuga" {
		t.Error("Invalid name.")
	}
	if user.Email != "fuga@example.com" {
		t.Error("Invalid Email.")
	}
}

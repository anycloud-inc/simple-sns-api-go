package account

import (
	"simple_sns_api/src/db"
	"testing"
)

func TestRegister(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, authToken, err := AccountService{}.Register(ctx, RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})

	t.Logf("user: %v", user)
	t.Logf("authToken: %v", authToken)

	if err != nil {
		t.Error(err)
	}
	if user.Name != "Hoge" {
		t.Error("Invalid name.")
	}
	if authToken == "" {
		t.Error("Invalid authToken.")
	}
	if user.Password == "hogehoge" {
		t.Error("Password is not encrypted.")
	}
}

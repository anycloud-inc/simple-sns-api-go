package auth

import (
	"simple_sns_api/db"
	"simple_sns_api/modules/account"
	"testing"
)

func TestSignInSuccess(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})
	user, authToken, err := AuthService{}.SignIn(ctx, SignInParams{
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
}

func TestSignInFailure(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	account.AccountService{}.Register(ctx, account.RegisterParams{
		Name:     "Hoge",
		Email:    "hoge@example.com",
		Password: "hogehoge",
	})
	user, authToken, err := AuthService{}.SignIn(ctx, SignInParams{
		Email:    "hoge@example.com",
		Password: "fugafuga",
	})

	t.Logf("error: %v", err)

	if err == nil {
		t.Error("Error should be occured.")
	}
	if user != nil {
		t.Error("User should be nil.")
	}
	if authToken != "" {
		t.Error("Auth token should be nil.")
	}
}

package auth

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/ent/user"
	"simple_sns_api/src/db"
	"simple_sns_api/src/lib/auth"
)

type AuthService struct{}

type SignInParams struct {
	Email    string
	Password string
}

func (as AuthService) SignIn(ctx context.Context, params SignInParams) (*ent.User, auth.AuthToken, error) {
	user, err := db.Client.User.Query().Where(user.Email(params.Email)).First(ctx)
	if err != nil {
		return nil, "", err
	}
	err = auth.ComparePassword(user.Password, params.Password)

	if err != nil {
		return nil, "", err
	}
	authToken, err := auth.MakeAuthToken(user.ID)

	return user, authToken, nil
}

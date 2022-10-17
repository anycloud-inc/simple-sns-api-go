package account

import (
	"context"
	"simple_sns_api/db"
	"simple_sns_api/ent"
	"simple_sns_api/lib/auth"
)

type AccountService struct{}

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

func (as AccountService) Register(ctx context.Context, params RegisterParams) (*ent.User, auth.AuthToken, error) {
	encrypted, err := auth.EncryptPassword(params.Password)
	if err != nil {
		return nil, "", err
	}
	user, err := db.Client.User.Create().
		SetName(params.Name).
		SetEmail(params.Email).
		SetPassword(string(encrypted)).
		Save(ctx)
	if err != nil {
		return nil, "", err
	}
	authToken, err := auth.MakeAuthToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, authToken, nil
}

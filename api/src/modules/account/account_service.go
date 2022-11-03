package account

import (
	"context"
	"simple_sns_api/src/db"
	"simple_sns_api/src/ent"
	"simple_sns_api/src/ent/user"
	"simple_sns_api/src/lib/auth"
)

type AccountService struct{}

type RegisterParams struct {
	Name     string
	Email    string
	Password string
}

type UpdateParams struct {
	Name  string
	Email string
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

func (as AccountService) Update(ctx context.Context, userId int, params UpdateParams) (*ent.User, error) {
	qb := db.Client.User.UpdateOneID(userId)
	if params.Name != "" {
		qb = qb.SetName(params.Name)
	}
	if params.Email != "" {
		qb = qb.SetEmail(params.Email)
	}
	err := qb.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return as.Find(ctx, userId)
}

func (as AccountService) Find(ctx context.Context, userId int) (*ent.User, error) {
	return db.Client.User.Query().Where(user.ID(userId)).First(ctx)
}

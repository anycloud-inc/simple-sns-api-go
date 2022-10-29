package post

import (
	"context"
	"simple_sns_api/src/db"
	"simple_sns_api/src/ent"
)

type PostService struct{}

type CreateParams struct {
	UserId int
	Body   string
}

func (s PostService) Create(ctx context.Context, params CreateParams) (*ent.Post, error) {
	post, err := db.Client.Post.Create().
		SetBody(params.Body).
		SetUserID(params.UserId).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return post, nil
}

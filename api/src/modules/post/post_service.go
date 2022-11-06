package post

import (
	"context"
	"simple_sns_api/ent"
	"simple_sns_api/ent/post"
	"simple_sns_api/src/db"
)

type PostService struct{}

type PaginationParams struct {
	Cursor int
	Size   int
}

type CreateParams struct {
	UserId int
	Body   string
}

func (s PostService) find(ctx context.Context, pagination PaginationParams) ([]*ent.Post, error) {
	query := db.Client.Post.Query().WithUser()
	if pagination.Cursor != 0 {
		query = query.Where(post.IDLT(pagination.Cursor))
	}
	if pagination.Size == 0 {
		pagination.Size = 20
	}
	posts, err := query.
		Limit(pagination.Size).
		Order(ent.Desc(post.FieldID)).
		All(ctx)
	return posts, err
}

func (s PostService) findOne(ctx context.Context, id int) (*ent.Post, error) {
	return db.Client.Post.Query().WithUser().Where(post.ID(id)).First(ctx)
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

func (s PostService) Delete(ctx context.Context, id int) error {
	return db.Client.Post.DeleteOneID(id).Exec(ctx)
}

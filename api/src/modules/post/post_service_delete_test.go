package post

import (
	"simple_sns_api/src/db"
	"testing"
)

func TestDelete(t *testing.T) {
	ctx := db.CreateTestConnection(t)
	user, _ := createAccount(ctx)

	post, _ := PostService{}.Create(ctx, CreateParams{UserId: user.ID, Body: "PostBody"})
	PostService{}.Delete(ctx, post.ID)
	post, _ = PostService{}.FindOne(ctx, post.ID)
	if post != nil {
		t.Error("Post should be deleted")
	}
}

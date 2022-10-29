package post

import (
	"net/http"
	"simple_sns_api/src/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type Controller struct{}

// GET /posts
func (_ Controller) Index(ctx *gin.Context) {
	// TODO: pagination
	posts := db.Client.Post.Query().AllX(ctx)
	ctx.JSON(http.StatusOK, gin.H{"posts": posts})
}

// POST /posts
func (_ Controller) Create(ctx *gin.Context) {
	body := ctx.PostFormMap("post")["body"]
	post, err := db.Client.Post.Create().SetBody(body).Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, post)
}

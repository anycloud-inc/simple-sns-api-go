package post

import (
	"net/http"
	"simple_sns_api/db"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type PostController struct{}

// GET /posts
func (pc PostController) Index(ctx *gin.Context) {
	post := []Post{
		{Id: 1, Body: "Hello World!"},
		{Id: 2, Body: "Hello Gin!!"},
	}
	ctx.JSON(http.StatusOK, post)
}

// POST /posts
func (pc PostController) Create(ctx *gin.Context) {
	body := ctx.PostFormMap("post")["body"]
	post, err := db.Client.Post.Create().SetBody(body).Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, post)
}

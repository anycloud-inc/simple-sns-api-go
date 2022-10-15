package post

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Post struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type PostController struct{}

// GET /posts
func (pc PostController) Index(c *gin.Context) {
	post := []Post{
		{Id: 1, Body: "Hello World!"},
		{Id: 2, Body: "Hello Gin!"},
	}
	c.JSON(http.StatusOK, post)
}

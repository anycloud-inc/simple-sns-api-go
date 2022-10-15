package main

import (
	"simple_sns_api/modules/post"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	engine := gin.Default()

	{
		p := engine.Group("/posts")
		controller := post.PostController{}
		p.GET("", controller.Index)
	}

	return engine
}

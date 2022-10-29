package main

import (
	"simple_sns_api/src/modules/post"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	engine := gin.Default()

	{
		p := engine.Group("/posts")
		controller := post.Controller{}
		p.GET("", controller.Index)
		p.POST("", controller.Create)
	}

	return engine
}

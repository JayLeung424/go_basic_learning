package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由和路由组

func main() {
	r := gin.Default()
	// 普通路由
	r.GET("/index", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"method": "GET"})
	})

	// Any 请求方法大集合(杂烩)
	r.Any("/user", func(context *gin.Context) {
		switch context.Request.Method {
		case http.MethodGet:
			context.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			context.JSON(http.StatusOK, gin.H{"method": "POST"})
		}
	})
	// NoRoute
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"message": "default page"})
	})

	// 路由组
	videoGroup := r.Group("/video")
	{
		// 实际的url = /video/index
		videoGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"message": "index"})

		})
	}

	r.Run()
}

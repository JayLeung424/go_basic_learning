package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取path上的参数

func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		context.JSON(http.StatusOK,
			gin.H{
				"URI-name": name,
				"URI-age":  age,
			})
	})

	r.Run()
}

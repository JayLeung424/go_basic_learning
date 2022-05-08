package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取form表单的参数

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html")
	r.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", nil)
	})
	// login post
	r.POST("/login", func(context *gin.Context) {
		// byKey
		username := context.PostForm("username")
		password := context.PostForm("password")
		// 默认值 DefaultPostForm
		xxx := context.DefaultPostForm("xxx", "xxx")
		// 状态
		ace, ok := context.GetPostForm("ace")
		if !ok {
			ace = "ace"
		}
		context.JSON(
			http.StatusOK,
			gin.H{
				"username": username,
				"password": password,
				"xxx":      xxx,
				"ace":      ace,
			})
	})

	r.Run()
}

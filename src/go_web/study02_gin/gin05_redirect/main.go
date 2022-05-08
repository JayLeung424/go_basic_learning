package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 重定向

func main() {

	r := gin.Default()

	// 重定向
	r.GET("/index", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://baidu.com")
	})

	// 路由重定向
	r.GET("/a", func(context *gin.Context) {
		// 访问a 跳转到  /b 对应的路由处理函数
		context.Request.URL.Path = "/b" //指定url
		r.HandleContext(context)
	})
	r.GET("/b", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"msg": "ok",
			})
	})

	r.Run()
}

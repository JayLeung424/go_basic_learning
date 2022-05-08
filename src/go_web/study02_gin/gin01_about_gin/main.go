package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 返回默认的路由路径
	r := gin.Default()

	// restful开发规范
	// GET 请求 path: /hello
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "hello gin",
			})
	})

	// 启动服务
	r.Run() // 默认8080 ->可以设置修改
}

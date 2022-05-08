package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回json字段

func main() {
	// 返回默认的路由路径
	r := gin.Default()

	// 方式1 使用map or gin.H
	r.GET("/json/map", func(context *gin.Context) {
		// data := map[string]interface{}{
		// 	"name":    "jay",
		// 	"age":     22,
		// 	"message": "hello go !",
		// }
		data := gin.H{
			"name":    "jay",
			"age":     22,
			"message": "hello go !",
		}
		context.JSON(
			http.StatusOK,
			data)
	})

	// 方式2 结构体
	type msg struct {
		Name, Message string
		Age           int
	}
	r.GET("/json/struct", func(context *gin.Context) {
		data := msg{
			Name:    "jay",
			Age:     22,
			Message: "hello go !",
		}
		context.JSON(
			http.StatusOK,
			data)
	})

	// 启动服务
	r.Run() // 默认8080 ->可以设置修改
}

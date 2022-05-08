package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 参数绑定  将请求参数和结构体进行绑定

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pass"`
}

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("./index.html")

	// 原始方法 使用query一一赋值
	r.GET("/user", func(context *gin.Context) {
		username := context.Query("username")
		password := context.Query("password")
		user := UserInfo{
			username,
			password,
		}
		fmt.Printf("user: %#v\n", user)
		context.JSON(
			http.StatusOK,
			gin.H{
				"message": "ok",
			})
	})

	r.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// shouldBind 方法
	r.GET("/user/should_bind", func(context *gin.Context) {
		var u UserInfo // 声明一个userinfo类型的变量
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		} else {
			fmt.Printf("user: %#v\n", u)
			context.JSON(
				http.StatusOK,
				gin.H{
					"message": "ok",
				})
		}
	})

	r.POST("/form", func(context *gin.Context) {
		var u UserInfo // 声明一个userinfo类型的变量
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		} else {
			fmt.Printf("user: %#v\n", u)
			context.JSON(
				http.StatusOK,
				gin.H{
					"message": "ok",
				})
		}
	})
	r.POST("/json", func(context *gin.Context) {
		var u UserInfo // 声明一个userinfo类型的变量
		err := context.ShouldBind(&u)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		} else {
			fmt.Printf("user: %#v\n", u)
			context.JSON(
				http.StatusOK,
				gin.H{
					"message": "ok",
				})
		}
	})

	r.Run()
}

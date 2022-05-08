package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 中间件

func indexHandle(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"msg": "index"})
}

// m1 定一个一个中间件
func m1(c *gin.Context) {
	// 计算时间
	start := time.Now()
	fmt.Println("m1 run ...")
	cost := time.Since(start)
	c.Next() // 调用后续的函数
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
}

// m2 也是一个中间件
func m2(c *gin.Context) {
	// 计算时间
	start := time.Now()
	fmt.Println("m2 run ...")
	cost := time.Since(start)
	c.Next() // 调用后续的函数
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m2 out ...")
}

// authMiddleware 登陆权限校验
func authMiddleware(doCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 或者处理一些其他准备工作
	return func(c *gin.Context) {
		if doCheck {
			// 存放具体的逻辑
			// 是否登陆判断
			// if 登陆用户
			// c.Next()
			// else
			// c.Abort()
		}
	}
}

func main() {
	r := gin.Default()
	// 全局注册中间件函数
	r.Use(m1, m2, authMiddleware(false))

	// GET(relativePath string, handlers ...HandlerFunc) IRoutes
	r.GET("/index", indexHandle)
	r.GET("/user", indexHandle)

	// 路由组注册中间件方法1
	xxGroup := r.Group("/xx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	// 路由组注册中间件方法2
	xx2Group := r.Group("/xx")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{"msg": "xxGroup"})
		})
	}
	r.Run(":9001")
}

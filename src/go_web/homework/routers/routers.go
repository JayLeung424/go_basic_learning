package routers

import (
	"github.com/gin-gonic/gin"
	"go_web_learning/go_web/homework/controller"
)

// SetupRouter  存放路由
func SetupRouter() *gin.Engine {

	r := gin.Default()
	// 静态文件加载
	r.StaticFile("/static", "static")
	// 模版文件加载
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	// v1版本接口
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看 所有的待办事项
		v1Group.POST("/todo/list", controller.GetTodoList)
		// 查看 单个
		v1Group.GET("/todo/:id", controller.GetTodoById)
		// 修改
		v1Group.PUT("/todo/:id", controller.UpdateTodoById)
		// 删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodoById)
	}
	return r
}

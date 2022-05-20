package controller

import (
	"github.com/gin-gonic/gin"
	"go_web_learning/go_web/homework/models"
	"net/http"
)

/**
URL -> controller -> service -> model(具体增删改查)
*/

// IndexHandler index页面的跳转
func IndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}

// CreateTodo 创建一个todo
func CreateTodo(context *gin.Context) {
	// 从请求中拿出数据
	var todo models.Todo
	context.ShouldBindJSON(&todo) // ShouldBindJSON 从json中拿到结果值
	// 存入数据库
	err := models.CreateTodo(&todo)
	if err != nil {
		// 返回响应
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	// 返回响应
	context.JSON(http.StatusOK, gin.H{"message": todo})
}

// GetTodoList 获取todo的list
func GetTodoList(context *gin.Context) {
	todoList, err := models.GetTodoList()
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": todoList})
}

// GetTodoById 根据id获取todo
func GetTodoById(context *gin.Context) {
	// id的判断
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	// 查询该id是否存在
	todo, err := models.GetTodoById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": todo})
}

// UpdateTodoById 根据id修改todo
func UpdateTodoById(context *gin.Context) {
	// id的判断
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	// 查询该id是否存在
	todo, err := models.GetTodoById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	// 修改状态
	context.ShouldBindJSON(&todo)
	err = models.UpdateTodoById(id, todo)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"data": todo})
}

// DeleteTodoById 根据id删除todo
func DeleteTodoById(context *gin.Context) {
	// id的判断
	id, ok := context.Params.Get("id")
	if !ok {
		context.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	// 查询该id是否存在
	_, err := models.GetTodoById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": "该todo不存在,无法删除！"})
	}
	err = models.DeleteTodoById(id)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"message": "编号:" + id + "的todo已删除"})
}

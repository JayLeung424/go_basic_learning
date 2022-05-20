package models

import (
	"go_web_learning/go_web/homework/dao"
)

// Todo 建立结构体
type Todo struct {
	ID     int    `json:"id""`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// TableName 指定表名
func (Todo) TableName() string {
	return "t_bubble"
}

// CreateTodo 创建一个todo
func CreateTodo(todo *Todo) (err error) {
	// 存入数据库
	err = dao.DB.Create(&todo).Error
	if err != nil {
		return err
	}
	return
}

// GetTodoById 创建一个todo
func GetTodoById(id string) (todo *Todo, err error) {
	// 存入数据库
	err = dao.DB.Where("id=?", id).First(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// GetTodoList 创建一个todo
func GetTodoList() (todoList []*Todo, err error) {
	err = dao.DB.Debug().Find(&todoList).Error
	if err != nil {
		return nil, err
	}
	return todoList, nil
}

// UpdateTodoById 创建一个todo
func UpdateTodoById(id string, todo *Todo) (err error) {
	// 存入数据库
	err = dao.DB.Debug().Model(&todo).Where("id=?", id).Update(todo).Error
	if err != nil {
		return err
	}
	return
}

// DeleteTodoById 创建一个todo
func DeleteTodoById(id string) (err error) {
	// 存入数据库
	err = dao.DB.Debug().Where("id=?", id).Unscoped().Delete(&Todo{}).Error
	if err != nil {
		return err
	}
	return
}

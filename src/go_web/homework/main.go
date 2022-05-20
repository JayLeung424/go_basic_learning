package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go_web_learning/go_web/homework/dao"
	"go_web_learning/go_web/homework/models"
	"go_web_learning/go_web/homework/routers"
)

func main() {
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.CloseMySQL()
	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{})

	routers.SetupRouter().Run()
}

package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 基础使用

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "root:root1234@(1.116.114.133:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表 自动迁移( 把结构体和数据表进行对应 )
	db.AutoMigrate(&UserInfo{})

	// 创建记录
	// u1 := UserInfo{1, "jay", "男", "swim"}
	// db.Create(u1)

	// 查询
	var u UserInfo
	first := db.First(&u)
	fmt.Println(first)
	fmt.Println(u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	db.Delete(&u)
}

package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// User 定义模型
type User struct {
	gorm.Model   // 内置的gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"` // 设置类型
	Role         string  `gorm:"size:255"`                       // 设置字段大小
	MemberNumber *string `gorm:"unique;not null"`                // 设置唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`                 // 自增类型
	Address      string  `gorm:"index:addr"`                     // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`                              // 忽略本字段

}

// Animal 主键
type Animal struct {
	// 设置主键
	// ID string // 监测到ID 默认主键
	AnimalID string `gorm:"primary_key"` // 指定主键
	Name     string
	Age      int
}

// TableName 将 User 的表名设置为 `profiles`
func (User) TableName() string {
	return "profiles"
}
func main() {

	// 默认表名称规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}

	db, err := gorm.Open("mysql", "root:root1234@(1.116.114.133:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Animal{})
	// 表名禁用复数
	db.SingularTable(true)

	// 使用User结构体创建名为`t_user`的表
	db.Table("t_user").CreateTable(&User{})

}

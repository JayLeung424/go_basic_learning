package dao

import "github.com/jinzhu/gorm"

// 存放mysql相关操作

var (
	DB *gorm.DB
)

// InitMySQL 初始化mysql 并建立连接
func InitMySQL() (err error) {
	// 默认表名称规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	// 连接数据库
	DB, err = gorm.Open("mysql", "root:root1234@(1.116.114.133:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	// 表名禁用复数
	DB.SingularTable(true)
	return DB.DB().Ping()
}

// CloseMySQL 关闭mysql
func CloseMySQL() {
	DB.Close() // 程序退出 关闭数据库连接
}

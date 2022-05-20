package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
参考文档: https://www.liwenzhou.com/posts/Go/gorm_crud/#autoid-1-2-0
*/

// AdminUser 实体类
type AdminUser struct {
	ID    int64
	Name  string         `gorm:"default:'默认值'"`
	NName *string        `gorm:"default:'默认值'"` // 使用指针方式实现零值存入数据库
	WName sql.NullString `gorm:"default:'默认值'"` // 使用Scanner/Valuer接口方式实现零值存入数据库

	Age int16
}

// User 实体类
type User struct {
	gorm.Model
	Name string `gorm:"default:'默认值'"`
	Age  int16
}

// Result 查询后拼装的对象 DTO
type Result struct {
	Name string
	Age  int
}

func main() {

	// 1、连接Mysql数据库 初始化表规则
	db, err := gorm.Open("mysql", "root:root1234@(1.116.114.133:13306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 默认表名称规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "t_" + defaultTableName
	}
	// 关闭复数
	db.SingularTable(true)

	// 2、将模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3、创建数据
	// create(db)

	// 4、查询数据
	read(db)
}

// create 创建数据
func create(db *gorm.DB) {
	u := User{Name: "jay", Age: 22} // 创建AdminUser对象
	fmt.Println(db.NewRecord(&u))   // 判断主键是否为空
	db.Create(&u)

	// 默认值
	u2 := User{Age: 20}
	db.Create(&u2)

	// 例如PostgreSQL数据库中可以使用下面的方式实现合并插入, 有则更新, 无则插入。
	// 为Insert语句添加扩展SQL选项
	db.Set("gorm:insert_option", "ON CONFLICT").Create(&u)
}

// read 查询数据
func read(db *gorm.DB) {

	var user User // 声明模型结构体类型变量user
	// 查询到第一条记录 SELECT * FROM users ORDER BY id LIMIT 1;
	db.First(&user) // id必须是int类型
	fmt.Printf("user:%#v\n", user)
	// 随机获取一条记录 SELECT * FROM users LIMIT 1;
	db.Take(&user)
	// 根据主键查询最后一条记录 SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)

	// 查询所有的记录 SELECT * FROM users;
	var users []User
	db.Debug().Find(&users)
	fmt.Printf("users:%#v\n", users)

	// Struct  SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;
	db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)

	// Scan 将查询出来的对象 映射到result中
	var result Result
	// 连表查询
	db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Find(&result)

	// 链式查询
	tx := db.Where("name = ?", "jz")

	// 添加更多条件
	if true {
		tx = tx.Where("age = ?", 20)
	} else {
		tx = tx.Where("age = ?", 30)
	}
	if false {
		tx = tx.Where("active = ?", 1)
	}
	tx.Find(&user)
}

// update 修改数据
func update(db *gorm.DB) {
	var user User
	db.First(&user)

	user.Name = "七米"
	user.Age = 99
	// 默认会更新该对象的所有字段，即使你没有赋值。 UPDATE `users` SET `created_at` = '2020-02-16 12:52:20', `updated_at` = '2020-02-16 12:54:55', `deleted_at` = NULL, `name` = '七米', `age` = 99, `active` = true  WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 1
	db.Save(&user)

	// 更新单个属性，如果它有变化 UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Update("name", "hello")

	// 根据给定的条件更新单个属性 UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
	db.Model(&user).Where("active = ?", true).Update("name", "hello")

	// 使用 map 更新多个属性，只会更新其中有变化的属性 UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})

	// 使用 struct 更新多个属性，只会更新其中有变化且为非零值的字段 UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;
	db.Model(&user).Updates(User{Name: "hello", Age: 18})

	// 警告：当使用 struct 更新时，GORM只会更新那些非零值的字段
	// 对于下面的操作，不会发生任何更新，"", 0, false 都是其类型的零值
	// db.Model(&user).Updates(User{Name: "", Age: 0, Active: false})
}

// delete 修改数据
func delete(db *gorm.DB) {
	var user User
	// 软删除 - 批量删除  UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;
	db.Where("age = ?", 20).Delete(&User{})

	// Unscoped 方法可以物理删除记录  DELETE FROM orders WHERE id=10;
	db.Unscoped().Delete(&user)
}

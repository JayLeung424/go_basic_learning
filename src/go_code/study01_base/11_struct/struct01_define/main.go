package main

import "fmt"

// 结构体

// 定义结构体
type person struct {
	// 结构体是内存对其的 相同类型的放一起 会节省内存
	name, city string
	age        int
}

func main() {
	var p person // p: 变量名 person: 变量类型
	p.name = "jay"
	p.city = "sh"
	p.age = 22
	fmt.Println(p)

	// 匿名结构体
	var user struct {
		name string
	}
	user.name = "jay"
	fmt.Println(user)

}

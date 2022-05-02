package main

import "fmt"

// 结构体指针
type person struct {
	name, city string
	age        int
}

func main() {
	var p = new(person)
	fmt.Printf("%T\n", p) // *main.person
	// 写法1
	//(*p).name = "jay"
	//(*p).city = "sh"
	//(*p).age = 22
	// 写法2 语法堂 自动识别
	p.name = "jay"
	p.city = "sh"
	p.age = 22

	fmt.Printf("%#v\n", p) // &main.person{name:"jay", city:"sh", age:22}

	// 取结构体的地址进行示例化
	p2 := &person{}
	fmt.Printf("%T\n", p2)  // *main.person
	fmt.Printf("%#v\n", p2) // &main.person{name:"", city:"", age:0}
	p2.name = "jay"
	fmt.Printf("%T\n", p2)  // *main.person
	fmt.Printf("%#v\n", p2) // &main.person{name:"jay", city:"", age:0}
}

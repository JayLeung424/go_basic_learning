package main

import "fmt"

// 结构体构造函数: 构造一个结构体实例的函数
// 结构体是值类型

type person struct {
	name, city string
	age        int
}

// 构造函数
func newPerson(name, city string, age int) person {
	return person{
		name: name,
		city: city,
		age:  age,
	}
}

// 构造函数 当结构体很大的时候 我们通常返回指针类型 可以保证程序的性能
func newPerson2(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}
func main() {
	p := newPerson("jay", "sh", 22)
	p1 := newPerson2("jay", "sh", 22)
	fmt.Printf("%T\n", p)  // *main.person
	fmt.Printf("%#v\n", p) // &main.person{name:"", city:"", age:0}

	fmt.Printf("%T\n", p1)  // *main.person
	fmt.Printf("%#v\n", p1) // &main.person{name:"", city:"", age:0}
}

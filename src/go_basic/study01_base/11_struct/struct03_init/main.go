package main

import "fmt"

// 结构体的初始化

type person struct {
	// 结构体是内存对其的 相同类型的放一起 会节省内存
	name, city string
	age        int
}

func main() {
	// 1. 键值对初始化
	p1 := person{
		name: "jay",
		city: "sh",
		age:  22,
	}
	fmt.Printf("%#v\n", p1)

	p2 := &person{
		name: "jay",
		city: "sh",
		age:  22,
	}
	fmt.Printf("%#v\n", p2)

	// 2. 值列表进行初始化  根据结构体里面的顺序初始化(必须都要有)
	p3 := person{
		"jay", "sh", 22,
	}
	fmt.Printf("%#v\n", p3)

}

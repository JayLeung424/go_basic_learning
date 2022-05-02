package main

import "fmt"

// 使用值接受者实现接口 和 使用指针接收者实现接口 的区别

// 接口嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}
type sayer interface {
	say()
}
type person struct {
	name string
	age  int
}

// 使用值接受者实现接口：类型的值和类型的指针都可以保存到接口变量中
// func (p person) move() {
// 	fmt.Printf("%s在跑步", p.name)
// }

// 使用指针接受者实现接口： 只有类型指针能够保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑步", p.name)
}

func (p *person) say() {
	fmt.Printf("%s在跑步", p.name)
}

func main() {
	var m mover
	// p1 := person{ // person类型的值
	// 	name: "jay",
	// 	age:  18,
	// }
	// m = p1  // p1是person类型的值，没有实现move接口
	// m.move()
	// fmt.Println(p1)
	// fmt.Println(m)

	p2 := &person{ // person类型的指针
		name: "josiah",
		age:  22,
	}
	m = p2
	m.move()
	fmt.Println(m)

	var a animal
	a = p2
	a.say()
	a.move()

}

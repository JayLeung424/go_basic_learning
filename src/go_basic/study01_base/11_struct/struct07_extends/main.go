package main

import "fmt"

// 结构体的继承

// Animal 动物 父类
type Animal struct {
	name string
}

// move 动物的公共方法
func (a *Animal) move() {
	fmt.Printf("%s会动～～", a.name)
}

// Dog 使用指针嵌套 实现继承
type Dog struct {
	Feet    int
	*Animal // 匿名嵌套 嵌套的是指针
}

// wang Dog的方法 汪汪叫
func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪～～", d.name)
}

func main() {
	d := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "lele",
		},
	}
	d.wang()
	d.move()
}

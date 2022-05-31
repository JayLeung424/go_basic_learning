package main

import (
	"fmt"
)

/**
 * 接口是一个类型  一种抽象的类型
 * type 接口类型名 interface{
 *   方法名1( 参数列表1 ) 返回值列表1
 *   方法名2( 参数列表2 ) 返回值列表2
 *    …
 * }
 */

// 为什么需要接口

// dog 结构体-狗
type dog struct{}

func (d dog) say() {
	fmt.Println(" 汪汪汪～～～")
}

// cat 结构体-猫
type cat struct{}

func (c cat) say() {
	fmt.Println(" 喵喵喵～～～")
}

// person
type person struct {
	name string
}

func (p person) say() {
	fmt.Println(" 哒咩～～～")
}

// 接口不管你是什么类型 只关心实现什么方法
// sayer 定义一个类型 一个抽象的类型  只要实现了say这个方法的类型 都可以成为sayer
type sayer interface {
	say()
}

// fit 击打
func fit(arg sayer) {
	// 无论传进来的是什么。我都要打ta 打ta ta就会叫, 需要执行ta的say()方法
	arg.say()
}

func main() {
	c1 := cat{}
	fit(c1)

	d1 := dog{}
	fit(d1)

	p1 := person{
		name: "jay",
	}
	fit(p1) // 因为person没有实现say的方法 所以不是sayer类型 无法当作参数传入fit
	//

	var s sayer
	c2 := cat{}
	s = c2
	fmt.Println(s)
}

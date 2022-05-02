package main

import "fmt"

//  为任意类型添加方法

// MyInt 基于内置的基本类型自定义类型
type MyInt int

// 不允许给基本类型定义方法 但是可以通过自定义类型
//func (i int)sayHi() {
func (i MyInt) sayHi() {
	fmt.Println("Hi~")
}

func main() {
	var i MyInt
	i.sayHi()
}

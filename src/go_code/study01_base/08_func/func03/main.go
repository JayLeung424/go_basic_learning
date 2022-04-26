package main

import (
	"fmt"
	"strings"
)

// 匿名函数 和 闭包
func main() {
	// 匿名且直接执行
	func() {
		fmt.Println("这是一个匿名函数")
	}()

	// 接收到返回的函数
	r := a() // 此时是一个闭包
	// 执行了a函数定义的返回匿名函数
	r()

	r1 := a1("jay")
	r1()

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}

// 闭包 = 函数+外层变量的引用
// a 定义一个函数返回一个函数
func a() func() {
	name := "jay"
	return func() {
		fmt.Println("hello", name)
	}
}
func a1(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}
}

// 闭包demo
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

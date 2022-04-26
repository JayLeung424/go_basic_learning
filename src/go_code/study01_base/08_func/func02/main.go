package main

import (
	"errors"
	"fmt"
)

// 函数进阶 - 变量作用域

func main() {
	testGlobal()
	// 外层不能访问函数内部的变量
	// fmt.Println(name)

	// 函数当成参数传入
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30
}

// num: 定义全局变量
var num = 10

// testGlobal  定义函数
func testGlobal() {
	// 可以在函数中访问全局变量
	// 先找自己的作用域的变量 如果有 使用
	// 如果找不到 往外层找
	num = 100
	fmt.Println(num)

	name := "jay"
	fmt.Println(name)
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

// calc 函数作为参数
func calc(x, y int, op func(int, int) int) int {
	// 接收到 x y  然后调用op方法
	return op(x, y)
}

// do 函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

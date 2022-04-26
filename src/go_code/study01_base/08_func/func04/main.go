package main

import "fmt"

// 函数类型与变量
func main() {
	// 定义函数类型
	var c calculation               // 声明一个calculation类型的变量c
	c = add                         // 把add赋值给c
	fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	fmt.Println(c(1, 2))            // 像调用add一样调用c
}

// 定义方法类型
// type: 关键字
// calculation 类型名
// func(int, int) 参数
// int 返回
type calculation func(int, int) int

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

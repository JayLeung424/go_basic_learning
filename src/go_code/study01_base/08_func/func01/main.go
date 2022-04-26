package main

import "fmt"

// 函数

func main() {
	// 函数的调用
	function1()
	function2("jay")
	function3("a", "b")
	function4("xx", "x", "xxx")
	x, y := function5(10, 5)
	fmt.Println(x, y)
	function6(10, 5)
}

// 定义无参数无返回值的函数
func function1() {
	fmt.Println("hello world")
}

// 定义有接收无返回的函数
func function2(name string) {
	fmt.Println("hello world", name)
}

// 两个参数相同时候 可以缩写
func function3(name, age string) {
	fmt.Println("hello world", name, age)
}

// 可变参数  x 是切片类型
func function4(x ...string) {
	for _, s := range x {
		fmt.Println(s)
	}
}

// 定义有参数又返回的函数
func function5(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

// 声明好返回值
func function6(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub // 可以省略!
}

// defer: 延迟调用
// 将延迟处理的语句按defer定义的逆序进行执行
func deferFunc() {
	fmt.Println("start ...")
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end ...")
	// start ...
	// end ...
	// 3
	// 2
	// 1
}

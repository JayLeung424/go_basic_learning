package main

import "fmt"

// 指针
// & 取地址
// * 根据地址取值
func main() {

	a := 10
	b := &a
	c := *b
	fmt.Println(a) // 10
	fmt.Println(b) // 地址0xc00001a0b0
	fmt.Println(c) // 找到 "地址0xc00001a0b0" 里面存储的值

	x := 1
	modify1(x)
	fmt.Println(x) // 1 没有修改
	modify2(&x)
	fmt.Println(x) // 200 把地址修改了

	// var n *int // 声明一个地址 没有初始化 此时=nil
	// n = 100 // 会报错
	//
	// var b = map[string]int // 声明map要用make函数

	// new 函数 声明一个类型的指针 对应的值是默认值
	n := new(int)
	*n = 10 // n的地址= 10
}

func modify1(x int) {
	x = 100
}

// * 标记- 要地址！
func modify2(y *int) {
	*y = 200
}

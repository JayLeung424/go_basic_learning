package main

import "fmt"

// 自定义类型 和 类型别名示例

// 1、自定义类型

// MyInt 基于int类型的自定义类型
type MyInt int

// 2、类型别名

// MyNewInt 类型别名 实际上还是MyInt  只是MyInt的另一个名字
type MyNewInt = MyInt

func main() {
	var i MyInt
	fmt.Printf("type:%T value:%v\n", i, i) //type:main.MyInt value:0
	var j MyNewInt
	fmt.Printf("type:%T value:%v\n", j, j) // type:main.MyInt value:0

}

package main

import "fmt"

// for 循环
func main() {
	// 基础写法
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}
	// 省略初始语句 必须保留初始语句的分号
	j := 0
	for ; j < 10; j++ {
		fmt.Println("j:", j)
	}

	// 省略初始和结束
	for j < 10 {
		fmt.Println("j:", j)
		j++
	}
	// 无限循环
	// for{
	//	 fmt.Println("hello world")
	// }

	// break & continue
	for i := 0; i < 10; i++ {
		if i == 7 {
			// 结束循环
			break
		}
		if i == 4 {
			// 跳出本次循环  进入下次循环
			continue
		}
		fmt.Println(i)
	}

}

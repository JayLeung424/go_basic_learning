package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() { // 开启一个主goroutine去执行main函数

	wg.Add(1000) // 计数牌 登记记录一共有多少个goroutine
	for i := 0; i < 1000; i++ {
		// 闭包  线程安全问题
		// go func() {
		// 	fmt.Println("hello ", i) // i要从外面找 所以会重复
		// 	wg.Done()
		// }()
		// 解决方法  i 从内部找
		go func(i int) {
			fmt.Println("hello ", i) // i要从外面找 所以会重复
			wg.Done()
		}(i)
	}
	// hello  1000
	// hello  1000
	// hello  1000
	// hello  1000
	// hello  1000
	// 大量的goroutine
	wg.Wait()
}

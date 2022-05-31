package main

import (
	"fmt"
	"sync"
	"time"
)

/**
context_demo
*/

var wg sync.WaitGroup

// 全局变量的方式
var exit bool = false

// work func
func work() {
	defer wg.Done() // 计数器-1
	// 死循环 如果不停止的话 wg.Done() 永远都不能执行
	for {
		fmt.Println("worker ...")
		time.Sleep(time.Second)

		// 方式1、使用全局变量 - 如果exit=true 就break
		if exit {
			break
		}
	}
	// 如何接受外部命令实现退出
}

func main() {
	wg.Add(1) // 计数牌
	go work() // 开启了一个独立的goroutine
	// 如何优雅的实现结束goroutine
	// 1、全局变量的方式
	time.Sleep(time.Second * 5)
	exit = true // 修改全局变量的方式 通知子goroutine退出

	wg.Wait() // 等待结束
	fmt.Println("over")

}

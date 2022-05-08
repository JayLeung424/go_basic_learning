package main

import (
	"fmt"
	"sync"
)

/**
并发 与 并行
	并发: 同一时间段内执行多个任务 (你用微信同时和两个女朋友在聊天)
	并行: 同一时刻执行多个任务 (你和你的朋友同时在和你的女朋友聊天)
*/

var wg sync.WaitGroup

// goroutine
func hello() {
	fmt.Println("hello world!")
	wg.Done() // 计数器-1  小弟通知wg计数器-1
}

func main() { // 开启一个主goroutine去执行main函数

	wg.Add(2)  // 计数牌 登记记录一共有多少个goroutine
	go hello() // 开启了一个独立的goroutine 执行hello这个函数
	go hello() // 开启了一个独立的goroutine 执行hello这个函数
	fmt.Println("hello main!")
	// 方式1 不可取
	// time.Sleep(time.Second) // 睡眠1s 然后等待时间

	// 方式2 sync.WaitGroup
	wg.Wait() // 等所有的小弟干完以后 才结束
}

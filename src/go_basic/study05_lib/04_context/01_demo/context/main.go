package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/**
context_demo
*/

var wg sync.WaitGroup

// work func
func work(ctx context.Context) {
	defer wg.Done() // 计数器-1
	go work2(ctx)
	// 死循环 如果不停止的话 wg.Done() 永远都不能执行
LABEL:
	for {
		fmt.Println("worker ...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:

		}
	}
	// 如何接受外部命令实现退出
}

// work2 func
func work2(ctx context.Context) {
	defer wg.Done() // 计数器-1
	// 死循环 如果不停止的话 wg.Done() 永远都不能执行
LABEL:
	for {
		fmt.Println("worker2 ...")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break LABEL
		default:

		}
	}
	// 如何接受外部命令实现退出
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(2)    // 计数牌
	go work(ctx) // 开启了一个独立的goroutine
	// 如何优雅的实现结束goroutine
	time.Sleep(time.Second * 5)

	cancel() // 调用cancle函数 告诉子goroutine 退出！

	wg.Wait() // 等待结束
	fmt.Println("over")

}

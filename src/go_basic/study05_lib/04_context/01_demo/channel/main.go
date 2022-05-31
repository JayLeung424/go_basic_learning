package main

import (
	"fmt"
	"sync"
	"time"
)

/**
使用 channle的方式
*/

var wg sync.WaitGroup

// work func
func work(ch <-chan bool) { // 单向通道 只能取值
	defer wg.Done() // 计数器-1
LABEL:
	// 死循环 如果不停止的话 wg.Done() 永远都不能执行
	for {
		select {
		case <-ch: // 如果能从ch取出来值  就退出
			// break // 只能退出select
			break LABEL // 退出整个for循环 （全局的意思）
		default:
			fmt.Println("worker ...")
			time.Sleep(time.Second)
		}
	}
	// 如何接受外部命令实现退出
}

func main() {
	exitChan := make(chan bool)

	wg.Add(1)         // 计数牌
	go work(exitChan) // 开启了一个独立的goroutine
	// 如何优雅的实现结束goroutine
	time.Sleep(time.Second * 5)
	exitChan <- true //
	wg.Wait()        // 等待结束
	fmt.Println("over")

}

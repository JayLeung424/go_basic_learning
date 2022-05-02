package main

import "fmt"

// channel
// var 变量 chan 元素类型
func main() {

	// 声明 - 引用类型 必须初始化才可以使用
	var ch1 chan int    // 生命一个传递int的通道
	var ch2 chan string // 生命一个传递string的通道
	var ch3 chan bool   // 生命一个传递bool的通道

	// 使用make函数进行初始化 - 记得定义长度
	ch1 = make(chan int, 1)
	ch2 = make(chan string, 1)
	ch3 = make(chan bool, 1)

	// 发送
	ch1 <- 10
	ch2 <- "jay"
	ch3 <- true

	// 接收
	c1 := <-ch1
	c2 := <-ch2
	c3 := <-ch3

	// 关闭通道 - 所有消息都发送出去以后 被垃圾回收器回收
	close(ch1)
	close(ch2)
	close(ch3)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
}

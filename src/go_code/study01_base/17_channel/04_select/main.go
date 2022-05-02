package main

import "fmt"

// select
func main() {
	ch := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // 如果可以从ch中取到值
			fmt.Println(x)
		case ch <- i: // 如果把值塞入ch中
			fmt.Println("在ch塞入了", i)
		default:
			fmt.Println("啥也不干")
		}
	}
}

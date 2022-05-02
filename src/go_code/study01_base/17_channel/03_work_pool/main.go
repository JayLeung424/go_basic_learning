package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker:%d start job :%d\n", id, job)
		results <- job * job
		time.Sleep(time.Second)
		fmt.Printf("worker:%d start job :%d\n", id, job)
	}
}

// work pool
func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- 1
	}
	close(jobs)
	// 输出结果
	for i := 0; i < 3; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}

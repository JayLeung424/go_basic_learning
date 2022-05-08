package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁
var (
	x     int64
	wg    sync.WaitGroup
	mutex sync.Mutex // 互斥锁

	// 当读远远大于写的时候 推荐使用读写互斥锁
	rwMutex sync.RWMutex
)

// writeWithLock 使用互斥锁的写操作
func writeWithLock() {
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

// readWithLock 使用互斥锁的读操作
func readWithLock() {
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

// writeWithLock 使用读写互斥锁的写操作
func writeWithRWLock() {
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

// readWithRWLock 使用读写互斥锁的读操作
func readWithRWLock() {
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

func main() {
	start := time.Now()
	// 读1000次
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go readWithRWLock()
		// go readWithLock()
	}
	// 写10次
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writeWithRWLock()
		// go writeWithLock()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}

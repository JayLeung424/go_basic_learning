package main

import (
	"fmt"
	"sync"
)

// sync.map 并发安全的map

var (
	wg sync.WaitGroup
	// 使用并发安全的map
	sm sync.Map
)
var m = make(map[int]int, 100)

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

func main() {
	// 原生的map 不支持并发插入
	// for i := 0; i < 30; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		set(i, i+10)
	// 		fmt.Printf("key:%v, value:%v\n", i, get(i))
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait()

	for i := 0; i < 30; i++ {
		// store 存储
		sm.Store(i, i+100)
		// Load取
		value, _ := sm.Load(i)
		fmt.Printf("key:%v, value:%v\n", i, value)
		// func (m *Map) Store(key, value interface{})	存储key-value数据
		// func (m *Map) Load(key interface{}) (value interface{}, ok bool)	查询key对应的value
		// func (m *Map) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)	查询或存储key对应的value
		// func (m *Map) LoadAndDelete(key interface{}) (value interface{}, loaded bool)	查询并删除key
		// func (m *Map) Delete(key interface{})	删除key
		// func (m *Map) Range(f func(key, value interface{}) bool)	对map中的每个key-value依次调用f
	}
}

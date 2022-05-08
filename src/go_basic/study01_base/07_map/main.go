package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map
func main() {
	// 声明一个map key:string value:int
	var m map[string]int
	fmt.Println(m == nil) // true 只有声明map类型 没有初始化 a就是nil

	// map的初始化
	m = make(map[string]int)

	// 赋值
	m["jay"] = 7
	m["go"] = 10
	fmt.Println(m)
	fmt.Printf("type:%T\n", m)

	// 声明并初始化
	b := map[int]bool{
		1: true,
		0: false,
	}
	fmt.Println(b)
	fmt.Printf("type:%T\n", b)

	// 判断某个key是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["tian"] = 100
	scoreMap["jay"] = 200

	// 判断 alen 是否存在scoremap中
	value, ok := scoreMap["alen"]
	if ok {
		fmt.Println("Alen在scoreMap中", value)
	} else {
		fmt.Println("查无此人!")
	}

	// 删除某个key
	delete(scoreMap, "tian")
	fmt.Println(scoreMap)

	// for range
	// map是无序的 键值对和添加的顺序无关!
	for k, v := range scoreMap {
		fmt.Println("k:", k, "value:", v)
	}
	for _, v := range scoreMap {
		fmt.Println("value:", v)
	}

	// 按照某个固定的顺序遍历map
	var score = make(map[string]int, 100)

	// 添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0～99随机整数
		score[key] = value
	}
	// key 按照从小到大的顺序
	// 1、取出所有的key 存放到切片中
	keys := make([]string, 0, 100) // 目前0个元素 能放100个元素
	for k, _ := range score {
		keys = append(keys, k)
	}
	// 2、进行排序
	sort.Strings(keys)
	// 3、根据排序后的key 取值
	for _, key := range keys {
		value := score[key]
		fmt.Println("k:", key, "value: ", value)
	}

	// 元素类型为map的切片
	var mapSlice = make([]map[string]int, 8, 8) // 只是完成了切片的初始化
	// 现在里面都是nil  无法赋值
	// mapSlice[0]["jay"] = 10 // assignment to entry in nil map
	// fmt.Println(mapSlice)

	// 完成对内部map的初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["jay"] = 10
	fmt.Println(mapSlice)

	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	key := "中国"
	v, ok := sliceMap[key]
	if !ok {
		v = make([]string, 0, 2)
	}
	v = append(v, "北京", "上海")
	sliceMap[key] = v
	fmt.Println(sliceMap)

}

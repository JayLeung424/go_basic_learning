package main

import "fmt"

// 切片函数
// 切片一定要初始化再使用
func main() {
	// 切片的扩容
	var a []int // 此时没有申请内存
	//a[0] = 100
	fmt.Println(a) //index out of range [0] with length 0

	// append 动态扩容
	//a = append(a,10) //[10]
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", a, len(a), cap(a), a)
	}
	// cap 是成2倍的增加
	//[0]  len:1  cap:1  ptr:0xc0000b2010
	//[0 1]  len:2  cap:2  ptr:0xc0000b2030
	//[0 1 2]  len:3  cap:4  ptr:0xc0000b6020
	//[0 1 2 3]  len:4  cap:4  ptr:0xc0000b6020
	//[0 1 2 3 4]  len:5  cap:8  ptr:0xc0000aa080
	//[0 1 2 3 4 5]  len:6  cap:8  ptr:0xc0000aa080
	//[0 1 2 3 4 5 6]  len:7  cap:8  ptr:0xc0000aa080
	//[0 1 2 3 4 5 6 7]  len:8  cap:8  ptr:0xc0000aa080
	//[0 1 2 3 4 5 6 7 8]  len:9  cap:16  ptr:0xc0000b8000
	//[0 1 2 3 4 5 6 7 8 9]  len:10  cap:16  ptr:0xc0000b8000

	// 一次追加多个
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
	b := []int{1, 2, 3, 4}
	b = append(a, b...)
	fmt.Println(b)

	// copy()复制切片
	// 只是拷贝值 不是指向同一个地址 浅拷贝
	co1 := []int{1, 2, 3, 4, 5}
	co2 := make([]int, 5, 5)
	copy(co2, co1)   //使用copy()函数将切片a中的元素复制到切片c
	fmt.Println(co1) //[1 2 3 4 5]
	fmt.Println(co2) //[1 2 3 4 5]
	co2[0] = 1000
	fmt.Println(co1) //[1 2 3 4 5]
	fmt.Println(co2) //[1000 2 3 4 5]

	// 从切片中删除元素
	d := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	d = append(d[:2], d[3:]...) // 在切片[0:2]只有追加 [3:]以后的切片
	fmt.Println(d)              //[30 31 33 34 35 36 37]

}

package main

import "fmt"

// 切片
// 切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
func main() {
	// 基于数组得到切片
	a := [5]int{55, 66, 77, 88, 99}
	b := a[1:4]
	fmt.Printf("%T\n", b)
	fmt.Println(b)
	// 切片再次切片
	c := b[0:]
	fmt.Printf("%T\n", b)
	fmt.Println(c)

	// make 函数构造切片
	d := make([]int, 5, 10) // [0 0 0 0 0]  容量为5 长度为10
	fmt.Println(d)

	// 通过len函数 获取长度
	fmt.Println("长度:", len(d))
	fmt.Println("容量:", cap(d))

	var s1 []int         // 声明int切片         len(s1)=0;cap(s1)=0;s1==nil
	s2 := []int{}        // 声明int切片且初始化   len(s2)=0;cap(s2)=0;s2!=nil
	s3 := make([]int, 0) //                    len(s3)=0;cap(s3)=0;s3!=nil

	if s1 == nil {
		fmt.Println("s1是一个nil")
	}
	if s2 == nil {
		fmt.Println("s2是一个nil")
	}
	if s3 == nil {
		fmt.Println("s3是一个nil")
	}
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	// 切片赋值拷贝
	i := make([]int, 5, 10)
	j := i
	j[0] = 100
	// 指向同一个地址
	fmt.Println(i) // [100 0 0 0 0]
	fmt.Println(j) // [100 0 0 0 0]

	// 遍历
	for i := 0; i < len(j); i++ {
		fmt.Println(j[i])
	}
	for _, value := range j {
		fmt.Println(value)
	}

}

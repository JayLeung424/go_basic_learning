package main

import "fmt"

// 数组
// 数组是同一种数据类型元素的集合。
// 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
func main() {
	var a [3]int
	var b [4]int
	var c [5]string
	fmt.Println(a) // [0 0 0]
	fmt.Println(b) // [0 0 0 0]
	fmt.Println(c) // [    ]

	// a = b 长度不同 不能直接等 数组越界

	// 初始化
	// 1、定义的时候使用初始值列表的方式初始化
	var cityArray = [4]string{"北京", "上海", "杭州", "广州"}
	fmt.Println(cityArray)

	// 2、编译器自己推导数组长度
	var boolArray = [...]bool{true, false, true, false}
	fmt.Println(boolArray)
	fmt.Println(boolArray[0])

	// 3、使用索引值方式初始化
	var langArray = [...]string{1: "go", 3: "python", 7: "java"}
	fmt.Println(langArray)

	// 数组遍历
	// for循环
	for i := 0; i < len(cityArray); i++ {
		fmt.Println(cityArray[i])
	}
	// for range
	for index, value := range cityArray {
		fmt.Println("index :", index, " value:", value)
	}
	for _, value := range cityArray {
		fmt.Println("value:", value)
	}

	// 二维数组
	//school := [...][2]string 只有外层可以用... 让编译器推导
	school := [3][2]string{
		{"语文", "数学"},
		{"英语", "物理"},
		{"音乐", "体育"},
	}
	fmt.Println(school)
	for _, strings := range school {
		for _, i2 := range strings {
			fmt.Println(i2)
		}
	}
}

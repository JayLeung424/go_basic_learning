package main

import (
	"fmt"
	"math"
)

// 基本数据类型

func main() {
	// 十进制打印为二进制
	n := 23
	fmt.Printf("十进制:%d\n", n)  // 十进制 23
	fmt.Printf("二进制:%b\n", n)  // 二进制 10111
	fmt.Printf("八进制:%o\n", n)  // 八进制 27
	fmt.Printf("十六进制:%x\n", n) // 十六进制 17

	// unit8 范围: 0~255
	var age uint8
	// age = 256;
	fmt.Println("unit:", age)

	// 浮点数
	fmt.Println("MaxFloat32:", math.MaxFloat32) // 3.4028234663852886e+38
	fmt.Println("MaxFloat64:", math.MaxFloat64) // 1.7976931348623157e+308

	// 布尔值
	var b bool
	fmt.Println("bool默认值:", b) // false
	b = true
	fmt.Println("bool值:", b)

	// 字符串
	s1 := "hello world！"
	s2 := "你好 世界！"
	fmt.Println("英文字符串:", s1)
	fmt.Println("中文字符串:", s2)
	// 多行字符串
	s3 := `
		第一行
		第二行
		第三行
	`
	fmt.Println("多行字符串:", s3)

	// 转译符号
	//  r	回车符（返回行首）
	//  \n	换行符（直接跳到下一行的同列位置）
	//  \t	制表符
	//  \'	单引号
	//  \"	双引号
	//	\\	反斜杠
	// windows下打印自己的地址
	fmt.Println("str := \\golang\\golang_learning")

	// byte 和 rune
	// uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	// rune类型，代表一个 UTF-8字符。
	var c1 byte = 'c'
	var c2 = 'c'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T  c2:%T\n", c1, c2) // c1:uint8  c2:int32

	s := "hello梁杰"
	// %c ->helloæ¢ æ °%
	//
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}
	fmt.Println()

	// range ->hello梁杰
	for _, r := range s {
		fmt.Printf("%c", r)
	}
}

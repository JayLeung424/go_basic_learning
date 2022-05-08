package main

import "fmt"

// Go语言中的运算符
func main() {

	// 1、算术运算符
	a, b := 10, 20
	fmt.Println("加:", b+a)
	fmt.Println("减:", b-a)
	fmt.Println("乘:", b*a)
	fmt.Println("除:", b/a)
	fmt.Println("求余:", b%a)

	// 自增 自减
	a++
	a--

	// 2、关系运算符
	fmt.Println("大于:", b > a)
	fmt.Println("小于:", b < a)
	fmt.Println("等于:", b == a)
	fmt.Println("不等:", b != a)

	// 3、逻辑运算符
	// && 表示and
	fmt.Println("&& :", a == 10 && b == 20)
	// || 表示or
	fmt.Println("|| :", 10 > 20 || b == 20)
	// ! 取反
	fmt.Println("! :", !(10 > 20 || b == 20))

	// 4、位运算符
	i := 1 // i: 001
	j := 5 // j: 101
	// &	参与运算的两数各对应的二进位相与。（两位均为1才为1）
	fmt.Println(i & j) // 001
	// |	参与运算的两数各对应的二进位相或。（两位有一个为1就为1）
	fmt.Println(i & j) // 101
	// ^	参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1）
	fmt.Println(i ^ j) // 100
	// <<	左移n位就是乘以2的n次方。 “a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。
	fmt.Println(1 << 2) // 001 -> 100
	// >>	右移n位就是除以2的n次方。“a>>b”是把a的各二进位全部右移b位。
	fmt.Println(4 >> 2) // 100 -> 001

	// 5、赋值运算符
	a += 1
	fmt.Println(a) // 11

}

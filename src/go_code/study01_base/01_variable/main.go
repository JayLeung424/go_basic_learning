package main

import "fmt"

// 变量与常量
// 函数外面 只能放标识符(变量/常量/函数/类型)的声明

// 程序的入口函数
func main() {
	fmt.Println("Hello world！") // 换行打印

	// 声明变量 同时赋值
	var s1 = "s111"
	fmt.Println("s1:", s1)
	// 类型推导, 根据值 判断变量类型
	var s2 = 20
	fmt.Println("s2:", s2)
	// 简短变量声明(不能在函数外声明)
	s3 := "哈哈哈"
	fmt.Println("s3:", s3)
	// 匿名变量(当有些数据必须用变量接受，但是又不使用的话，就可以使用匿名变量)
	x, _ := foo()
	_, y := foo()
	a, b := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Printf("a=%d,b=%s", a, b)

	// iota
	fmt.Println("d1", d1)  // 1
	fmt.Println("d2", d2)  // 2
	fmt.Println("d3:", d3) // 2
	fmt.Println("d4:", d4) // 3

	fmt.Println("a1:", a1) // 0
	fmt.Println("a2:", a2) // 100
	fmt.Println("a3:", a3) // 100
	fmt.Println("a4:", a4) // 4
	fmt.Println("a5:", a5) // 5

	// 批量声明变量
	firstName2 = "t"
	lastName2 = "xd"
	age2 = 12
	// Go语言的变量声明后必须使用。 不使用 编译不过去
	fmt.Print()                                                                // 打印内容
	fmt.Printf("firstName:%s, lastName:%s,年龄:%d", firstName2, lastName2, age2) // Printf可以使用占位符 %d 十进制

}

// Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。 并且Go语言的变量声明后必须使用。
// var firstName1 string
// var lastName1 string
// var age1 int

// 批量声明变量
var (
	firstName2 string
	lastName2  string
	age2       int
)

// 常量 一旦定义  无法修改
const pi = 3.1415926

// 批量声明常量
const (
	STATUS_OK    = 200
	NOT_FOUND    = 404
	STATUS_ERROR = 500
)

// 如果无值, 和上面一样
const (
	n1 = 100
	n2
	n3
)

// 常量计数器 iota是go语言的常量计数器，只能在常量的表达式中使用。
// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。
const (
	// 实现枚举  一个个累加
	a1 = iota // 0
	a2 = 100  // 100
	a3        // 100
	_         // 100
	a4 = iota // 4
	a5        // 5
)

const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

func foo() (int, string) {
	return 10, "Jay"
}

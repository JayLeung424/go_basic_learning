package main

import "fmt"

// 匿名字段
// 结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

// Person 结构体Person类型
type Person struct {
	string
	int
	// int  相同字段只能有一个
}

func main() {
	p1 := Person{
		"Jay",
		18,
	}
	fmt.Printf("%#v\n", p1)        //main.Person{string:"北京", int:18}
	fmt.Println(p1.string, p1.int) //北京 18
}

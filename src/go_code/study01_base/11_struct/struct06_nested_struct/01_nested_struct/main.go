package main

import "fmt"

// 嵌套结构体

// Address 地址的属性
type Address struct {
	Province string
	City     string
}

// Person 人的属性
type Person struct {
	Name   string
	Gender string
	Age    int
	// 嵌套其他的结构体
	Address
}

func main() {
	p := Person{
		Name:   "jay",
		Gender: "男",
		Age:    22,
		Address: Address{
			Province: "SH",
			City:     "SH",
		},
	}
	fmt.Println(p)
	fmt.Println(p.Address.Province) // 通过嵌套的结构体字段 拿到对应字段
	fmt.Println(p.Province)         // 直接访问对应的字段

}

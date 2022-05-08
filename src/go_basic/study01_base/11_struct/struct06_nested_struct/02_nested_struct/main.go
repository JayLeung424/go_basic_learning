package main

import "fmt"

// 匿名结构体字段冲突如何解决？

// Address 地址的属性
type Address struct {
	Province   string
	City       string
	UpdateTime string
}

// Email 邮件的属性
type Email struct {
	Addr       string
	UpdateTime string
}

// Person 人的属性
type Person struct {
	Name   string
	Gender string
	Age    int
	// 嵌套其他的结构体
	Address
	Email
}

func main() {
	p := Person{
		Name:   "jay",
		Gender: "男",
		Age:    22,
		Address: Address{
			Province:   "SH",
			City:       "SH",
			UpdateTime: "Address UpdateTime",
		},
		Email: Email{
			Addr:       "AAA@gmail",
			UpdateTime: "Email UpdateTime",
		},
	}
	fmt.Println(p)
	fmt.Println(p.Address.Province) // 通过嵌套的结构体字段 拿到对应字段
	fmt.Println(p.Province)         // 直接访问对应的字段
	// fmt.Println(p.UpdateTime)         // UpdateTime 存在两个 不能直接访问
	fmt.Println(p.Address.UpdateTime) // 通过嵌套的结构体字段 拿到对应字段
	fmt.Println(p.Email.UpdateTime)   // 通过嵌套的结构体字段 拿到对应字段

}

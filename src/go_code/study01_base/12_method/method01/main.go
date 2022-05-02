package main

import "fmt"

// 方法的定义
// 方法和函数的区别
// 方法有个接受者 只有接受者的类型才可以调用
// 函数是谁都可以调用

// Person 是一个结构体
type Person struct {
	name string
	age  int8
}

// Student 是一个结构体
type Student struct {
	name string
	age  int8
}

// NewPerson 是Person的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream 是为 Person 定义方法 [(p Person) 接受者！]
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好go语言\n", p.name)
}

// SetAge : set age  接受者是指针！
func (p *Person) SetAge(age int8) {
	p.age = age
}

// SetAge2 : set age  接受者是值！
func (p Person) SetAge2(age int8) {
	// 这个时候只有p的age修改了  但是实际上调用的对象的age没有修改
	p.age = age
}

func main() {
	p1 := NewPerson("jay", 22)
	// p1 现在是地址
	fmt.Printf("%T\n", p1)
	// 标准方法使用
	(*p1).Dream()
	// 简单使用
	p1.Dream()

	s := Student{name: "jay"}
	// s.Dream()  // 没有该方法 因为Dream属于Person类型的方法
	fmt.Println(s)

	// 接受者是指针的使用
	fmt.Println(p1.age) // 22
	p1.SetAge(11)
	fmt.Println(p1.age) // 11
	p1.SetAge2(30)
	fmt.Println(p1.age) // 11  因为函数传参是值拷贝

}

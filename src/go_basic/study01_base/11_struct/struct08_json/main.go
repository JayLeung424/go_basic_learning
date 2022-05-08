package main

import (
	"encoding/json"
	"fmt"
)

// 字段的可见性 与 Json序列化

// Student 如果go语言包中定义的标识符是首字母大写的  那么就是对外可见的
// 通常小写 需要的话再大写
type Student struct {
	ID   int // 对外可见的 通过反射可以拿到
	Name string
}

// newStudent 构造函数
func newStudent(id int, name string) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}

// class 班级
type class struct {
	// Title   string
	// title   string
	title   string `Title` // ``标记为返回的名字
	Student []Student
}

func main() {
	c1 := class{
		// Title:   "火箭101",
		title:   "火箭101",
		Student: make([]Student, 0, 20),
	}
	// 往班级中添加学生信息
	for i := 0; i < 10; i++ {
		// 造十个学生信息
		stu := newStudent(i, fmt.Sprintf("stu%02d", i))
		c1.Student = append(c1.Student, stu)
	}
	fmt.Printf("%#v\n", c1)

	// Json的序列化 和反序列化
	// Json序列化：Go语言中的数据 -> json格式的字符串
	data, err := json.Marshal(c1)
	// 发现没有title  因为title是小写 不能被别的包访问到
	// {"Student":[{"ID":0,"Name":"stu00"},{"ID":1,"Name":"stu01"},{"ID":2,"Name":"stu02"},{"ID":3,"Name":"stu03"},{"ID":4,"Name":"stu04"},{"ID":5,"Name":"stu05"},{"ID":6,"Name":"stu06"},{"ID":7,"Name":"stu07"},{"ID":8,"Name":"stu08"},{"ID":9,"Name":"stu09"}]}
	// main.class{title:"", Student:[]main.Student{main.Student{ID:0, Name:"stu00"}, main.Student{ID:1, Name:"stu01"}, main.Student{ID:2, Name:"stu02"}, main.Student{ID:3, Name:"stu03"}, main.Student{ID:4, Name:"stu04"}, main.Student{ID:5, Name:"stu05"}, main.Student{ID:6, Name:"stu06"}, main.Student{ID:7, Name:"stu07"}, main.Student{ID:8, Name:"stu08"}, main.Student{ID:9, Name:"stu09"}}}
	if err != nil {
		fmt.Println("Json序列化失败了,失败原因:", err)
	} else {
		fmt.Printf("%T\n", data)
		fmt.Printf("%s\n", data)
	}

	// Json反序列化: Json格式的字符串 -> Go语言的数据
	jsonStr := data
	var c2 class
	// 用地址接收
	json.Unmarshal(jsonStr, &c2)
	fmt.Printf("%#v\n", c2)

}

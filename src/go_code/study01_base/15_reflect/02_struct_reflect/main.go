package main

import (
	"fmt"
	"reflect"
)

// 结构体反射

type student struct {
	Name  string `json:"name" ini:"n_name"`
	Score int    `json:"score" ini:"s_score"`
}

func (s student) Study() string {
	msg := "好好学习,天天向上"
	fmt.Println(msg)
	return msg
}
func (s student) Sleep() string {
	msg := "好好睡觉,快点长大"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	// 获取变量方法的数量
	fmt.Println(t.NumMethod())
	// 因为下面要拿到具体的方法取调用 所以使用值信息
	v := reflect.ValueOf(x)
	for i := 0; i < v.NumMethod(); i++ {
		// 方法名
		fmt.Printf("method name :%s\n", t.Method(i).Name)
		// 方法类型
		fmt.Printf("method type :%s\n", t.Method(i).Type)

		// 通过反射调用方法传递的参数必须是[] reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}

	// 通过方法名获取
	t.MethodByName("Sleep")              // (method, bool)
	methodObj := v.MethodByName("Sleep") // value
	fmt.Println(methodObj.Type())        // func() string
	fmt.Println(methodObj)               // 0x10851e0

}

func main() {
	stu1 := student{
		Name:  "jay",
		Score: 97,
	}

	// 通过反射获取结构体中的所有的字段
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v ,score:%v\n", t.Name(), t.Kind())

	// 遍历结构体变量的所有字段   t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 根据索引去拿字段
		fieldObj := t.Field(i)
		// fmt.Println(fieldObj)
		fmt.Printf("Name:%v, type:%v, tag:%v\n", fieldObj.Name, fieldObj.Type, fieldObj.Tag)
		fmt.Println(fieldObj.Tag.Get("json")) // score
		fmt.Println(fieldObj.Tag.Get("ini"))  // s_score%
		// Name 字段名  type:类型  tag 信息
		// Name:Name, type:string, tag:json:"name" ini:"n_name"
		// Name:Score, type:int, tag:json:"score" ini:"s_score"
	}

	// 根据名字取结构体的字段
	field, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("Name:%v, type:%v, tag:%v\n", field.Name, field.Type, field.Tag)
	}

	printMethod(stu1)
}

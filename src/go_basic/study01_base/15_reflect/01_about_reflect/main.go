package main

import (
	"fmt"
	"reflect"
)

// 反射
// reflectType 拿到参数的类型
func reflectType(x interface{}) {
	// 不知道调用这个函数的时候 会传递什么类型的变量
	// 方式1 类型断言 缺点:一个一个猜
	// 方式2 使用反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj)
	fmt.Printf("%T\n", obj)
	// type:自定义类型   kind:底层的类型
	fmt.Println("Name:", obj.Name(), ", Kind:", obj.Kind())
}

// reflectValue 拿到参数的值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T\n", v, v)
	k := v.Kind() // 拿到值对应的类型种类
	// 如何得到一个传入时候类型的变量
	switch k {
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Println(ret)
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// 用来根据指针取到对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Float32:
		v.Elem().SetFloat(100.123)
	}
}

type Cat struct {
}

func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 10
	reflectType(b)
	// 结构体类型
	var c = Cat{}
	reflectType(c) // Name:  Cat , Kind: struct
	var d []int
	reflectType(d) // Name:  , Kind: slice
	var e []string
	reflectType(e) // Name:  , Kind: slice

	// reflect.ValueOf(x)
	reflectValue(a) // 1.234,reflect.Value
	reflectValue(b) // 10,reflect.Value
	reflectValue(c)
	reflectValue(d)
	reflectValue(e)

	// v.Elem().SetFloat
	fmt.Println(a)
}

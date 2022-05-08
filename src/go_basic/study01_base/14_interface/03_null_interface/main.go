package main

import "fmt"

/**
- 空接口 -
	接口中没有定义任何方法时候  该接口就是一个空接口
	空接口的特点： 任意类型都实现了空接口 -> 空接口的变量可以存储任意值
	空接口一般不需要提前定义
*/

type xxx interface {
}

/**
空接口的应用:
	1、空接口类型作为函数的参数 例如 ：fmt.Println()
	2、空接口作为map的value
*/

func main() {
	var x interface{} // 定义一个空接口变量 x
	x = "hello"
	x = 100
	x = false
	fmt.Println(x)

	var m = make(map[string]interface{}, 36)
	m["name"] = "josiah"
	m["age"] = 20
	m["hobby"] = []string{"篮球", "足球"}
	fmt.Println(m)

	// 类型断言
	ret, ok := x.(string)
	if !ok {
		fmt.Println("不是string类型")
	} else {
		fmt.Println(ret)
	}

	// 使用switch进行断言
	switch v := x.(type) {
	case string:
		fmt.Printf("%v是string类型", v)
		break
	case bool:
		fmt.Printf("%v是bool类型", v)
		break
	case int:
		fmt.Printf("%v是int类型", v)
		break
	default:
		fmt.Printf("%v不知道是什么类型", v)
		break
	}
}

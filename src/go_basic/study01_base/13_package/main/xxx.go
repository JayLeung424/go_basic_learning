package main

import (
	// AliasName 引用自己的包  当代码都在$GOPATH目录下的话
	// 我们导入包的路径要从$GOPATH/src后面继续写
	AliasName "../calc" // 别名
	// _ 匿名导入包
	_ "../snow"

	"fmt"
	// Go语言中不允许导入包而不使用！！
	// Go语言不允许相互引用包 A引用B B引用A
)

func main() {
	// 记得首字母大写
	fmt.Println(AliasName.Name)
	ret := AliasName.Add(10, 20)
	fmt.Println(ret)
}

// init函数在包导入的时候 自动执行
// init函数没有参数 也没有返回值
func init() {
	fmt.Println("main.init")
}

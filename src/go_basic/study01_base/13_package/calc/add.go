package calc

import (
	snow "../snow"
	"fmt"
)

// 标识符首字母大写 表示对外可见
// 通常不对外的标识符都是要首字母小写
// 当某些变量函数被其他的包访问的时候需要加上注释

// Name 定义一个测试的全局变量
var Name = "Jay"

// Add 定义一个测试的全局函数
func Add(x, y int) int {
	// 不同包下需要引用包在调用
	snow.Snow()
	// 同一个包的不同文件可以直接调用
	Sub(x, y)
	return x + y
}

// init函数在包导入的时候 自动执行
// init函数没有参数 也没有返回值
func init() {
	fmt.Println("init 函数run...")
}

package main

import (
	"fmt"
	"strings"
)

// 字符串中的常见操作
func main() {
	str := "hello jay!"
	// len(str)	求长度
	fmt.Println("长度:", len(str)) // 10

	// +或fmt.Sprintf	拼接字符串
	fmt.Println("拼接字符串:", str+"my bro")        // hello jay!my bro
	fmt.Printf("拼接字符串:%s%s\n", str, "my bro!") // hello jay!my bro!

	// strings.Split(v1,v2)	 分割 v1:切割字符 v2:以v2开始切割
	fmt.Println("分割:", strings.Split(str, "l")) // 分割: [he  o jay!]

	// strings.contains	判断是否包含
	fmt.Println("包含:", strings.Contains(str, "jay")) // true

	// strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
	fmt.Println("前缀:", strings.HasPrefix(str, "jay!")) //false
	fmt.Println("后缀:", strings.HasSuffix(str, "jay!")) //true

	// strings.Index(),strings.LastIndex()	子串出现的位置
	fmt.Println("出现的位置:", strings.Index(str, "l"))     // 2
	fmt.Println("最后一次出现的位置:", strings.Index(str, "l")) // 2

	// strings.Join(a[]string, sep string)	join操作
	str2 := []string{"how", "do", "you", "do"}
	fmt.Println(str2)
	fmt.Println("拼接:", strings.Join(str2, "+")) // 拼接: how+do+you+do
}

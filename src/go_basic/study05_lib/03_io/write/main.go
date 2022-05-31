package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

/**
文件操作 - 写
func OpenFile(name string, flag int, perm FileMode) (*File, error) {
	...
}
*/
func main() {
	writeFileByIoutil()
}

// writeFileByIoutil 使用bufio写文件
func writeFileByIoutil() {
	str := "hello jj"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed,err:", err)
		return
	}
}

// writeFileByBufio 使用bufio写文件
func writeFileByBufio() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file faild,err:%v\n", fileObj)
		return
	}
	writer := bufio.NewWriter(fileObj)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello 林俊杰\n") // 将数据都先写入缓存
	}
	// 将缓存中的内容写入文件
	writer.Flush()
}

// writeFile io包写文件
func writeFile() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("open file faild,err:%v\n", fileObj)
		return
	}

	str := "河南林俊杰\n"
	fileObj.Write([]byte(str))       // []byte类型
	fileObj.WriteString("hello 林俊杰") // string 类型
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
文件操作 - 读取
*/
func main() {
	// readAllFromFile()
	// readByBufio()
	readByIoutil()
}

// readByIoutil ioutil读取文件
func readByIoutil() {
	// 文件不要太大 容易内存溢出
	file, err := ioutil.ReadFile("./xx.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Print(string(file))
}

// readByBufio bufio读取文件
func readByBufio() {

	// 打开文件
	fileObj, err := os.Open("./xx.txt") // 相对路径
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", nil)
		return
	}
	// 关闭文件
	defer fileObj.Close()

	// 读取文件内容 bufio
	reader := bufio.NewReader(fileObj)
	// // readLine 读取每一行
	// line, _, err := reader.ReadLine()
	// for i := range line {
	// 	fmt.Println(i)
	//
	// }
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}
}

// readAllFromFile 无法估量字节长度
func readAllFromFile() {

	// 打开文件
	fileObj, err := os.Open("./xx.txt") // 相对路径
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", nil)
		return
	}
	// 关闭文件
	defer fileObj.Close()

	// 读取文件内容
	for {

		var tmp = make([]byte, 10)
		read, err := fileObj.Read(tmp)
		if err == io.EOF { // EOF = end of file
			// 把当前读了多少个字节的数据打印出来，然后瑞出
			fmt.Println(string(tmp[:read]))
		}
		if err != nil {
			fmt.Printf("read file failed,err:%v\n", nil)
			return
		}
		fmt.Printf("read  %d bytes from file.\n", read)
		fmt.Println(string(tmp))

	}
}

// readFromFile 读取文件 自定义好最大字节数
func readFromFile() {

	// 打开文件
	fileObj, err := os.Open("./xx.txt") // 相对路径
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", nil)
		return
	}
	// 关闭文件
	defer fileObj.Close()

	// 读取文件内容
	var tmp = make([]byte, 128)
	read, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", nil)
		return
	}
	fmt.Printf("read  %d bytes from file.\n", read)
	fmt.Println(string(tmp))
}

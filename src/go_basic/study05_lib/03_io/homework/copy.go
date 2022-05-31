package main

import (
	"fmt"
	"io"
	"os"
)

/**
拷贝文件
*/
func main() {
	_, err := copyFile("./src.txt", "./dst.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}
	fmt.Println("copy done!")

}

// copyFile 打开两个文件
func copyFile(srcPath, targetPath string) (written int64, err error) {

	// 打开源文件
	src, err := os.Open(srcPath)
	if err != nil {
		fmt.Printf("src file open faild ,err:%d\n", nil)
	}
	defer src.Close()

	// 设置目标文件
	target, err := os.OpenFile(targetPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("src file open faild ,err:%d\n", nil)
	}
	defer target.Close()
	return io.Copy(target, src) // 调用io.Copy()拷贝内容
}

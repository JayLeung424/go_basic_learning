package main

import "fmt"

// if判断
func main() {
	// 写法1
	var score = 65
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// 特殊写法
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score >= 70 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}

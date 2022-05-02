package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "/Users/jiel/File/adZipPath/20220427231456/自动化/110.jpg"
	rex := regexp.MustCompile(`.*\/(\d+)\.(jpg)`)
	str1 := rex.ReplaceAllString(str, "$1")
	fmt.Println(str1)
}

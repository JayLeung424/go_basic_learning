package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 使用http实现 web 开发

func sayHello(writer http.ResponseWriter, r *http.Request) {
	// 找到对应的h5页面
	file, _ := ioutil.ReadFile("./hello.html")
	_, _ = fmt.Fprintln(writer, string(file))
}

func main() {
	// 当浏览器执行 /hello 时候，执行 sayHello()
	http.HandleFunc("/hello", sayHello)
	// 启动服务
	err := http.ListenAndServe("127.0.0.1:9001", nil)
	if err != nil {
		fmt.Printf("http server failed ,err:%v\n", err)
	}
}

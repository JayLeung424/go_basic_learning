package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server

func process(con net.Conn) {
	defer con.Close() // 处理完以后要关闭链接
	// 针对当前的连接做数据的发送和接收操作
	for {
		reader := bufio.NewReader(con)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed,err:%v\n", err)
			break
		}

		recv := string(buf[:n])
		fmt.Println("接收到的数据:", recv)
		con.Write([]byte(recv)) // 把收到的数据返回给客户端
	}

}

func main() {
	// 1、开启服务
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err信息:", err)
		return
	}
	for {
		// 2、等待客户端建立连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%v\n", err)
			continue
		}
		// 3、启动一个单独的goroutine去处理链接
		go process(conn)
	}

}

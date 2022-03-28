package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.dial err=", err)
		return
	}
	defer conn.Close()

	//从键盘输入 我给它一个协程
	go func() {
		// 从键盘输入，给服务器发送内容
		str := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(str) //从键盘输入，放入到str里面
			if err != nil {
				fmt.Println("os.stdin.err=", err)
				return
			}
			// 把内容发给服务器
			conn.Write(str[:n])
		}
	}()

	// 接收服务器回复的数据
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收服务器请求
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println("客户端读取服务端返回的数据：", string(buf[0:n]))
	}

}

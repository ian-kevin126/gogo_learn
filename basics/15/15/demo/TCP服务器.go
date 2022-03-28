package main

import (
	"fmt"
	"net"

	"strings"
)

func HandleConn(conn net.Conn) {
	// 函数调用完毕，自动关闭conn
	defer conn.Close()

	// 获取客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect successful")
	buf := make([]byte, 2048)

	for {
		n, err := conn.Read(buf)
		fmt.Println("nnnn", n) // 返回的是客户端输入的字符长度
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Printf("[%s]:%s\n", addr, string(buf[:n]))
		fmt.Println("len=", len(string(buf[:n])))

		if "exit" == string(buf[0:n-2]) { //"\r\n"
			fmt.Println(addr, "exit")
			return
		}

		//把用户小写转换成大写 在发送给用户
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	//接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
			return
		}

		//处理用户请求，新建立一个协程
		go HandleConn(conn)
	}

}

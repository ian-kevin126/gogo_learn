package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//文件的接收操作
func RecvFile(filename string, conn net.Conn) {
	//新建文件
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) //接收对方发送过来的文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err=", err)

			}
			return
		}
		if n == 0 {
			fmt.Println("n==0 文件接收完毕")
			break
		}

		// 往文件里写如内容
		f.Write(buf[:n])
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	defer listener.Close()
	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err=", err)
		return
	}
	defer conn.Close()

	//缓冲
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err =", err)
		return
	}

	filename := string(buf[:n])
	//回复ok
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(filename, conn)

}

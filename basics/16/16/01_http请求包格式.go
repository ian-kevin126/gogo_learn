package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listener err = ", err)
		return
	}

	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err =", err)
		return
	}

	defer conn.Close()

	//接收客户端数据
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read =", err)
		return
	}
	if n == 0 {
		fmt.Println("read err=", err)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))

}

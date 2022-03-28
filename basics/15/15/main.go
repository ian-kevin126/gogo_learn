package main

import (
	"fmt"
	"net"
)

/*
	应用层  <——————应用程序打牌应用程序（协议：FTP、TFTP、Telnet、NFS）——————>  应用层
	传输层  <——————         进程到进程（协议：TCP、UDP）             ——————>  传输层
	网络层  <——————         主机到主机（协议：IP、ICMP、IGMP）       ——————>  网络层
	链路层  <——————         设备到设备（协议：ARP、RARP）            ——————>  链路层
*/

/**
服务器端代码
*/
func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer listener.Close()

	// 阻塞等待用户的连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	//接收用户的请求
	buf := make([]byte, 1024) //1024大小进行缓冲

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println("buf=", string(buf[:n]))
	defer conn.Close()

}

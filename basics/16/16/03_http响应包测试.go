package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("dail err =", err)
		return
	}

	defer conn.Close()

	// 模仿客户端访问服务端
	requestBuf := "GET /go HTTP/1.1\r\nAccept: image/gif, image/jpeg, image/pjpeg, application/x-ms-application," +
		" application/xaml+xml, application/x-ms-xbap, */*\r\nAccept-Language: zh-Hans-CN,zh-Hans;q=0.8," +
		"en-US;q=0.5,en;q=0.3\r\nUser-Agent: Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; WOW64; Trident/7.0; " +
		".NET4.0C; .NET4.0E; .NET CLR 2.0.50727; .NET CLR 3.0.30729; .NET CLR 3.5.30729)\r\nAccept-Encoding: gzip, " +
		"deflate\r\nHost: 127.0.0.1:8000\r\nConnection: Keep-Alive\r\n\r\n"

	//发送请求包，服务器才会有响应
	conn.Write([]byte(requestBuf))

	//接收服务器端回复的数据
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)

	if err != nil || n == 0 {
		fmt.Println("read err=", err)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))

}

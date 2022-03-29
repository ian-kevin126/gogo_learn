package main

import (
	"fmt"
	"net/http"
)

func main() {

	//resp,err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http.Get err=", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("status=", resp.Status)
	fmt.Println("statusCode=", resp.StatusCode)
	fmt.Println("header=", resp.Header)

	//fmt.Println("body=",resp.Body)

	buf := make([]byte, 1024*4)

	var temp string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err=", err)
			break
		}
		temp += string(buf[:n])
	}

	fmt.Println("tem=", temp)
	/*
		status= 200 OK
		statusCode= 200
		header= map[Content-Length:[20] Content-Type:[text/html; charset=utf-8] Date:[Mon, 28 Mar 2022 09:32:49 GMT]]
		read err= EOF
		tem= <h1>hello world</h1>
	*/

}

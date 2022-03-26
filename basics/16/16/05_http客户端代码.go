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
	fmt.Println("statuscode=", resp.StatusCode)
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

}

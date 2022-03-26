package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args //接收用户传递参数，都是以字符串方式传递

	n := len(list)
	fmt.Println("n=", n)

	for i := 0; i < n; i++ {
		fmt.Printf("list[%d]=%s\n", i, list[i])
	}

	fmt.Println("-----------------------------")

	for i, data := range list {
		fmt.Printf("list[%d]=%s\n", i, data)
	}

	// 执行：go run 09获取命令行参数.go 10 20
}

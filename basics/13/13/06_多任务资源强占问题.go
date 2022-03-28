package main

import (
	"fmt"
	"time"
)

//定义一个打印机，参数字符串，按照每个字符打印
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("hello")
}

func Person2() {
	Printer("world")
}
func main() {
	//2个协程共有一个资源
	go Person1()
	go Person2()

	for {

	}

	// 输出：hweorllldo
}

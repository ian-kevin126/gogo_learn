package main

import (
	"fmt"
	"time"
)

//定义全局的变量，创建一个channel
var ch = make(chan int)

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
	ch <- 66666

}

func Person2() {
	<-ch //从管道读取数据，接收，如果没有数据它会阻塞
	Printer("world")
}
func main() {
	//2个协程共有一个资源
	go Person1()
	go Person2()

	for {

	}

}

package main

import (
	"fmt"
	"time"
)

//创建一个管道
//var ch = make(chan int)
var ch = make(chan int)

//定义一个打印机，参数字符串，按照每个字符打印
func Printer1(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//person3 执行完之后 就会执行person4 ，在协程直接通信我们一般使用chan
func person3() {
	Printer1("hello")
	ch <- 6666 //给管道写数据 发送
}

func person4() {
	<-ch //重管道读取数据，接收，如果通道没有数据就会阻塞
	Printer1("world")

}
func main() {
	go person3()
	go person4()

	for {

	}
}

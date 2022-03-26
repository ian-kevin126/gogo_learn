package main

import "fmt"

//定义在{}外部的变量就是全局变量，全局变量在任何地方都能使用，伴生程序

func Test() {
	fmt.Println("test a", a) // test a 10
}

var a int // 全局变量

func main() {

	a = 10

	fmt.Println("a", a) // a 10

	Test()
}

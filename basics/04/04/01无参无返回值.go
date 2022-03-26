package main

import "fmt"

//无参无返回值

func MyFunction() {
	b := "发射子弹"
	fmt.Println(b)
}

func main() {
	//无参无返回值函数调用: 函数名()
	MyFunction()
	MyFunction()
}

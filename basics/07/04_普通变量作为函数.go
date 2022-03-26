package main

import "fmt"

func swap(a, b int) {
	a, b = b, a
	fmt.Printf("swap:a=%d,b=%d\n", a, b) //20 10
}

func main() {
	a, b := 10, 20
	//通过一个函数swap 交换a b的内容
	swap(a, b)
	fmt.Printf("main: a=%d,b=%d\n", a, b) //10 20
}

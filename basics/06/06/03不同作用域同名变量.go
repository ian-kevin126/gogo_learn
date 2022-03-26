package main

import "fmt"

var aa byte //全局变量

func main() {

	var aa int // 局部变量
	// testaa()

	fmt.Printf("%T\n", aa)

	{
		//不同作用域，允许定义同名变量
		//使用变量的原则，就近原则
		var aa float32
		fmt.Printf("2222: %T\n", aa) //float32
	}
}

func testaa() {
	fmt.Printf("444 %T\n", aa)
}

package main

import "fmt"

func test1(a int) {

	if a == 1 { //函数的终止条件，非常重要
		fmt.Println("a=", a)
		return //终止条件
	}
	//函数调用本身
	test1(a - 1)
	fmt.Println("a=", a)
}

func main() {
	test1(3)
	fmt.Println("main")
}

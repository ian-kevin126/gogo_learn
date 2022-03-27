package main

import (
	"fmt"
)

func main() {
	//声明定义同时初始化

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a=", a) // a= [1 2 3 4 5]

	b := [5]int{1, 2, 3, 4, 5} //自动推导
	fmt.Println("b=", b)       // b= [1 2 3 4 5]

	//部分初始化 没有初始化的int自动补全0
	c := [5]int{1, 2, 3}
	fmt.Println("c=", c) // c= [1 2 3 0 0]

	//指定某个元素的初始化
	d := [5]int{2: 10, 4: 20}
	fmt.Println("d=", d) // d= [0 0 10 0 20]

}

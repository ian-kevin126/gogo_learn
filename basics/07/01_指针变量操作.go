package main

import "fmt"

func main() {

	var a int = 10
	//每个变量有2层含义：变量的内存 变量地址
	fmt.Printf("a=%d\n", a)     //变量的内存
	fmt.Printf("&a = %v\n", &a) //变量的地址

	//保存地址必须要用指针（地址）

	var p *int //定义一个变量p，类型为int类型
	p = &a     //指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p=%v,&a=%v,*p=%v\n", p, &a, *p)

	*p = 1000
	fmt.Printf("p=%v,&a=%v,*p=%v\n", p, &a, *p)

}

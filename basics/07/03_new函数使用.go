package main

import "fmt"

func main() {
	//a := 10     //整形变量
	var p *int //指针变量
	// p = &a
	// *p = 1000 //a = 1000

	p = new(int)
	fmt.Printf("p=%v,*p=%d\n", p, *p)

	*p = 6000

	fmt.Printf("p=%v,*p=%d\n", p, *p)

	q := new(int) //自动推导
	*q = 88888
	fmt.Println("*q=", *q)
}

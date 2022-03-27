package main

import "fmt"

type Student2 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//先定义一个普通变量
	var s Student2
	//定义一个指针变量
	var p1 *Student2
	p1 = &s

	p1.id = 1
	(*p1).name = "jack"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "bj"
	fmt.Println("s=", s) // s= {1 jack 109 18 bj}

	p2 := new(Student2) //new 地址
	p2.id = 1
	p2.name = "mike"
	p2.sex = 'M'
	p2.age = 18
	p2.addr = "bj"
	fmt.Println("p2=", p2) // p2= &{1 mike 77 18 bj}

}

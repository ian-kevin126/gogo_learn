package main

import "fmt"

type Person5 struct {
	name string //名字
	sex  byte   // 姓名
	age  int    // 年龄
}

type Student5 struct {
	*Person5 //指针类型 Person5地址
	id       int
	addr     string
}

func main() {
	s1 := Student5{&Person5{"mike", 'm', 18}, 777, "bj"}
	fmt.Printf("s1=%+v\n", s1)
	fmt.Println(s1.Person5.name)
	fmt.Println(s1.name)

	//先定义变量
	var s2 Student5
	s2.Person5 = new(Person5) //分配空间地址
	s2.name = "tom"
	s2.sex = 'm'
	s2.age = 18
	s2.id = 2222
	s2.addr = "bj"
	fmt.Println(s2.name, s2.sex, s2.age, s2.id, s2.addr)

}

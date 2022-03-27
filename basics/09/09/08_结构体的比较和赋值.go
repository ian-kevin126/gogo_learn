package main

import "fmt"

type Student3 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	s1 := Student3{1, "mike", 'm', 18, "bj"}
	s2 := Student3{1, "mike", 'm', 18, "bj"}
	s3 := Student3{2, "mike", 'm', 18, "bj"}
	fmt.Println("s1==s2", s1 == s2) // s1==s2 true
	fmt.Println("s1==s3", s1 == s3) // s1==s3 false

	//同类型的2个结构体变量是可以互相赋值的
	var tmp Student3
	tmp = s1
	fmt.Println("tmp=", tmp) // tmp= {1 mike 109 18 bj}
}

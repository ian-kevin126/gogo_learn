package main

import "fmt"

type Person1 struct {
	name string //名字
	sex  byte   // 姓名
	age  int    // 年龄
}

type Student1 struct {
	Person1 //只有类型没有字段，匿名字段 基础person所有的成员
	id      int
	addr    string
}

func main() {
	s1 := Student1{Person1{"mike", 'm', 18}, 1, "bj"}

	//对象成员的操作方式一
	s1.name = "yoyo"
	s1.sex = 'f'
	s1.age = 22
	s1.id = 7777
	s1.addr = "sh"

	//对象操作匿名字段方式二
	s1.Person1 = Person1{"tom", 'm', 19}

	fmt.Println(s1.Person1.name)
	fmt.Println(s1.name) // s1.Person1.name == s1.name
	fmt.Println("s1=", s1)

}

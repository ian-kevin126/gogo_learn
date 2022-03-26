package main

import "fmt"

type Person2 struct {
	name string //名字
	sex  byte   // 姓名
	age  int    // 年龄
}

type Student2 struct {
	Person2 //只有类型没有字段，匿名字段 基础person所有的成员
	id      int
	addr    string
	name    string //和person同名了
}

func main() {
	//声明（定义一个变量）
	var s Student2
	//就近原则：如果能在本作用域找到此成员，就操作此成员，如果找不到，就找到继承的字段
	s.name = "zhangsan" //Student的name
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	fmt.Printf("s=%+v\n", s)

}

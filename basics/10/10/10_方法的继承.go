package main

import "fmt"

type Person6 struct {
	name string
	sex  byte
	age  int
}

//Person6类型，实现了一个方法
func (tmp Person6) PrintInfo() {
	fmt.Printf("name=%s,sex=%c,age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有个学生，继承person字段，成员和方法都继承了
type Student6 struct {
	Person6 //
	id      int
	addr    string
}

//Student6 实现了一个方法，这个方法和Person方法同名，这种叫方法重写
func (temp Student6) PrintInfo() {
	fmt.Println("Student6=", temp)
}

func main() {
	s := Student6{Person6{"mike", 'm', 18}, 666, "bj"}
	//就近原则：先找本作用域的方法，找不到在去找继承的方法，如果找不到 调用报错
	s.PrintInfo()
	//显示调用继承的方法
	s.Person6.PrintInfo()
}

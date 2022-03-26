package main

import "fmt"

type mystr string //自定义类型，给一个类型改名

type Person3 struct {
	name string //名字
	sex  byte   // 姓名
	age  int    // 年龄
}

type Student3 struct {
	Person3 //只有类型没有字段，匿名字段 基础person所有的成员
	int     //基础类型的匿名字段
	mystr
}

func main() {
	s := Student3{Person3{"mike", 'm', 18}, 666, "zhangsan"}
	fmt.Printf("s=%+v\n", s)
	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person3, s.int, s.mystr)

}

package main

import "fmt"

type Person struct {
	name string // 名字
	sex  byte   // 姓名
	age  int    // 年龄
}

type Student struct {
	Person // 只有类型没有字段，匿名字段 基础person所有的成员
	id     int
	addr   string
}

func main() {
	//结构体顺序初始化，把所有字段都要初始化
	var s1 Student = Student{Person{"mike", 'm', 18}, 1, "bj"}
	fmt.Println("s1=", s1)                   // s1= {{mike 109 18} 1 bj}
	fmt.Println("s1.Person", s1.Person.name) // s1.Person mike

	//自动类型推导
	s2 := Student{Person{"mike", 'm', 18}, 1, "bj"}
	//%+v 显示更详细的信息
	fmt.Printf("s2=%+v\n", s2) // s2={Person:{name:mike sex:109 age:18} id:1 addr:bj}

	//指定类型初始化，没有初始化的常用类型按照默认值赋值 int 0 string 为空
	s3 := Student{id: 1}
	fmt.Printf("s3=%+v\n", s3) // s3={Person:{name: sex:0 age:0} id:1 addr:}

	s4 := Student{Person: Person{name: "zhangsan"}, id: 1}
	fmt.Printf("s4=%+v", s4) // s4={Person:{name:zhangsan sex:0 age:0} id:1 addr:}

	//s5 := Student{"mike",'m',18,1,"bj"} //err

}

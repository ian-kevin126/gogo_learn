package main

import "fmt"

func main() {
	//定义一个map变量，类型为map[int]string,如果使用map一定初始化，make分配空间
	var m1 map[int]string = make(map[int]string)
	fmt.Println("m1=", m1) // m1= map[]
	m1[1] = "jake"
	fmt.Println("m1=", m1) // m1= map[1:jake]

	m2 := make(map[int]string)   //make方式创建然后自动推导类型
	fmt.Println("len=", len(m2)) // len=0
	m2[1] = "mike"
	fmt.Println("m2=", m2)       // m2= map[1:mike]
	fmt.Println("len=", len(m2)) // len=1

	//map 先给map指定一个可以容纳长度，一旦超过这个长度 重新分配底层空间
	m3 := make(map[int]string, 2)
	m3[1] = "mile"
	m3[2] = "jack"
	m3[3] = "go"
	fmt.Println("m3=", m3)       // m3= map[1:mile 2:jack 3:go]
	fmt.Println("len=", len(m3)) // len=3

	//map另一种初始化并且赋值
	m4 := map[int]string{1: "mike", 2: "jack", 3: "go"}
	fmt.Println("m4=", m4) // m4= map[1:mike 2:jack 3:go]
}

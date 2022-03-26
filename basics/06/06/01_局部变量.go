package main

import (
	"fmt"
)

func test() {
	a := 1000
	fmt.Println("a", a)
}

func test1(a int) { //test1  b = a
	a = 2000
	fmt.Println("a=", a)

}

func main() {
	//定义在{}里面的变量就是局部变量，只能在{里面}有效，执行到定义变量的那句话，才开始分配空间，
	//离开{}自动释放
	a := 111
	fmt.Println("a", a)

	test1(a)

	{
		b := 222
		fmt.Println("b", b)
	}
	//fmt.Println("b=",b)

	if flag := 3; flag == 3 {
		fmt.Println("flag=", flag)
	}

	//fmt.Println("flag=",flag)
}

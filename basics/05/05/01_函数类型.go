package main

import "fmt"

//通过type给一个函数类型起名 函数也是一种数据类型（int byte）
type FuncType func(a, b int) int //没有名字 没有{}

func Add(a, b int) int {
	return a + b
}

func Minus(c, d int) int {
	return c - d
}

func main() {
	var a int
	a = 10
	fmt.Println("a=", a)

	var fTest FuncType
	fTest = Add
	result := fTest(3, 12) //等价于Add(3,12)
	fmt.Println("result=", result)

	fTest = Minus
	result = fTest(3, 2) //等价于Minus（3,2）
	fmt.Println("result=", result)
}

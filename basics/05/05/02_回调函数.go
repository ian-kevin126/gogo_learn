package main

import (
	"fmt"
)

type FuncType01 func(int, int) int

//回调函数 就是函数有一个参数是函数类型，这个函数就是回调函数
//fTest FuncType01 多态接口 调用一个接口，可以有不同实现方式

func Calc(a, b int, fTest FuncType01) (result int) { //fTest = Add1
	fmt.Println("Calc")
	result = fTest(a, b)
	//result =Add1（3，4）
	return
}

func Add1(a, b int) int {
	return a + b
}

func Minus1(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func main() {

	a := Calc(3, 4, Add1)
	fmt.Println("a=", a)

	a = Calc(7, 5, Minus1)
	fmt.Println("a=", a)

	a = Calc(4, 3, Mul)
	fmt.Println("a=", a)

}

package main

import "fmt"

func test() func() int {
	var x int //x没有初始化，值为0
	return func() int {
		x++
		return x * x
	}
}

func main() {
	//返回值是一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来	调用闭包函数
	//f不关心这些捕获的变量和常量是否超出作用范围，只要你的f存在，还在使用x，这个变量x就会一直存在

	f := test()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16

	f = test()

	fmt.Println(f()) // 1
}

/*
func test() int  {
	//函数调用时候，x才会分配空间，才初始化为0
	var x int //没有初始化，值为0
	x++
	return x*x //函数调用完毕，x自动释放
}
func main(){
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())
	fmt.Println(test())

}
*/

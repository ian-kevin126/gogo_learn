package main

import (
	"fmt"
)

/*func (){

}*/
func main() {

	a := 10
	str := "mike"
	//匿名函数 ：就是名字的函数,函数在定义 还没有被调用
	f1 := func() {
		fmt.Println("匿名函数")
	}
	//fmt.Printf("%T\n",f1)
	f1()

	/*f2 := func(a,b int){
		fmt.Println("a+b=",a+b)
	}
	//fmt.Printf("f2= %T\n",f2)

	f2(2,3)*/

	//给一个函数类型起一个别名
	type FuncType03 func()
	//声明别名
	var f2 FuncType03
	f2 = f1
	f2()

	//匿名函数 定义匿名函数的时候同时调用
	func() {
		fmt.Println("匿名函数直接调用")
	}()

	func(a, b int) {
		fmt.Println("匿名函数直接调用", a, b)
	}(2, 3)

	func() {
		fmt.Printf("a=%d,str=%s\n", a, str)
	}()

	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(10, 20)
	fmt.Printf("x=%d,y=%d", x, y)

}

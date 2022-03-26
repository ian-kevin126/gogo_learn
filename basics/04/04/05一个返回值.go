package main

import "fmt"

//无参有返回值，只有一个返回值，通过return返回
func myfunc01() int {
	return 5555
}

//这种是官方推荐 给返回值起一个名字
func myfunc02() (a int) {
	return 6666
}

//常用法
func myfunc03() (a int) {
	a = 666
	return
}

func main() {
	var a int
	a = myfunc01()
	fmt.Println("a=", a)

	b := myfunc01()
	fmt.Println("b=", b)

	c := myfunc03()
	fmt.Println("c=", c)

}

package main

import "fmt"

func main() {
	a := 20
	str := "aaa"

	func() {
		//闭包以引用的方式捕获外部变量
		a = 40
		str = "bbb"
		fmt.Printf("a=%d,str=%s\n", a, str) // a=40,str=bbb
	}()

	fmt.Printf("a=%d,str=%s\n", a, str) // a=40,str=bbb

}

package main

import "fmt"

func myfunc001() (int, int, int) {
	return 11, 22, 33
}

//go 推荐写法
func myfunc002() (a int, b int, c int) {
	a, b, c = 11, 22, 33
	return
}

func myfunc003() (a, b, c int) {
	a, b, c = 11, 22, 33
	return
}

func main() {
	a, b, c := myfunc001()
	// fmt.Printf("a=%d,b=%d,c=%d",a,b,c)
	a, b, c = myfunc002()
	fmt.Printf("a=%d,b=%d,c=%d", a, b, c)
}

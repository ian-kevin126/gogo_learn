package main

import "fmt"

//函数定义 比较俩个数的大小
func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return
}

func main() {
	max, min := MaxAndMin(10, 15)
	fmt.Printf("max=%d,min=%d\n", max, min)

	//通过匿名变量丢弃某个返回值 _
	a, _ := MaxAndMin(10, 15)
	fmt.Printf("a=%d", a)

	_, b := MaxAndMin(10, 15)
	fmt.Printf("b=%d", b)
}

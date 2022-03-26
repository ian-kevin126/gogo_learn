package main

import "fmt"

func main() {
	//只支持 == ！= 比较是不是每个元素都一样 2个数组比较类型要一致 [3]int [4]int
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}

	fmt.Println("a==b", a == b)
	fmt.Println("a==c", a == c)
	//同类型的数组可以赋值
	var d [5]int
	d = c
	fmt.Println("d=", d)

}

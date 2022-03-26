package main

import (
	"fmt"
)

func main() {
	//切片和数组区别
	//数组 数组【】的长度是固定的常量 ，数组不能修改长度，len和cap固定
	a := [5]int{}
	fmt.Printf("len=%d,cap=%d\n", len(a), cap(a))
	//切片 []里面为空 【...】切片长度不固定 可以追加
	s := []int{} //底层引入一个数组 不能转载下 就会新引入地址 不过对切片本身来说不变
	//只要知道切片不定长就可以 arraylist
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))

	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))
	s = append(s, 1)
	fmt.Printf("len=%d,cap=%d\n", len(s), cap(s))

}

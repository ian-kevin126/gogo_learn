package main

import "fmt"

func main() {
	//自动推导类型，同时进行初始化
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1=", s1) // s1= [1 2 3 4]

	//借助make的方式创建切片（类型 长度 容量）
	s2 := make([]int, 5, 10)
	fmt.Println(s2)                                 // [0 0 0 0 0]
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2)) // len=5,cap=10

	//如果没有指定容量  容量和长度一样
	s3 := make([]int, 5)
	fmt.Printf("len=%d,cap=%d\n", len(s3), cap(s3)) // len=5,cap=5
}

package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:hight:max] s【1:2:4】 取下标从low开始 len =hight-low cap =max-low
	s1 := array[:] //[0:len(array):len(array)] 容量和长度是一致
	fmt.Println("s1=", s1)
	fmt.Printf("len=%d,cap=%d\n", len(s1), cap(s1))

	//操作某个元素 和数组操作方式一样
	data := array[1]
	fmt.Println("data=", data)

	s2 := array[3:6:7] //a[3] a[4] a[5]  len = 6-3 cap = 7-3
	fmt.Println("s2=", s2)
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2))

	s3 := array[:5] //从0开始 6个元素 容量6
	fmt.Println("s3=", s3)
	fmt.Printf("len=%d,cap=%d\n", len(s3), cap(s3))

	s4 := array[3:]

	fmt.Println("s4=", s4)
	fmt.Printf("len=%d,cap=%d\n", len(s4), cap(s4))

}

package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] s【1:2:4】 取下标从low开始 len = high-low cap =max-low

	s1 := array[:]                                  //[0:len(array):len(array)] 容量和长度是一致
	fmt.Println("s1=", s1)                          // s1= [0 1 2 3 4 5 6 7 8 9]
	fmt.Printf("len=%d,cap=%d\n", len(s1), cap(s1)) // len=10,cap=10

	//操作某个元素 和数组操作方式一样
	data := array[1]
	fmt.Println("data=", data) // data=1

	s2 := array[3:6:7]                              //a[3] a[4] a[5]  len = 6-3 cap = 7-3
	fmt.Println("s2=", s2)                          // s2= [3 4 5]
	fmt.Printf("len=%d,cap=%d\n", len(s2), cap(s2)) // len=3,cap=4

	// 注意这种切片的容量，它会取原数组的cap
	s3 := array[:5]                                 //从0开始 5个元素 容量10
	fmt.Println("s3=", s3)                          // s3= [0 1 2 3 4]
	fmt.Printf("len=%d,cap=%d\n", len(s3), cap(s3)) // len=5,cap=10
	s3 = append(s3, 11)
	fmt.Println("s3=", s3)                          // s3= [0 1 2 3 4 11]
	fmt.Printf("len=%d,cap=%d\n", len(s3), cap(s3)) // len=6,cap=10

	s4 := array[3:]
	fmt.Println("s4=", s4)                          // s4= [3 4 5 6 7 8 9]
	fmt.Printf("len=%d,cap=%d\n", len(s4), cap(s4)) // len=7,cap=7
	s4 = append(s4, 12)
	fmt.Printf("len=%d,cap=%d\n", len(s4), cap(s4)) // len=8,cap=14
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {

	//设置种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) //100 以内的值
	}

}

//冒泡排序
func BubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	n := 10

	//创建切片
	s := make([]int, n)
	InitData(s)
	fmt.Println("排序前", s)
	BubbleSort(s)

	fmt.Println("排序后", s)

}

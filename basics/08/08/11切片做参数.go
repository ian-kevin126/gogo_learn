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

// BubbleSort 冒泡排序
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
	fmt.Println("排序前", s) // 排序前 [67 37 70 72 75 42 93 62 52 24]
	BubbleSort(s)

	fmt.Println("排序后", s) // 排序后 [24 37 42 52 62 67 70 72 75 93]

}

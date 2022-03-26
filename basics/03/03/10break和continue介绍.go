package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		time.Sleep(time.Second) //睡觉1s
		if i == 5 {
			//break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环，本层循环
			continue //跳出本次循环，继续下一次循环
		}

		fmt.Println("i=", i)

	}
}

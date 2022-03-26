package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//如果种子参数一致，每次运行的程序产生的结果是一致的
	//每次运行让种子参数不一致
	//rand.Seed(6666)
	rand.Seed(time.Now().UnixNano()) // 以当前系统时间作为种子参数

	for i := 0; i < 6; i++ {
		//fmt.Println("rand=",rand.Int()) //随机数很大
		fmt.Println("rand=", rand.Intn(100)) //限制在100以内的数字
	}

}

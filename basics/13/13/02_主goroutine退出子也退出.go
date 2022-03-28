package main

import (
	"fmt"
	"time"
)

//主程序退出 子程序跟着退出，子程序退出需要时间
func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("main i=", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}

	/*
		main i= 1
		子协程 i= 1
		子协程 i= 2
		main i= 2
	*/
}

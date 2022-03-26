package main

import (
	"fmt"
	"time"
)

func main() {
	//创建channel
	ch := make(chan string)
	defer fmt.Println("主协程也结束")

	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程i=", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程，子协程工作完毕"
	}()

	str := <-ch //没有数据，阻塞
	fmt.Println("str=", str)

}

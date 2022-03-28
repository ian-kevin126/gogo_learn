package main

import (
	"fmt"
	"time"
)

func main() {
	//创建channel
	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程i=", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协工程，工作完毕"
	}()

	str := <-ch //没有数据前 阻塞
	fmt.Println("str=", str)

	/*
		子协程i= 0
		子协程i= 1
		str= 我是子协工程，工作完毕
		子协程调用完毕
	*/
}

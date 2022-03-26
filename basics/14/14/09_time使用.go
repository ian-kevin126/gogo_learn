package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写内容
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("当前时间：", time.Now())

	t := <-timer.C //channel 没有数据前后阻塞
	fmt.Println("t=", t)

}

package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写内容
	timer := time.NewTimer(3 * time.Second)
	fmt.Println("当前时间：", time.Now()) // 当前时间： 2022-03-28 12:54:36.694305 +0800 CST m=+0.000164296

	t := <-timer.C       //channel 没有数据前后阻塞
	fmt.Println("t=", t) // t= 2022-03-28 12:54:39.69559 +0800 CST m=+3.001359924
}

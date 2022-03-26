package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) //创建一个无缓存的channle

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		close(ch) //不需要写数据的时候 我们关闭channel
		// ch <-666  //send on closed channel 关闭channel后无法在发送数据
	}()

	for {
		//如果ok为true,说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num=", num)
		} else { //管道关闭
			break
		}
	}
}

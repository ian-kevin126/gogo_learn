package main

import (
	"fmt"
	//"time"
)

//创建通道只能写，不能读\
//out chan<- int
func producer(out chan<- int) { //out chan<- int = ch
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num=", num)
	}
}

func main() {
	//创建一个双向管道
	ch := make(chan int)
	//生成者，生成数据,开启一个协程
	go producer(ch)
	//消费者，从channel管道读取内容，打印
	consumer(ch)

	// time.Sleep(2*time.Second)

}

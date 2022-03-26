package main

import (
	"fmt"

	"time"
)

func main() {
	//创建一个有缓存区的管道
	ch := make(chan int, 3)
	//len(ch) 缓存区里面有多个数据，cap(ch )缓存区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))

	go func() {

		for i := 0; i < 10; i++ {
			ch <- i //往chan写内容
			fmt.Printf("子协程[%d]:len(ch)=%d,cap(ch)=%d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}
}

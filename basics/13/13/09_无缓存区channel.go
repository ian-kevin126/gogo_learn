package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	//len(ch)缓存区已使用数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d", len(ch), cap(ch))

	go func() {

		for i := 0; i < 3; i++ {
			fmt.Printf("子协程：i=%d\n", i)
			ch <- i //往chan写内容

		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}

}

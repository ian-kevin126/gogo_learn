package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
		}
		//不需要写数据的时候，关闭channel
		close(ch)
	}()

	for num := range ch {
		fmt.Println("num=", num)
	}
}

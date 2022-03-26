package main

import "fmt"

func fibonacci(ch chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		//监听channel数据流动
		select {
		case ch <- x:
			x, y = y, x+y
		case flag := <-quit:
			fmt.Println("flag=", flag)
			return

		}
	}
}

func main() {
	ch := make(chan int) //数字通信

	quit := make(chan bool) //程序是否退出 true

	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println(num)
		}
		quit <- true

	}()

	fibonacci(ch, quit)

}

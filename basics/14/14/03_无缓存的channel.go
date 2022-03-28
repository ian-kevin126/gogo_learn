package main

import (
	"fmt"
	"time"
)

/*
无缓冲的channel：是指在接收前没有能力保存任何值的通道

这种类型的通道要求发送goroutine和接收goroutine同时准备好，才能完成发送和接收操作。
如果两个goroutine没有同时准备好，通道会导致先执行发送或接收操作的goroutine阻塞等。

这种对通道进行发送和接收的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作而单独存在。

*/

func main() {
	//创建一个无缓存区的channel
	ch := make(chan int, 0)
	//len(ch)缓冲区使用数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch)) // len(ch)=0,cap(ch)=0

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程:i=%d\n", i)
			ch <- i //往chan写内容
			fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))
		}
	}()

	//延时2秒
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读管道内容，没有内容前，阻塞
		fmt.Println("num=", num)
	}

	/*
		子协程:i=0
		len(ch)=0,cap(ch)=0
		子协程:i=1
		num= 0
		num= 1
		len(ch)=0,cap(ch)=0
		子协程:i=2
		len(ch)=0,cap(ch)=0
		num= 2
	*/
}

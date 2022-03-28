package main

import (
	"fmt"

	"time"
)

/*
使用有缓冲的通道在goroutine之间同步数据
- 第一步，右侧的goroutine正在从通道接收一个值
- 第二步，右侧的这个goroutine独立完成了接收值的动作，而左侧的goroutine正在发送一个新值到通道里
- 第三步，左侧的goroutine还在向通道发送新值，而右侧的goroutine正在从通道接收另外一个值，这个步骤里的两个操作既不是同步的，也不会相互阻塞。
- 最后，所有的发送和接收都完成，而通道里还有几个值，也有一些空间可以存更多的值
*/

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

	/*
		len(ch)=0,cap(ch)=3
		子协程[0]:len(ch)=1,cap(ch)=3
		子协程[1]:len(ch)=2,cap(ch)=3
		子协程[2]:len(ch)=3,cap(ch)=3
		num= 0
		num= 1
		num= 2
		num= 3
		子协程[3]:len(ch)=0,cap(ch)=3
		子协程[4]:len(ch)=0,cap(ch)=3
		子协程[5]:len(ch)=1,cap(ch)=3
		子协程[6]:len(ch)=2,cap(ch)=3
		子协程[7]:len(ch)=3,cap(ch)=3
		num= 4
		num= 5
		num= 6
		num= 7
		num= 8
		子协程[8]:len(ch)=0,cap(ch)=3
		子协程[9]:len(ch)=0,cap(ch)=3
		num= 9
	*/
}

package main

import (
	"fmt"
	"time"
)

/*
select的作用：
Go里面提供一个关键字select，通过select可以监听channel上的数据流动。

select的写法与switch非常类似，由select开始一个新的选择快，每个选择条件由case语句来描述。

与switch语句可以选择任何可使用相等比较的条件相比，select有比较多的限制，其中最大的一条限制就是：
—— 每个case语句里必须是一个IO操作。

*/

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num=", num)
				// 超时执行语句（超过3s没有人读取ch里面的数据，就执行超时逻辑）
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
				//break
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	qt := <-quit
	fmt.Println("程序结束:qt=", qt)

}

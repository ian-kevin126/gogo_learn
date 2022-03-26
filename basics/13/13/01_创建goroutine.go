package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a newTask")
		time.Sleep(time.Second) //延时1s
	}
}

func main() {
	go newTask() //go 关键字就新建一个协程，新建一个任务

	/*for{
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}*/
	i := 0
	for {
		i++
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
		if i == 2 {
			break
		}

	}

}

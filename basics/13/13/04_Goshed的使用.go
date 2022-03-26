package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()

	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("hello")
		//让出时间片，先让别的协程执行，它执行完，最回来执行此协程

		//time.Sleep(time.Second)
	}

}

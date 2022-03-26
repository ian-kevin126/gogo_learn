package main

import (
	"fmt"
	"time"
)

//主程序退出 子程序跟着退出，导致子程序没有调用
func main() {
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i=", i)
			time.Sleep(time.Second)
		}
	}()

}

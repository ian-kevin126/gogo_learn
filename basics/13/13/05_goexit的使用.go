package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccc")
	//return //终止函数
	runtime.Goexit() //终止所有的协程
	fmt.Println("ddddddddd")
}

func main() {
	//创建新的协程
	go func() {
		fmt.Println("aaaaaaa")
		test()
		fmt.Println("bbbbbbbb")
	}()

	for {

	}

	/*

		return：
			aaaaaaa
			ccccccc
			bbbbbbbb

		goexit：
			aaaaaaa
			ccccccc
	*/
}

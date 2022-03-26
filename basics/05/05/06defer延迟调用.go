package main

import "fmt"

func main() {
	// 关键字defer用于延迟一个函数或者方法（或者当前所创建的匿名函数）的执行，需要注意的是：defer语句只能出现在函数或者方法的内部
	// defer作为延迟调用，main函数结束之前调用
	defer fmt.Println("aaaaaaa")
	defer fmt.Println("cccc")
	defer fmt.Println("dddd")
	fmt.Println("bbbbbbb")
}

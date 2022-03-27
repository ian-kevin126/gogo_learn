package main

import "fmt"

// 系统首先调用init方法，再执行main函数
func init() {
	fmt.Println("包初始化操作")
}

func main() {
	fmt.Println("main 函数")
}

/**
包初始化操作
main 函数

当一个包被导入时，如果该包还导入了其他的包，那么会先将其它包导入进来，然后再对这些包中的包级常量和变量进行初始化，接着执行init函数
（如果有的话），以此类推。等所有的被导入的包都加载完毕之后，就会开始对main包中的包级常量和变量进行初始化，然后执行main包中的init函数（如果存在）
，最后再执行main函数。
*/

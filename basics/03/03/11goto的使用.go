package main

import (
	"fmt"
)

func main() {
	fmt.Println("11111111")
	goto END                //goto是关键字，END是用户起的名字，他叫标签，goto不能夸函数调用
	fmt.Println("22222222") // 这一行不会打印
END:
	fmt.Println("333333")
}

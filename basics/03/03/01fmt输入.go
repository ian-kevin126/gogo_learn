package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入变量a:\n")
	//阻塞等待输入的值
	//fmt.Scan(&a) //一定不能忘记&
	fmt.Scanf("%d", &a) //输入的字符按照格式进行转换
	fmt.Println("a=", a)
}

package main

import "fmt"

func main() {

	var ch byte
	ch = 'A'
	var t int
	t = int(ch) //不支持隐式转换
	fmt.Println("t=", t)

	var flag bool
	flag = true
	fmt.Printf("flag=%t\n", flag)
	//这种不能转换成int类型 因为不兼容类型 bool不能转换int
	//fmt.Printf("flag=%d\n",int(flag))

}

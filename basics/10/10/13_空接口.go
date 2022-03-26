package main

import "fmt"

func xxx(arg ...interface{}) { //interface{} 空接口
	for i, k := range arg {
		fmt.Println(i, k)
	}
}

func main() {
	//空接口万能类型，保存任意类型的值
	var i interface{} = 1
	fmt.Println("i=", i)
	i = "abc"
	fmt.Println("i=", i)
	xxx(i)
}

package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for {
			fmt.Println(i)
			goto LABEL //跳转到标签LABEL，从标签处，执行代码
		}
	}

	fmt.Println("this is test")

LABEL:
	fmt.Println("it is over")
}

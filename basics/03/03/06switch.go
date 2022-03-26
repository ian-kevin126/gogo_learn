package main

import "fmt"

func main() {
	b := 10
	if b == 10 {
		fmt.Println("b==10")
	}

	if b > 10 {
		fmt.Println("b>10")
	}

	if b < 10 {
		fmt.Println("b<10")
	}

	b = 3

	switch b { //swich case 默认在最后一行加入break，一旦满足条件就不会在往下运行
	//如果想让执行语句满足条件之后继续运行 就必须加fallthrough
	case 1:
		fmt.Println("这个数字是1")
		//break
	case 2:
		fmt.Println("这个数字是2")
		//break
	case 3:
		fmt.Println("这个数字是3")
		fallthrough //只要你不跳出switch语句，后面无条件执行
		//break
	case 4:
		fmt.Println("这个数字是4")
		fallthrough
		//break
	default:
		fmt.Println("按下的是xxx")

	}

}

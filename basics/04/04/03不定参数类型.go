package main

import (
	"fmt"
)

func MyFunc11(a int) {

}

func MyFunc12(a, b int) {

}

func MyFunc13(a, b, c int) {

}

func MyFunc14(args ...int) { //...int不定参数整形  ...type 不定参数类型
	fmt.Println("len(args)=", len(args))

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d\n", i, args[i])
	}

	fmt.Println("==========================")

	for i, data := range args { //第一个代表下标，第二个代表下标对应的值
		fmt.Printf("args[%d]=%d\n", i, data)
	}

}

func MyFunc15(a int, args ...int) {
	fmt.Println("len(args)=", len(args))
}

/*func MyFunc16(args ...int,a string)  { //不定参数一定只能放在形参中的最后一个位置

}*/

func main() {
	MyFunc15(1, 2)
}

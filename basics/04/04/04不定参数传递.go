package main

import "fmt"

func MyFunc111(temp ...int) {
	for _, data := range temp {
		fmt.Printf("data=%d\n", data)
	}
}

func test(args ...int) {
	// MyFunc111(args...)  //实参一定是args...
	//MyFunc111(args[0:2]...) //args[0]~args[2](不包含数字args【2】)传递过去
	//MyFunc111(args[:2]...)//args[0]~args[2](不包含数字args【2】)传递过去

	MyFunc111(args[2:]...) //args[2]开始（包含args【2】本身，把后面所有的元素都传递过去）
}

func main() {
	test(1, 2, 3, 4, 5, 6)

}

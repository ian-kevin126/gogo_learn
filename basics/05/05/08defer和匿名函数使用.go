package main

import "fmt"

func main() {
	a := 10
	b := 20

	/*defer func(){
		fmt.Printf("a=%d,b=%d\n",a,b)
	}()*/
	defer func(a, b int) {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}(a, b) //(10,20) 代表调用此函数，把参数传递过去，已经先传递参数，只是没有调用，所以这里还是原来的值

	a = 111
	b = 222
	fmt.Printf("外部 a=%d,b=%d\n", a, b)

	/** 打印：
	外部 a=111,b=222
	a=10,b=20
	*/
}

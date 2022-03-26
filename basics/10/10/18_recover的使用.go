package main

import "fmt"

func testa1() {
	fmt.Println("aaaaaaaaaaaaaaaaaa")
}

func testb1(x int) {
	defer func() {
		//recover() //通过revocer捕获错误信息，继续往后执行
		//fmt.Println(recover())
		if err := recover(); err != nil { //打印出产生的异常
			fmt.Println(err)
		}
	}() //别忘了调用此函数

	var a [10]int
	a[x] = 1111 //当x为20 时候，数组越界，产生一个panic，导致程序崩溃

}

func testc1() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa1()
	testb1(20)
	testc1()

}

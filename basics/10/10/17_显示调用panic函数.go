package main

import "fmt"

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaaa")
}

func testb() {
	fmt.Println("bbbbbbbbbbbbbbbbbb")
	//显示调用panic函数，导致程序的中断
	panic("this is a panic test")
	/* a := 10
	b :=0
	i := a/b //当分母为0的时候，产生一个panic。导致程序崩溃
	fmt.Println(i)*/
}

func testc() {
	fmt.Println("cccccccccccccccccc")
}

func main() {
	testa()
	testb()
	testc()

}

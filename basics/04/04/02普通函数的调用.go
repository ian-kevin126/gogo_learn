package main

import "fmt"

func MyFunc() {
	b := "发射子弹"
	fmt.Println("b=", b)
}

func MyFunc01(a string) { //a = "发射子弹" 形式参数
	b := a
	fmt.Println("b=", b)
}

func MyFunc02(a string, b int) {
	fmt.Printf("a=%s,b=%d", a, b)
}

func MyFunc03(a, b int) {
	fmt.Printf("a=%d,b=%d", a, b)
}

func MyFunc04(a, b string, c int) {

}

func MyFunc05(a, b string, c, d int) {

}

func main() {
	MyFunc01("发射子弹")    //实参
	MyFunc01("发射导弹")    //实参
	MyFunc02("发射子弹", 5) //在函数调用的时候，有几个值就传几个，是什么类型就传什么类型的
}

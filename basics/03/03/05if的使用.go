package main

import "fmt"

func main() {
	var a int
	a = 5 // = 赋值
	// a  := 10 //:= 10初始化并赋值给a
	if a == 10 { //==  判断 true fasle
		fmt.Println("a==10")
	}

	//if else 使用
	if a == 10 { //==  判断 true fasle
		fmt.Println("a==10")
	} else {
		fmt.Println("a !=10")
	}

	//if语句进行初始化 然后在判断
	if a := 10; a == 10 {
		fmt.Println("a==10")
	} else {
		fmt.Println("a!=10")
	}

	fmt.Println("a======", a)

	a = 8
	if a == 10 {
		fmt.Println("a==10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a <10")
	} else {
		fmt.Println("这是不可能走到的")
	}

	if a := 8; a == 10 {
		fmt.Println("a==10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a <10")
	} else {
		fmt.Println("这是不可能走到的")
	}

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

}

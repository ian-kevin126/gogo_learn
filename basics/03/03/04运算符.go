package main

import "fmt"

func main() {

	fmt.Println("1+2 结果：", 1+2)

	A := 20
	A++ //++只能放在后面
	A = A + 1
	//++A
	fmt.Println("A++ 的结果", A)

	fmt.Println("5 > 3 结果：", 5 > 3)

	fmt.Println("!(4>3) 结果：", !(4 > 3))

	//&& 与 并且 左边右边都为真，结果才为真，有一个为假 都为假
	fmt.Println("true&&true 结果", true && true)
	fmt.Println("true&&false 结果", true && false)

	//|| 或 左边和右边只要有个一个为真都为真

	fmt.Println("true||false 结果", true || false)
	fmt.Println("false||false 结果", false || false)

	a := 20

	fmt.Println("0<=a && a<=10 结果：", 0 <= a && a <= 10)
}

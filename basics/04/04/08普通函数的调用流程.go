package main

import "fmt"

func funccc(c int) {
	fmt.Println("c=", c)
}

func funccb(b int) {
	funccc(b - 1)
	fmt.Println("b=", b)
}

func funcca(a int) {
	funccb(a - 1)
	fmt.Println("a=", a)
}

func main() { /// 1 2 3
	funcca(3)
	fmt.Println("main")
}

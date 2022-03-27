package main

import "fmt"

func main() {
	var p *int //指针变量
	p = nil
	fmt.Println("p=", p) // p= <nil>

	//*p = 1000 invalid memory address or nil pointer dereference
	//因为p没有任何指向

	var a int
	p = &a
	*p = 1000
	fmt.Println("a=", a) // 1=1000

}

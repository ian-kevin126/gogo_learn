package main

import "fmt"

func test1(x int) {
	result := 1000 / x
	fmt.Println("result=", result)
}

func main() {

	defer fmt.Println("aaaaaaaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbbbbb")
	defer test1(0)
	defer fmt.Println("cccccccccccccccccc")

	/**
	  cccccccccccccccccc
	  bbbbbbbbbbbbbbbbb
	  aaaaaaaaaaaaaaaaa
	  panic: runtime error: integer divide by zero
	*/
}

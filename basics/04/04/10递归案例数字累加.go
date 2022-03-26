package main

import "fmt"

func testSum(i int) int {
	if i == 1 {
		return 1
	}

	return i + testSum(i-1)
}

func testSum1(i int) int {
	if i == 100 {
		return 100
	}

	return i + testSum1(i+1)
}

func main() {
	var sum int
	sum = testSum(100)
	fmt.Println("sum=", sum)

	sum = testSum1(1)
	fmt.Println("sum=", sum)
}

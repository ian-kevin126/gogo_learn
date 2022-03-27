package main

import "fmt"

func swap1(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func main() {
	a, b := 10, 20
	swap1(&a, &b)

	fmt.Printf("main: a= %d,b=%d\n", a, b) // a=20,b=10
}

package main

import "fmt"

func modify(a [5]int) {
	a[1] = 666
	fmt.Println("modify a=", a) // modify a= [1 666 3 4 5]
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	modify(a)
	fmt.Println("main a=", a) // main a= [1 2 3 4 5]

}

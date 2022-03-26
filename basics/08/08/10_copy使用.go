package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	destSlice := []int{7, 7, 7, 7}
	//copy(destSlice,srcSlice)
	copy(srcSlice, destSlice)
	fmt.Println("dst=", srcSlice)
}

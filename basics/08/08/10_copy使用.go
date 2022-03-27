package main

import "fmt"

func main() {
	srcSlice := []int{1, 2}
	destSlice := []int{7, 7, 7, 7}
	//copy(destSlice,srcSlice)
	//fmt.Println("dst=", srcSlice) // dst= [1 2 7 7]
	copy(srcSlice, destSlice)
	fmt.Println("dst=", srcSlice) // dst= [7 7]
}

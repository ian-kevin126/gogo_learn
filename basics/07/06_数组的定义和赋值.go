package main

import "fmt"

func main() {
	//id1 := 1
	//id2 :=2
	//id3 :=3
	var id [50]int //数组，同一类型的集合

	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d]=%d\n", i, id[i])
	}
}

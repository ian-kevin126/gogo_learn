package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	i := 0

	for {
		<-ticker.C // 1séçé´é
		i++
		fmt.Println("i=", i)

		if i == 5 {
			ticker.Stop()
			break
		}
	}

	/*
		i= 1
		i= 2
		i= 3
		i= 4
		i= 5
	*/

}

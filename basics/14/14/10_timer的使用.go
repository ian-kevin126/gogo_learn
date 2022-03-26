package main

import (
	"fmt"
	"time"
)

func main() {
	//验证time.newTimer（）时间到了 只会响应一次
	timer := time.NewTimer(1 * time.Second)

	for {
		<-timer.C
		fmt.Println("时间到")
	}
}

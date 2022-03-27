package main

import (
	"fmt"
	"regexp"
)

func main() {

	buf := "43.14 567 agsdg 1.23 7. 8.9 dddddss 6.66 7.8"
	//正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`) //4.5  \.
	//result := reg.FindAllString(buf,-1)
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", result) // result= [[43.14] [1.23] [8.9] [6.66] [7.8]]
}

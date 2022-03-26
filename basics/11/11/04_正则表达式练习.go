package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a7c aac 888 a9c ac tac" //axxxc

	//1,写一个解析规则
	//reg1 := regexp.MustCompile(`a(.)c`) //(.) 把满足条件放在一组
	//配置数字
	//reg1 := regexp.MustCompile(`a([0-9]+)c`) //a1c a2c a3c
	// reg1 := regexp.MustCompile(`a(\d+)c`) //a1c a2c a3c [0-9]=\d
	//配置字符串 [a-zA-Z0-9_]=\w
	//reg1 := regexp.MustCompile("a[a-zA-Z]c")
	reg1 := regexp.MustCompile(`a\wc`)
	//2.拿规则去配置字符串
	//reslut := reg1.FindAllString(buf,-1)
	reslut := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result=", reslut)
}

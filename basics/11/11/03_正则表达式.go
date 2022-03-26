package main

import (
	"fmt"
	"regexp"
)

func main() {
	//reg := regexp.MustCompile("\\w+") 正则表达式中的\需要转义
	reg := regexp.MustCompile(`^z.*l$`) //^以z开头 以l结尾 .任意字符 *重复零次或者多次

	result := reg.FindAllString("zhangsanl zhangsanl", -1)
	fmt.Printf("%v\n", result)

	reg1 := regexp.MustCompile(`^z(.*)l$`)

	result1 := reg1.FindAllString("zhangsanl", -1)
	fmt.Printf("%v\n", result1)
}

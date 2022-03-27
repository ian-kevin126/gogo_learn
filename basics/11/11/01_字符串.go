package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 字符串转化的函数在strconv中，如下是一些常用的，Append函数可以将整数等转换为字符串，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10) // 以十进制方式追加
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdef")
	str = strconv.AppendQuoteRune(str, '单')

	fmt.Println(string(str)) // 4567false"abcdef"'单'
}

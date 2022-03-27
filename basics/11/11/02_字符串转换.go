package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换成字符串后追加到字节数组中
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数要为追加是数，第三个指定10进制的方式追加
	slice = strconv.AppendInt(slice, 1234, 16)
	slice = strconv.AppendQuote(slice, "abcdefg")

	fmt.Println("slice=", string(slice)) // slice= true4d2"abcdefg"

	//把其他类型转换成字符串
	var str string
	str = strconv.FormatBool(true)

	str = strconv.FormatInt(10, 8)
	fmt.Println("str=", str) // str= 12

	str = strconv.FormatFloat(3.14, 'f', -1, 64)
	fmt.Println("str=", str) // str= 3.14

	// 常用 把整形转换成字符串
	str = strconv.Itoa(6666)
	fmt.Println("str=", str) // str= 6666

	//把字符串转换成其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("true")

	if err == nil {
		fmt.Println("flag=", flag) // flag= true
	} else {
		fmt.Println("err=", err)
	}

	// 把字符串转换成整形
	a, err := strconv.Atoi("aaa")
	if err == nil {
		fmt.Println("a=", a)
	} else {
		fmt.Println("err=", err) // err= strconv.Atoi: parsing "aaa": invalid syntax
	}

}

package main

import (
	"fmt"
	"strings"
)

func main() {
	//"hellogo" 中是否包含hello，包含返回true,不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "_")
	fmt.Println("buf=", buf)

	//Index 查找字段位置
	fmt.Println(strings.Index("abcdhello", "hello")) //4
	fmt.Println(strings.Index("abcdhello", "ddddd")) //-1

	//Repeat
	buf = strings.Repeat("go", 5)
	fmt.Println("buf=", buf)

	//split
	buf = "abc_hello_mike_go"
	s2 := strings.Split(buf, "_")
	fmt.Println("s2=", s2)

	//trim
	buf = strings.Trim(" are you aoka    ", " a") //去掉俩边头的空格指定的字符
	fmt.Printf("buf=#%s#\n", buf)

	//Fields 去掉空格，把元素放入切片中
	s3 := strings.Fields(" are you ok?   ")
	fmt.Println("s3=", s3)

	for i, data := range s3 {
		fmt.Println(i, data)
	}

}

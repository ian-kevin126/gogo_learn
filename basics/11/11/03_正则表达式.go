package main

import (
	"fmt"
	"regexp"
)

/**
. : 匹配除了换行符以外的任意字符
\w：匹配字母或数字或下划线或汉字
\s：匹配任意的空白字符
\d：匹配数字
\b：匹配单词的开始或结束
^ ：匹配字符的开始
& ：匹配字符的结束
* ：重复0或更多次
+ ：重复1次或更多次
? ：重复0次或1次
{n} ：重复n次
{n,} ：重复n次或更多次
{n,m} ：重复n次到m次
*/

func main() {
	//reg := regexp.MustCompile("\\w+") 正则表达式中的\需要转义
	reg := regexp.MustCompile(`^z.*l$`) //^以z开头 以l结尾 .任意字符 *重复零次或者多次

	result := reg.FindAllString("zhangsanl zhangsanl", -1)
	fmt.Printf("%v\n", result) // [zhangsanl zhangsanl]

	reg1 := regexp.MustCompile(`^z(.*)l$`)

	result1 := reg1.FindAllString("zhangsanl", -1)
	fmt.Printf("%v\n", result1) // [zhangsanl]
}

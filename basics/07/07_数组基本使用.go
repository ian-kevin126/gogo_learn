package main

import (
	"fmt"
)

func main() {
	//定义一个数组 [10]int和 [5]int 是不同的类型
	//[数字]，这个数字代表的是数组的元素个个数
	var a [10]int
	var b [5]int
	fmt.Printf("len(a)=%d,len(b)=%d\n", len(a), len(b))

	//注意点 定义数组，指定的数组的个数必须是常量 non-constant array bound n

	//n:=10
	//var c [n]int
	//fmt.Printf("len(c)=%d\n",len(c))
	// 操作数组元素，从0开始，到len()-1,不对称元素，这个数字叫下标,下标是可以使用变量的
	a[0] = 1
	i := 1
	a[i] = 10

	//赋值，每个元素
	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	//打印，第一个返回下标 ，第二个是返回元素对应的值
	for i, data := range a {
		fmt.Printf("a[%d]=%d\n", i, data)
	}

}

package main

import "fmt"

/**
复合类型

类型       名称     长度     默认值   说明
pointer   指针              nil
array     数组              0
slice     切片              nil     引用类型
map       字典              nil     引用类型
struct    结构体
*/

func main() {

	var a int = 10
	// 每个变量有2层含义：变量的内存 变量地址
	fmt.Printf("a=%d\n", a)     // 变量的内存  a=10
	fmt.Printf("&a = %v\n", &a) // 变量的地址  &a = 0xc0000bc008

	// 保存地址必须要用指针（地址）
	var p *int                                  // 定义一个变量p，类型为 *int类型
	p = &a                                      // 指针变量指向谁，就把谁的地址赋值给指针变量
	fmt.Printf("p=%v,&a=%v,*p=%v\n", p, &a, *p) // p=0xc0000bc008,&a=0xc0000bc008,*p=10

	*p = 1000
	fmt.Printf("p=%v,&a=%v,*p=%v\n", p, &a, *p) // p=0xc0000bc008,&a=0xc0000bc008,*p=1000

}

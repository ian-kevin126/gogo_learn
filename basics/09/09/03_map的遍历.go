package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "jack", 3: "tom"}
	//第一个返回值为key 第二个返回值为value 遍历结构是无序的
	for key, value := range m {
		fmt.Printf("%d======>%s\n", key, value)
	}

	//如何判断一个key是否存在，
	//第一个返回值为key的所对应的value，第二个返回值为key是否存在的条件，如果存在ok为true
	//value,ok := m[1]
	if value, ok := m[1]; ok == true {
		fmt.Println("m[1]=", value)
	} else {
		fmt.Println("key 不存在")
	}

	delete(m, 2) //删除key为2的内容
	fmt.Println("m=", m)

}

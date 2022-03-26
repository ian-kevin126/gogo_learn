package main

import "fmt"

func main() {
	m1 := map[int]string{1: "mike", 2: "jack"}
	fmt.Println("m1=", m1)

	m1[1] = "marry" //如果key存在，修改内容
	m1[3] = "tom"   //如果没有key 追加内容，切片append
	fmt.Println("m1=", m1)
}

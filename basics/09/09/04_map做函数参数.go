package main

import "fmt"

func test(m map[int]string) { //map 和我们切片一样 他们是引用类型
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "jack", 2: "mike", 3: "marry"}
	fmt.Println("m=", m) // m= map[1:jack 2:mike 3:marry]

	test(m)

	fmt.Println("m=", m) // m= map[2:mike 3:marry]

}

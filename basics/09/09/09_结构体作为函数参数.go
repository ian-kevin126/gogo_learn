package main

import "fmt"

type Student4 struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func test01(s Student4) {
	s.id = 666
	fmt.Println("test01:", s) // test01: {666 mike 109 18 bj}
}

func test02(s *Student4) {
	s.id = 666
	fmt.Println("test01:", s) // test01: &{666 mike 109 18 bj}
}

func main() {
	s := Student4{1, "mike", 'm', 18, "bj"}
	test01(s)               //值传递
	fmt.Println("main:", s) // main: {1 mike 109 18 bj}

	test02(&s)              //地址传递（引用传递），形参可以改变实参
	fmt.Println("main:", s) // main: {666 mike 109 18 bj}

}

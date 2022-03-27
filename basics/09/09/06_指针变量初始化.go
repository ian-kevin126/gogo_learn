package main

import "fmt"

type Person struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//p1 存放是地址Person
	var p1 *Person = &Person{1, "mike", 'm', 18, "bj"}
	fmt.Printf("p1=%v\n", p1) // p1=&{1 mike 109 18 bj}

	p2 := &Person{name: "mike", addr: "bj"}
	fmt.Printf("p2 type is => %T\n", p2) // p2 type is => *main.Person
	fmt.Println("p2=", p2)               // p2= &{0 mike 0 0 bj}

}

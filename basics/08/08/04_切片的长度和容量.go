package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} //左包含右不包含
	/*s := a[0:3:5]
	fmt.Println("s=",s)
	fmt.Println("len(s)=",len(s))
	fmt.Println("cap(s)=",cap(s))*/

	s1 := a[1:3:5]
	fmt.Println("s1=", s1)
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))
	fmt.Printf("%p\n", &s1)
	fmt.Println("a=", a)

	s1 = append(s1, 1)
	fmt.Println("s1=", s1)
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))
	fmt.Printf("%p\n", &s1)
	fmt.Println("a=", a)

	s1 = append(s1, 1)
	fmt.Println("s1=", s1)
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))
	fmt.Printf("%p\n", &s1)
	fmt.Println("a=", a)

	s1 = append(s1, 1)
	fmt.Println("s1=", s1)
	fmt.Println("len(s1)=", len(s1))
	fmt.Println("cap(s1)=", cap(s1))
	fmt.Printf("%p\n", &s1)
	fmt.Println("a=", a)

	s1[0] = 9999
	fmt.Println("a=", a)

}

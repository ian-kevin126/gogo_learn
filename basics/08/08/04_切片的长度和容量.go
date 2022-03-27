package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} //左包含右不包含
	/*s := a[0:3:5]
	fmt.Println("s=",s)
	fmt.Println("len(s)=",len(s))
	fmt.Println("cap(s)=",cap(s))*/
	fmt.Printf("a=%p\n", &a) // a=0xc0000aa060

	s1 := a[1:3:5]
	fmt.Println("s1=", s1)           // s1= [2 3]
	fmt.Println("len(s1)=", len(s1)) // len(s1)=2
	fmt.Println("cap(s1)=", cap(s1)) // cap(s1)=4
	fmt.Printf("%p\n", &s1)          // 0xc0000a4018
	fmt.Println("a=", a)             // a= [1 2 3 4 5]
	fmt.Printf("a=%p\n", &a)         // a=0xc0000aa060

	s1 = append(s1, 1)
	fmt.Println("s1=", s1)           // s1= [2 3 1]
	fmt.Println("len(s1)=", len(s1)) // len(s1)= 3
	fmt.Println("cap(s1)=", cap(s1)) // cap(s1)= 4
	fmt.Printf("%p\n", &s1)          // 0xc0000a4018
	fmt.Println("a=", a)             // a= [1 2 3 1 5]

	s1 = append(s1, 1)
	fmt.Println("s1=", s1)           // s1= [2 3 1 1]
	fmt.Println("len(s1)=", len(s1)) // len(s1)= 4
	fmt.Println("cap(s1)=", cap(s1)) // cap(s1)= 4
	fmt.Printf("%p\n", &s1)          // 0xc0000a4018
	fmt.Println("a=", a)             // a= [1 2 3 1 1]

	s1 = append(s1, 1)
	fmt.Println("s1=", s1)           // s1= [2 3 1 1 1]
	fmt.Println("len(s1)=", len(s1)) // len(s1)= 5
	fmt.Println("cap(s1)=", cap(s1)) // cap(s1)= 8
	fmt.Printf("%p\n", &s1)          // 0xc0000a4018
	fmt.Println("a=", a)             // a= [1 2 3 1 1]

	s1[0] = 9999
	fmt.Println("a=", a) // a= [1 2 3 1 1]

}

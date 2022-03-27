package main

import (
	"fmt"
)

type Worker1 struct {
	name string
	sex  byte
	age  int
}

// SetWorkerInfo 接收者为普通变量，非指针，值语义，一份拷贝
func (w Worker1) SetWorkerInfo(n string, s byte, a int) {
	w.name = n
	w.sex = s
	w.age = a
	fmt.Println("w=", w)
	fmt.Printf("SetWorkerInfo &p=%p\n", &w)
}

// SetWorkerPoint 引用传递
func (w *Worker1) SetWorkerPoint(n string, s byte, a int) {
	w.name = n
	w.sex = s
	w.age = a
	fmt.Printf("SetWorkerPoint &p=%p\n", w)
}

func main() {

	w1 := Worker1{"tom", 'm', 22}
	fmt.Printf("&w1=%p\n", &w1) // &w1=0xc0000b4000
	//w1.SetWorkerInfo("marry",'f',11)
	//fmt.Printf("&w1=%p\n",&w1)

	//(&w1).SetWorkerPoint("marry",'f',11)
	w1.SetWorkerPoint("marry", 'f', 11) // SetWorkerPoint &p=0xc0000b4000
	fmt.Printf("&w1=%p\n", &w1)         // &w1=0xc0000b4000
	fmt.Println(w1)                     // {marry 102 11}

}

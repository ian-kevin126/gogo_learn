package main

import "fmt"

type Worker2 struct {
	name string
	sex  byte
	age  int
}

//接收者为普通变量，非指针，值语义，一份拷贝
func (w Worker2) SetWorkerInfo() {
	fmt.Printf("SetWorkerInfo &p=%p\n", &w)
}

//引用传递
func (w *Worker2) SetWorkerPoint() {
	fmt.Printf("SetWorkerPoint &p=%p\n", w)
}

func main() {
	//结构体变量是一个指针类型的变量，它能够调用方法，这些可以调用的方法 简称方法集
	/*w := &Worker2{"mike",'m',19}
	w.SetWorkerPoint()
	(*w).SetWorkerPoint() // 先把（*w）转换成w在调用 定价与上面
	//先把指针w，转换成*w在调用
	//(*w).SetWorkerInfo()
	w.SetWorkerInfo()*/

	w := Worker2{"mike", 'm', 19}
	w.SetWorkerPoint() //w--->(&p) SetWorkerPoint &p=0xc00012a000
	w.SetWorkerInfo()  // SetWorkerInfo &p=0xc00012a020

}

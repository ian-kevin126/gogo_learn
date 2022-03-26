package main

import "fmt"

type Worker struct {
	name string //名字
	sex  byte   //性别
	age  int    //年龄
}

func PrintInfo1(a Worker) {
	fmt.Println("a=", a)
}

func (a Worker) PrintInfo() {
	fmt.Println("a=", a)
}

/*func (p Worker) SetInfo(n string,s byte,a int) {
	p.name = n
	p.sex = s
	p.age =a
	fmt.Println("SetInfo p=",p)
}*/

func (p *Worker) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("SetInfo p=", p)
}

func main() {

	// w := Worker{"zhangsan",'m',20}
	// w.PrintInfo()
	//PrintInfo1(w) //面向过程
	//  w.PrintInfo()

	//定义一个结构体变量
	var p Worker
	//(&p).SetInfo("tom",'f',33)
	p.SetInfo("tom", 'f', 33)
	p.PrintInfo()
}

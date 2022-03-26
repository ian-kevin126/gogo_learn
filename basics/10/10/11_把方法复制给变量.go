package main

import "fmt"

type Person7 struct {
	name string //姓名
	sex  byte   //性别，字符类型
	age  int    //年龄
}

func (p Person7) SetInfoValue() {
	fmt.Printf("SetInfoValue:%p,%v\n", &p, p)
}

func (p *Person7) SetInfoPoint() {
	fmt.Printf("SetInfoPoint:%p,%v\n", p, p)
}

func main() {
	p := Person7{"mike", 'm', 18}
	fmt.Printf("main:%p,%v\n", &p, p)

	p.SetInfoPoint() //传统的调用方式

	pFunc := p.SetInfoPoint //这个是方法值，调用函数时，无需传递接收者，
	pFunc()

	pSetinfo := p.SetInfoValue
	pSetinfo() //这个就等价于p.SetInfoValue()

}

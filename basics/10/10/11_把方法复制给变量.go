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
	fmt.Printf("main:%p,%v\n", &p, p) // main:0xc000136000,{mike 109 18}

	p.SetInfoPoint() // 传统的调用方式 SetInfoPoint:0xc000136000,&{mike 109 18}

	pFunc := p.SetInfoPoint // 这个是方法值，调用函数时，无需传递接收者，
	pFunc()                 // SetInfoPoint:0xc000136000,&{mike 109 18}

	pSetInfo := p.SetInfoValue
	pSetInfo() // 这个就等价于p.SetInfoValue() // SetInfoPoint:0xc000136000,&{mike 109 18}

}

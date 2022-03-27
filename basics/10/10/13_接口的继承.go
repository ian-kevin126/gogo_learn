package main

import "fmt"

type Humaner01 interface {
	sayHi()
}

type Personer interface {
	Humaner01 //匿名字段，继承sayHi()
	sing(lrc string)
}

type Student10 struct {
	name string
	id   int
}

func (temp *Student10) sayHi() {
	fmt.Printf("Student10[%s,%d] sayHi\n", temp.name, temp.id)
}

func (temp *Student10) sing(lrc string) {
	fmt.Println("student10在唱：", lrc)
}

func main() {
	//定义一个接口类型的变量
	var i Personer
	s := &Student10{"mike", 7777}
	i = s
	i.sayHi()       // Student10[mike,7777] sayHi
	i.sing("人民解放歌") // student10在唱： 人民解放歌

	var h Humaner01
	h = s
	h.sayHi() // Student10[mike,7777] sayHi

	h = i     // 子接口可以赋值给父接口
	h.sayHi() // Student10[mike,7777] sayHi

	//i = h // 父接口定义变量不能给子接口

}

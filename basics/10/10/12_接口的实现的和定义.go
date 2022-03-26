package main

import "fmt"

//定义一个接口类型
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型（自定义类型）实现
	sayHi()
}

type Student9 struct {
	name string
	id   int
}

func (temp *Student9) sayHi() {
	fmt.Printf("Student9[%s,%d] sayHi\n", temp.name, temp.id)
}

type Teacher struct {
	addr string
	id   int
}

func (temp *Teacher) sayHi() {
	fmt.Printf("Teacher[%s,%d] sayHi\n", temp.addr, temp.id)
}

type Worker4 struct {
	id int
}

func (temp *Worker4) sayHi() {
	fmt.Printf("Worker4[%d] sayHi\n", temp.id)
}

//定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以表现有不同的形式，多态
func WhoSayHi(i Humaner) { //接口作为参数
	i.sayHi()
}

func main() {

	//定义一个接口的类型
	//var i Humaner
	// 只有实现了此接口的方法的类型，那个这个类型的变量才能赋值给接口变量
	s := &Student9{"mike", 7777}

	//i =s
	//i.sayHi()

	t := &Teacher{"bj", 9999}
	//i = t

	//i.sayHi()

	w := &Worker4{8888}
	//i= w
	//i.sayHi()

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(w)

	//创建一个切片,切片放入是接口
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = w

	for _, i := range x {
		i.sayHi()
	}

}

package main

import "fmt"

/*
在go语言中，接口（Interface）是一个自定义的类型，接口类型具体描述了一系列的方法集合。集合类型是一种抽象的类型，
它不会暴露出所有的对象的内部值的结构和这个对象支持的基础操作集合，它们只展现它们自己的反复噶。因此接口不能实例化
，必须通过自雷来实例化。

go通过接口实现鸭子类型（Duck Typing）："当看到一直鸟走起来像鸭子，游泳像鸭子，叫起来像鸭子，那么这只鸟就可以被看做是鸭子类型。"
即就是，我们不关心对象是什么类型，到底是不是真的鸭子，我们只关心行为。

- 接口命名习惯以 "er" 结尾
- 接口只有方法声明，没有方法实现，没有数据字段
- 接口可以匿名嵌入其他接口，或嵌入到接口提中

注意点：接口用来定义行为类型，这些被定义的行为不由借口自己直接实现，而是通过方法由用户定义的类型实现。
*/

// Humaner 定义一个接口类型
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

// WhoSayHi 定义一个普通函数，函数的参数为接口类型
//只有一个函数，可以表现有不同的形式，多态
func WhoSayHi(i Humaner) { //接口作为参数
	i.sayHi()
}

func main() {

	//定义一个接口的类型
	/*
		var i Humaner

		// 只有实现了此接口的方法的类型，那个这个类型的变量才能赋值给接口变量
		s := &Student9{"mike", 7777}
		i = s
		i.sayHi() // Student9[mike,7777] sayHi
	*/

	s := &Student9{"mike", 7777}

	t := &Teacher{"bj", 9999}
	//i = t

	//i.sayHi()

	w := &Worker4{8888}
	//i= w
	//i.sayHi()

	// 多态：同样一个函数，不同的实现
	WhoSayHi(s) // Student9[mike,7777] sayHi
	WhoSayHi(t) // Teacher[bj,9999] sayHi
	WhoSayHi(w) // Worker4[8888] sayHi

	//创建一个切片,切片放入是接口
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = w

	for _, i := range x {
		i.sayHi()
	}
	/*
		Student9[mike,7777] sayHi
		Teacher[bj,9999] sayHi
		Worker4[8888] sayHi
	*/

}

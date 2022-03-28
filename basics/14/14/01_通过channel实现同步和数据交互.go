package main

import (
	"fmt"
	"time"
)

/*
goroutine运行在相同的地址空间，因此访问共享内存必须做好同步。goroutine奉行通过通信来共享内存。
而不是共享内存来通信。引用类型channel是CSP模式的具体实现，用于多个goroutine通信，其内部实现了同步
，确保并发安全。

和map类似，channel也是一个对应make创建的底层数据结构的引用。

当我们复制一个channel或用于函数参数传递的时候，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用
同一个channel对象。和其他的引用类型一样，channel的零值也是nil。

顶一个channel时，也需要定义发送到channel的值的类型，channel可以使用内置的make()函数来创建：

make(chan Type, capacity)

当capacity=0时，channel是无缓冲阻塞读写的（有人写，就要有人读），当capacity>0时，channel有缓冲，是非阻塞的，
直到写满capacity个元素之后才会发生阻塞写入。

——————————————————————————————————————————————————————
<====   1   2   3   4   5   6   7   8   9   10   <====
——————————————————————————————————————————————————————

channel操作：

channel <- value  // 发送value到channel
<- channel  // 接收并将其丢弃
x := <- channel // 从channel中接收数据，并将其赋值给x
x, ok := <- channel // 功能同上，同时检查通道是否已经关闭或者是否为空

默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得goroutine同步变得更加简单，而不需要显式地lock。


*/

func main() {
	//创建channel
	ch := make(chan string)
	defer fmt.Println("主协程也结束")

	go func() {
		defer fmt.Println("子协程结束")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程i=", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程，子协程工作完毕"
	}()

	str := <-ch //没有数据，阻塞
	fmt.Println("str=", str)

	/*
		子协程i= 0
		子协程i= 1
		子协程结束
		str= 我是子协程，子协程工作完毕
		主协程也结束
	*/
}

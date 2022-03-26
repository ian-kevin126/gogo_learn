package main

func main() {
	//创建一个管道
	ch := make(chan int)
	//双向channel能隐式的转换为单向的channel
	var writeCh chan<- int = ch //只能写 不能读
	var readch <-chan int = ch  //只能读 不能写

	writeCh <- 666
	//<- writeCh  //invalid operation: <-writeCh (receive from send-only type chan<- int)

	<-readch //读
	//readch <- 666 //invalid operation: readch <- 666 (send to receive-only type <-chan int)

	//单向的无法转换成双向的
	//var ch2 chan int =writeCh //cannot use writeCh (type chan<- int) as type chan int in assignment
}

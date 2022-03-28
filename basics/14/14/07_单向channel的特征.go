package main

/*
var ch1 chan int  // 正常channel，可读可写
var ch2 chan<- float64 // 单向channel，可写不可读
var ch3 <-chan int // 单向channel，可读不可写
*/
func main() {
	//创建一个管道
	ch := make(chan int)

	// 双向channel能隐式的转换为单向的channel
	var writeCh chan<- int = ch //只能写 不能读
	var readCh <-chan int = ch  //只能读 不能写

	writeCh <- 666
	//<- writeCh  //invalid operation: <-writeCh (receive from send-only type chan<- int)

	<-readCh //读
	//readCh <- 666 //invalid operation: readCh <- 666 (send to receive-only type <-chan int)

	//单向的无法转换成双向的
	//var ch2 chan int =writeCh //cannot use writeCh (type chan<- int) as type chan int in assignment
}

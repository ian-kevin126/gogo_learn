package main

import (
	"fmt"
)

/*
注意点：
- channel不像文件一样需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式地结束range循环之类的，才关闭channel
- 关闭channel后，无法向channel再发送数据了（引发panic错误后导致接收立即返回零值）
- 关闭channel后，可以继续向channel接收数据
- 对于nil channel，无论接收和发送都会被阻塞
*/

func main() {
	ch := make(chan int) //创建一个无缓存的channel

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //往通道写数据
			fmt.Printf("写数据%d\n", i)
		}
		// 如果不手动关闭就会死锁：fatal error: all goroutines are asleep - deadlock!
		close(ch) //不需要写数据的时候 我们关闭channel
		// ch <-666  //send on closed channel 关闭channel后无法在发送数据
	}()

	for {
		// 如果ok为true,说明管道没有关闭
		if num, ok := <-ch; ok {
			fmt.Println("num=", num)
		} else { //管道关闭
			break
		}
	}

	/*
		写数据0
		num= 0
		num= 1
		写数据1
		写数据2
		num= 2
		num= 3
		写数据3
		写数据4
		num= 4
	*/
}

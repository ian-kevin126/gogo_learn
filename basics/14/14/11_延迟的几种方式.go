package main

import (
	"fmt"
	"time"
)

//time.NewTimer(2*time.Second).C
func main() {
	<-time.After(2 * time.Second) //定时2s 阻塞2s 2s后产生一个事假往channel写内容，和第一种一样
	fmt.Println("时间到")
}

/*
func main()  {
	time.Sleep(2*time.Second)
	fmt.Println("时间到")
}
*/

/*
func main(){
	//延时2秒方法
	timer := time.NewTimer(2*time.Second).C

	<-timer.C
	fmt.Println("时间到")
}
*/

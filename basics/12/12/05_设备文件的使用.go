package main

import (
	"fmt"
	//"os"
	"os"
)

func main() {
	//os.Stdout.Close()  //关闭后，无法输出  os.Stdout标准设备文件 默认是打开的，用户之间使用
	fmt.Println("are you ok?") //往标准输入设备（屏幕）写内容
	//os.Stdout.
	os.Stdout.WriteString("are you ok?\n")

	//os.Stdin
	var a int
	fmt.Println("请输入一个a")
	fmt.Scan(&a)
	fmt.Println("a=", a)

}

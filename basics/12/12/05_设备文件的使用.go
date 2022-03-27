package main

import (
	"fmt"
	//"os"
	"os"
)

/*
数据存储的位置：
- 内存（memory）：缺点是断电之后数据会丢失  优点是读取速度快
- 磁盘（disk）：缺点是读取速度慢 优点是数据永久存储（除非手动清除）
*/

func main() {
	//os.Stdout.Close()  //关闭后，无法输出  os.Stdout 标准设备文件 默认是打开的，用户之间使用

	fmt.Println("are you ok?") // 往标准输入设备（屏幕）写内容
	/*
		实现原理：os.Stdout：
		func Println(a ...interface{}) (n int, err error) {
			return Fprintln(os.Stdout, a...)
		}
	*/

	// 等价的语句
	os.Stdout.WriteString("are you ok?\n")

	//os.Stdin
	var a int
	fmt.Println("请输入一个a")
	fmt.Scan(&a)
	/* 实现原理：
	func Scan(a ...interface{}) (n int, err error) {
		return Fscan(os.Stdin, a...)
	}
	*/
	fmt.Println("a=", a)
}

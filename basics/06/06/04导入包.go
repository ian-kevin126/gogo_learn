package main

//_ 忽略此包
import _ "fmt"

//给导入包起一个别名
import io "fmt"

//无须通过包名调用，不建议这样操作
/*import  "fmt"
import  "os"*/

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("aaaa")
	io.Printf("aaaaa")
	fmt.Printf("bbbbbb")
	fmt.Println(os.Args)
}

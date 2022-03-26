package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//文件的写入
func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err =", err)
		return
	}
	//使用完毕，关闭文件
	defer f.Close()

	var buf string
	for i := 0; i < 10; i++ {
		//"i=%d\n" 这个字符串存储到buf里面
		buf = fmt.Sprintf("i=%d\n", i)
		n, err := f.WriteString(buf)
		if err != nil {
			fmt.Println("err=", err)
			return
		}
		fmt.Println("n=", n)
	}

}

//文件读取
func ReadFile(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*2) //2k大小
	n, err := f.Read(buf)
	if err != nil && err != io.EOF { //文件出错同时文件没有到结尾
		fmt.Println("err=", err)
		return
	}
	fmt.Println("buf=", string(buf[:n]))
}

//每次读取一行
func ReadFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	//关闭文件
	defer f.Close()
	//新建一个缓冲区。把内容先放着缓冲里面
	r := bufio.NewReader(f)
	for {
		//遇到\n结束读取，但是\n也进入缓冲区
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err=", err)
		}
		fmt.Printf("buf=#%s#\n", string(buf))

	}
}

func main() {
	path := "./demo.txt"
	//WriteFile(path)
	//ReadFileLine(path)
	ReadFile(path)

}

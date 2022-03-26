package main

import (
	"fmt"
	"net/http"
)

//w 给客户端回复数据
//r 读取客户端的发送数据请求
func HandConn(w http.ResponseWriter, r *http.Request) {

	fmt.Println("r.Method=", r.Method)
	fmt.Println("r.URL=", r.URL)
	fmt.Println("r.HEAd=", r.Header)
	fmt.Println("r.Body=", r.Body)

	w.Write([]byte("<h1>hello world</h1>"))

}

func main() {
	//注册处理函数，用户连接，自动调用指定的处理函数
	http.HandleFunc("/", HandConn)

	//监听绑定
	http.ListenAndServe("127.0.0.1:8000", nil)
}

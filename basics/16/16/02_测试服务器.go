package main

import (
	"fmt"
	"net/http"
)

//如果匹配成功进入业务逻辑处理
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world") //fmt.Fprintln往流量器写内容
}

//如果匹配成功进入业务逻辑处理
func goHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gogogo") //fmt.Fprintln往流量器写内容
}

func main() {
	http.HandleFunc("/go", goHandler)
	http.HandleFunc("/hello", myHandler)
	http.ListenAndServe(":8000", nil)
}

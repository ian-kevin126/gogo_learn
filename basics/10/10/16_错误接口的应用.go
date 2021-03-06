package main

import (
	"errors"
	"fmt"
)

func MyDiv(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不为零") // fmt.Errorf("%s","分母不为零")
	} else {
		result = a / b
	}
	return
}

func main() {
	result, err := MyDiv(10, 0)
	if err != nil {
		fmt.Println("err", err) // err 分母不为零
	} else {
		fmt.Println("result=", result)
	}
}

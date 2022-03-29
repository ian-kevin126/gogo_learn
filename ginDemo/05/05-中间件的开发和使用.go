package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/hello", func(context *gin.Context) {
		fullPath := context.FullPath()
		fmt.Println("请求路径：" + fullPath)

		context.Writer.Write([]byte(fullPath))
	})

	engine.Run()
}

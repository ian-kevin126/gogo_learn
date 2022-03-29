package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// 路由组
	routerGroup := engine.Group("/user")

	routerGroup.POST("/register", registerHandler)
	routerGroup.POST("/login", loginHandler)
	routerGroup.DELETE("/:id", deleteHandler)

	engine.Run()
}

func registerHandler(context *gin.Context) {
	fullPath := context.FullPath()
	fmt.Println("请求路径：" + fullPath)

	context.Writer.Write([]byte(fullPath))
}

func loginHandler(context *gin.Context) {
	fullPath := context.FullPath()
	fmt.Println("请求路径：" + fullPath)

	context.Writer.Write([]byte(fullPath))
}

func deleteHandler(context *gin.Context) {
	fullPath := context.FullPath()
	fmt.Println("请求路径：" + fullPath)

	userID := context.Param("id")

	context.Writer.Write([]byte(fullPath + userID))
}

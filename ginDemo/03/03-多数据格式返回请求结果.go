package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func main() {
	engine := gin.Default()

	engine.GET("/helloByte", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		context.Writer.Write([]byte(context.FullPath()))
	})

	engine.GET("/helloString", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		// 也可以直接使用WriteString方法
		context.Writer.WriteString(context.FullPath())
	})

	// 将map类型的数据转化成json数据格式
	engine.GET("/helloJson", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "ok",
			"data":    context.FullPath(),
		})
	})

	// 将结构体类型的数据转化成json数据格式
	engine.GET("/helloStruct", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		resp := Response{Code: 1, Message: "ok", Data: context.FullPath()}

		// 注意这里传的是引用
		context.JSON(200, &resp)
	})

	// 解析html
	// 首先要设置html目录
	engine.LoadHTMLGlob("./html/*")
	// 设置静态资源文件的路径（图片），将图片放置的文件映射为 /img，然后再html文件中就可以访问到assets中的图片了
	engine.Static("/img", "./assets")

	engine.GET("/helloHtml", func(context *gin.Context) {
		fullPath := context.FullPath()
		fmt.Println("请求路径：" + fullPath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
			"title":    "Gin hello",
		})
	})

	engine.Run()
}

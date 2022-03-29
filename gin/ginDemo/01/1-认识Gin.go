package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

func main() {
	// 在Gin框架中，engine被定义为一个结构体，engine代表gin框架的一个结构体定义，其中包含了路由、中间件、页面渲染接口、框架配置等相关的内容。
	engine := gin.Default()

	/*
		还可以使用New()创建engine，区别在于：gin.Default()其实也是使用New()创建，但是会默认使用Logger和Recovery中间件。
		engine1 := gin.New()

		- Logger是负责进行打印并输出日志的中间件，方便开发者进行程序调试；
		- Recovery中间件的作用是如果程序执行过程中遇到panic中断了服务，则Recovery会恢复程序执行，并返回服务器500内部错误。
		  通常情况下，我们使用默认的gin.Default()创建engine实例。
	*/

	/*
		http协议中一共定义了8种方法或者称之为类型来表明对请求网络资源的不同的操作方式。分别是：
		OPTIONS, HEAD, GET, POST, PUT, DELETE, TRACE, CONNECT，虽然有8种，但实际开发中常用的就Get、Post、Delete等几种。

		通用处理：engine中可以直接进行http请求的处理，在engine中使用handle方法进行http请求的处理。handle方法包含三个参数，具体如下：

		func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...handlers){}

		- httpMethod：第一个参数表示要处理的HTTP请求类型
		- relativePath：第二个参数表示要解析的接口，有开发者定义
		- handlers：第三个参数是处理对应的请求的代码定义
	*/

	// 第一种处理方式
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println("请求路径: ", context.FullPath())
		context.Writer.Write([]byte("Hello gin\n"))
	})

	// 处理POST请求
	engine.POST("/new", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		username, exist := context.GetPostForm("username")
		if exist {
			fmt.Println("username", username)
		}

		password, exists := context.GetPostForm("username")
		if exists {
			fmt.Println("password", password)
		}

		context.Writer.Write([]byte("hello, " + username))
	})

	// 处理DELETE请求
	engine.DELETE("/user/:id", func(context *gin.Context) {
		// 获取id
		userID := context.Param("id")
		fmt.Println(userID)
		context.Writer.Write([]byte("delete user: " + userID))
	})

	// 第二种处理方式
	engine.Handle("GET", "/test", func(context *gin.Context) {
		// 获取请求接口
		fmt.Println(context.FullPath())
		// 获取字符串参数
		// name := context.Query("name")
		// DefaultQuery()与Query的区别在于，DefaultQuery有两个参数，第二个是默认参数
		name := context.DefaultQuery("name", "hello")
		fmt.Println("name: ", name)

		// 输出
		context.Writer.Write([]byte("hello, " + name))
	})

	// 处理POST请求
	engine.Handle("POST", "/login", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		username := context.PostForm("username")
		password := context.PostForm("password")

		fmt.Println("username - password", username+password)

		context.Writer.Write([]byte(username + "登录"))
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

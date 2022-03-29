package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/** 中间件的开发和使用
自定义中间件，有两条标准：
- func函数
- 返回值类型是HandleFunc
*/

// RequestInfos 自定义一个中间件：打印接口请求信息
func RequestInfos() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method

		fmt.Println("请求path: ", path)
		fmt.Println("请求method: ", method)
		fmt.Println("状态码1：", context.Writer.Status()) // 状态码1： 200
		/*
			context.Next函数：

			在上面自定义的中间件RequestInfos中，打印了请求的path和method，接着去执行了正常的业务处理函数。如果我们
			想输出业务处理结果的信息，该如何实现呢？答案是使用context.Next函数。

			context.Next函数可以将中间件代码的执行顺序一分为二，Next函数调用之前的代码在请求处理之前，当程序执行到context.Next
			函数时，会中断向下执行，转而先去执行具体的业务逻辑，执行完业务逻辑处理函数之后，程序再次回到context.Next函数处，继续
			执行中间件后面的代码：
		*/
		context.Next()
		// 请求完成之后，拿到请求结果之后，再回到这里执行这一句
		fmt.Println("状态码2：", context.Writer.Status()) // 状态码2： 404
	}
}

func main() {
	engine := gin.Default()

	// 自定义中间件使用
	engine.Use(RequestInfos())

	engine.GET("/query", func(context *gin.Context) {
		context.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    1,
			"message": context.FullPath(),
		})
	})

	// 单独为某一个接口使用中间件
	engine.GET("/query1", RequestInfos(), func(context *gin.Context) {
		fullPath := context.FullPath()
		fmt.Println("请求路径：" + fullPath)

		context.Writer.Write([]byte(fullPath))
	})

	engine.Run()
}

## 1，HTTP请求和参数解析
```go
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

```

## 2，请求参数绑定与多数据处理
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
)

/*
像上面那种只有两个字段（username和password）的表单数据进行提交时，可以使用context.PostForm和context.GetPostForm获取，但是如果
表单数据比较多的时候，使用这种方式一次获取一个表单数据，开发效率是很慢的。

Gin框架提供给开发者表单实体绑定的功能，可以将表单数据与结构体绑定：
type UserRegister struct {
	Username string `form:"username" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required"`
}
*/

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

type Register struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Phone    string `form:"phone"`
}

func main() {
	engine := gin.Default()

	// http://localhost:8080/hello?name=dave&classes=软件工程
	engine.GET("hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Student
		// 数据表单绑定结构体，注意传递的是student的引用
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		// 能获取到值的情况
		fmt.Println(student.Name)
		fmt.Println(student.Classes)

		context.Writer.Write([]byte("hello, " + student.Name))
	})

	// 处理Post请求
	engine.POST("/register", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var register Register

		if err := context.ShouldBind(&register); err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(register.Username)
		fmt.Println(register.Phone)

		context.Writer.Write([]byte("username: " + register.Username))
	})

	// 使用json格式的数据
	engine.POST("/add", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Register
		// 完成json格式数据的绑定
		if err := context.BindJSON(&student); err != nil {
			log.Fatal(err.Error())
			return
		}

		/*
			    通过Postman提交接口请求：
				{
				    "username": "kevin12121",
				    "phone": "1232323232",
				    "password": "222222222"
				}
		*/

		fmt.Println(student.Username)
		fmt.Println(student.Phone)

		context.Writer.Write([]byte("username: " + student.Username))

	})

	if err := engine.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

```


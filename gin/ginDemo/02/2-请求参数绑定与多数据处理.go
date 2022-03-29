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

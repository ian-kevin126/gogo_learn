package main

import (
	"github.com/gin-gonic/gin"
	"gogo_learn/gin/ginApiProject/controller"
	"gogo_learn/gin/ginApiProject/utils"
)

func main() {

	config := utils.GetConfig()

	app := gin.Default()

	registerRouter(app)

	app.Run(config.AppHost + ":" + config.AppPort)
}

// 路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}

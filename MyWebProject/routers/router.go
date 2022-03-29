package routers

import (
	"MyWebProject/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/test/arg", &controllers.TestArgController{}, "get:Get;post:Post")
	beego.Router("/test/orm", &controllers.TestModelController{}, "get:Get;post:Post")
}

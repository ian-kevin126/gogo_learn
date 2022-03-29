package controllers

import beego "github.com/beego/beego/v2/server/web"

type User1 struct {
	Id       int    `form:"-"`
	Username string `form:"username"`
	Age      string `form:"age"`
	Email    string `form:"email"`
}

type TestArgController struct {
	beego.Controller
}

func (c *TestArgController) Get() {
	c.TplName = "index.tpl"
}

func (c *TestArgController) Post() {
	//id := c.GetString("id")
	//name := c.GetString("name")
	//c.Ctx.WriteString(id + name)

	u := User1{}

	if err := c.ParseForm(&u); err != nil {

	}

	c.Ctx.WriteString("username: " + u.Username + "Age: " + u.Age + "Email: " + u.Email)
}

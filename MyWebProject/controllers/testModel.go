package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type TestModelController struct {
	beego.Controller
}

type User struct {
	Id       int64
	Username string
	Password string
}

func (c *TestModelController) Get() {
	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/db_name?charset=utf8&loc=Local")

	// register model
	orm.RegisterModel(new(User))

	// create table
	o := orm.NewOrm()
	// 插入数据
	user := User{Username: "zhangsan", Password: "123456"}
	id, _ := o.Insert(&user)

	// 操作数据库
	c.Ctx.WriteString(fmt.Sprintf("操作数据库id=%v", id))
}

package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["IsLogin"] = CheckAccount(c.Ctx)
	c.Data["IsHome"] = true
	fmt.Println("aaaa=>", c.Data["IsLogin"])
	c.TplName = "home.html"
}

package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsHome"] = true
	c.Data["aa"] = "aaaaa"
	c.TplName = "home.html"
}

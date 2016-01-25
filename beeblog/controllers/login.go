package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"math"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
	if c.Input().Get("exit") == "true" {
		c.Ctx.SetCookie("uname", "", -1)
		c.Ctx.SetCookie("pwd", "", -1)
	}
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		if autoLogin {
			maxAge := math.MaxInt32
			c.Ctx.SetCookie("uname", uname, maxAge)
			c.Ctx.SetCookie("pwd", pwd, maxAge)
		} else {
			c.Ctx.SetCookie("uname", uname)
			c.Ctx.SetCookie("pwd", pwd)
		}
	} else {
		c.Ctx.SetCookie("uname", "", -1)
		c.Ctx.SetCookie("pwd", "", -1)
	}

	c.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	uname := ctx.GetCookie("uname")
	pwd := ctx.GetCookie("pwd")

	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		return true
	}
	return false
}

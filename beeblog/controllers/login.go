package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("uname", " ", -1, "/")
		this.Ctx.SetCookie("pwd", " ", -1, "/")
		this.Redirect("/", 302)
	}
	this.TplName = "login.html"
	//this.Data["IsLogin"] = true
}

func (c *LoginController) Post() {
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	if uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass") {
		if autoLogin {
			maxAge := 1<<31 - 1
			c.Ctx.SetCookie("uname", uname, maxAge, "/")
			c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
		} else {
			c.Ctx.SetCookie("uname", uname)
			c.Ctx.SetCookie("pwd", pwd)
		}

	} else {
		c.Ctx.SetCookie("uname", "", -1, "/")
		c.Ctx.SetCookie("pwd", "", -1, "/")
	}

	c.Redirect("/", 302)
	return
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value
	return uname == beego.AppConfig.String("adminName") &&
		pwd == beego.AppConfig.String("adminPass")
}

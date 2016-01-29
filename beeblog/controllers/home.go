package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsHome"] = true
	var err error
	c.Data["Topics"], err = models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "home.html"
}

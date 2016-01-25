package controllers

import (
	"beeblog/models"
	//"fmt"
	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "topic.html"
	c.Data["IsTopic"] = true
}

func (c *TopicController) Add() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//确保已注册
	if c.Data["IsLogin"] == false {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTopic"] = true
	c.TplName = "topic_add.html"
}

func (c *TopicController) Post() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//确保已注册
	if c.Data["IsLogin"] == false {
		c.Redirect("/login", 302)
		return
	}

	title := c.Input().Get("title")
	content := c.Input().Get("content")

	if title != "" && content != "" {
		err := models.AddTopic(title, content)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/topic", 302)
		return
	} else {
		c.Redirect("/topic", 401)
		return
	}
	return
}

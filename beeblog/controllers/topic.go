package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-study/beeblog/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "topic.html"
	c.Data["IsTopic"] = true
	var err error
	c.Data["Topics"], err = models.GetAllTopics("", true)
	fmt.Println(c.Data["Topics"])
	if err != nil {
		beego.Error(err)
	}
}

func (c *TopicController) Add() {
	//确保已注册
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	if c.Data["IsLogin"] == false {
		c.Redirect("/login", 302)
		return
	}
	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.Data["IsTopic"] = true
	c.TplName = "topic_add.html"
}

func (c *TopicController) Post() {
	//确保已注册
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	if c.Data["IsLogin"] == false {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTopic"] = true

	title := c.Input().Get("title")
	content := c.Input().Get("content")
	id := c.Input().Get("tid")
	category := c.Input().Get("category")

	if title != "" && content != "" && category != "" {
		if id == "" {
			err := models.AddTopic(title, content, category)
			if err != nil {
				beego.Error(err)
			}
		} else {
			err := models.ModifyTopic(id, title, content, category)
			if err != nil {
				beego.Error(err)
			}
		}

		c.Redirect("/topic/view/"+id, 302)
		return
	} else {
		c.Redirect("/topic", 401)
		return
	}
	return
}

func (c *TopicController) Delete() {
	//确保已注册
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	if c.Data["IsLogin"] == false {
		c.Redirect("/login", 302)
		return
	}
	tid := c.Input().Get("tid")
	models.DeleteTopic(tid)
	c.Redirect("/topic", 302)
	return
}

func (c *TopicController) View() {
	//确保已注册
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.Data["IsTopic"] = true
	tid := c.Ctx.Input.Params()["0"]
	if tid == "" {
		c.Redirect("/topic", 302)
		return
	}
	var err error
	c.Data["Topic"], err = models.ShowTopic(tid)
	c.Data["Tid"] = tid
	if err != nil {
		beego.Error(err)
		c.Redirect("/topic", 302)
		return
	}
	c.TplName = "topic_view.html"
}

func (c *TopicController) Modify() {
	//确保已注册
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	if c.Data["IsLogin"] == false {
		c.Redirect("/login", 302)
		return
	}
	c.Data["IsTopic"] = true
	tid := c.Input().Get("tid")
	if tid == "" {
		c.Redirect("/", 401)
		return
	}
	var err error
	c.Data["Topic"], err = models.ShowTopic(tid)
	if err != nil {
		beego.Error(err)
		c.Redirect("/", 401)
		return
	}
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	c.TplName = "topic_modify.html"
}

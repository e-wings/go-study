package controllers

import (
	"github.com/astaxie/beego"
	"github.com/e-wings/go-tools/enet/log"
	"go-study/ablog/models"
)

type TopicController struct {
	beego.Controller
}

func (c *TopicController) Get() {
	c.TplName = "topic.html"
}

func (c *TopicController) Add() {
	cats, err := models.GetCategories(1)

	if err != nil {
		log.Critical(err.Error())
		return
	}
	c.Data["Categories"] = cats
	c.TplName = "topic_add.html"
}

func (c *TopicController) Post() {
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	cat_id := c.Input().Get("category")

	if title == "" || content == "" || cat_id == "" {
		log.Warn("信息不全，无法新建文章。")
	} else {
		err := models.AddTopic(title, content, cat_id)
		if err != nil {
			log.Warn(err.Error())
		}
	}
	c.Redirect("/topic", 302)
	return
}

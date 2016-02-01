package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) Add() {
	tid := c.Input().Get("tid")
	if tid == "" {
		c.Redirect("/topic/view/tid", 302)
		return
	}
	nickname := c.Input().Get("nickname")
	content := c.Input().Get("content")
	models.AddReply(tid, nickname, content)
	c.Redirect("/topic/view/"+tid, 302)
	return
}

func (c *ReplyController) Delete() {
	tid := c.Input().Get("tid")
	if tid == "" {
		c.Redirect("/topic", 302)
		return
	}
	rid := c.Input().Get("rid")
	if rid == "" {
		c.Redirect("/topic/view/"+tid, 302)
		return
	}
	models.DeleteReply(tid, rid)
	c.Redirect("/topic/view/"+tid, 302)
	return
}

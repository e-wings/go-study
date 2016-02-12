package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
	//"go-study/ablog/log"
	"go-tools/enet/log"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	c.Data["IsCategory"] = true
	op := c.Input().Get("op")
	if op == "add" {
		cName := c.Input().Get("name")
		if len(cName) < 50 {
			log.Warn("类型名不能为空")
		}
	}

	c.TplName = "category.html"
}

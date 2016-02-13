package controllers

import (
	"github.com/astaxie/beego"
	//"fmt"
	"github.com/astaxie/beego/orm"
	"go-study/ablog/models"
	"go-tools/enet/log"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	o := orm.NewOrm()

	c.Data["IsCategory"] = true
	op := c.Input().Get("op")
	if op == "add" {
		cName := c.Input().Get("name")
		if len(cName) == 0 {
			log.Warn("类型名不能为空")
		} else {
			cat := &models.Category{
				Title: cName,
			}

			_, err := o.Insert(cat)
			if err != nil {
				log.Warn(err.Error())
			}
		}
		c.Redirect("/category", 302)
		return
	} else if op == "switch" {
		id := c.Input().Get("id")
		status := c.Input().Get("status")
		err := models.SwitchCategoryStatus(id, status)
		if err != nil {
			log.Critical(err.Error())
		}
	}
	cats, err := models.GetCategories(-1)
	if err != nil {
		log.Warn(err.Error())
	}
	c.Data["Categories"] = cats
	c.TplName = "category.html"
}

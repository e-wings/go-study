package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go-study/beeblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	c.Data["IsCategory"] = true
	c.TplName = "category.html"

	op := c.Input().Get("op")
	switch op {
	case "add":
		name := c.Input().Get("name")
		if name == "" {
			return
		}
		err := models.AddCategory(name)
		if err != nil {
			fmt.Println(err)
		}
		c.Redirect("/category", 302)
		return
	case "del":
		id := c.Input().Get("id")
		if id == "" {
			return
		}
		models.DeleteCategory(id)
		c.Redirect("/category", 302)
		return
	}

	var err error
	c.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

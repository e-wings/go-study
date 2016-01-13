package controllers

import (
	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Get() {

}

func (c *FormController) Post() {
	c.SaveToFile("uploadname", "/www/uploaded_file.txt")
}

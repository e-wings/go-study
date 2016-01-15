package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

func (c *FormController) Get() {

}

func (c *FormController) Post() {
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		c.SaveToFile("uploadname", "/www/"+h.Filename)
	}
}

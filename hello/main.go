package main

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	this.Ctx.WriteString("Hello world!")
}

func main() {
	beego.Router("/", &IndexController{})
	beego.Run(":8081")
}

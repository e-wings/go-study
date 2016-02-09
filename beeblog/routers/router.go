package routers

import (
	"github.com/astaxie/beego"
	"go-study/beeblog/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.AutoRouter(&controllers.ReplyController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/login", &controllers.LoginController{})
}

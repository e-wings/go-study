package routers

import (
	"github.com/astaxie/beego"
	"go-study/ablog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.AutoRouter(&controllers.CategoryController{})
	beego.Router("/login", &controllers.LoginController{})
}

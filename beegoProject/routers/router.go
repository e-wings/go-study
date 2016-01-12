package routers

import (
	"beegoProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//beego.Router("/users", &controllers.UsersController{}, "*:Index")
	beego.Router("/users/:id", &controllers.UsersController{}, "get:Get")
	beego.Router("/users/list", &controllers.UsersController{}, "get:GetAll")
	beego.Router("/users/update/:id", &controllers.UsersController{}, "put:Put")
	beego.Router("/users/add", &controllers.UsersController{}, "post:Post")
	beego.Router("/users/delete/:id", &controllers.UsersController{}, "delete:Delete")
}

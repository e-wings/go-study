package main

import (
	_ "beegoProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.EnableAdmin = true
	beego.EnableXSRF = true
	beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.XSRFExpire = 3600 //过期时间，默认60秒
	beego.Run()
}

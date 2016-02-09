package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-study/beeblog/models"
	_ "go-study/beeblog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	//自动建表
	if beego.AppConfig.String("runmode") == "pro" {
		orm.Debug = false
		orm.RunSyncdb("default", false, false)
	} else {
		orm.Debug = true
		orm.RunSyncdb("default", false, false)
	}

	beego.Run()
}

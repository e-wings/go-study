package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-study/ablog/models"
	_ "go-study/ablog/routers"
)

func init() {
	models.RegisterDB()
}

func main() {
	if beego.AppConfig.String("runmode") == "pro" {
		orm.Debug = false
		orm.RunSyncdb("default", false, false)
	} else {
		orm.Debug = true
		orm.RunSyncdb("default", false, true) //数据库链接、是否强制覆盖、是否更新
	}
	beego.Run()
}

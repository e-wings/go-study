// package main

// import (
// 	_ "beegoProject/routers"
// 	"fmt"
// 	"github.com/astaxie/beego"
// 	"github.com/astaxie/beego/orm"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	beego.EnableAdmin = true
// 	beego.EnableXSRF = true
// 	beego.XSRFKEY = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
// 	beego.XSRFExpire = 3600 //过期时间，默认60秒
// 	beego.Run()
// }

// func init() {
// 	fmt.Println("=====================")
// 	orm.RegisterDriver("mysql", orm.DR_MySQL)
// 	err := orm.RegisterDataBase("default", "mysql", "root@tcp(localhost:3306)/beego_test", 30)
// 	fmt.Println(err)
// 	//beego.error()
// }

package main

import (
	_ "beegoProject/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	Id       int64  `beedb:"PK"`
	Username string `orm:"size(128)"`
	Age      int
}

func init() {
	orm.RegisterModel(new(Users))
	dbLink := beego.AppConfig.String("dev::mysqluser") + ":" + beego.AppConfig.String("dev::mysqlpass") + "@tcp(" + beego.AppConfig.String("dev::mysqlurls") + ":3306)/" + beego.AppConfig.String("dev::mysqldb")
	orm.RegisterDataBase("default", "mysql", dbLink, 30)
}

func main() {
	o := orm.NewOrm()
	user := Users{Username: "hello", Age: 35}
	id, err := o.Insert(&user)
	fmt.Printf("ID:%d, Err: %v\n ", id, err)

	user.Username = "ttt"
	user.Age = 36
	num, err := o.Update(&user)
	fmt.Printf("Num:%d ,ERR: %v\n", num, err)
	//beego.StaticDir["/upload"] = "upload"
	beego.SetStaticPath("/upload", "upload")
	beego.Run()

}

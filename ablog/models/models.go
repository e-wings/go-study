package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
)

const (
	_DB_Name        = "data/ablog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id        int64
	Title     string
	Created   string `orm:"index"`
	Views     int64  `orm:"index"`
	TopicTime string `orm:"index"`
	//Topic     []*Topic `orm:"reverse(many)"`
}

func RegisterDB() {
	if !com.IsExist(_DB_Name) {
		os.MkdirAll(path.Dir(_DB_Name), os.ModePerm)
		os.Create(_DB_Name)
	}

	orm.RegisterModel(new(Category))
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_Name, 10)
}

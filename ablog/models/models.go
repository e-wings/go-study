package models

import (
	"errors"
	//"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
)

const (
	_DB_Name        = "data/ablog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id         int64
	Title      string
	Created    string `orm:"index"`
	Views      int64  `orm:"index"`
	TopicTime  string `orm:"index"`
	TopicCount int64  `orm:"-"`
	Status     int8   `orm:"default(1)"`
	//Topic     []*Topic `orm:"reverse(many)"`
}

type Topic struct {
	Id         int64
	Title      string
	Content    string `orm:"size(20000)"`
	Attachment string
	Category   *Category `orm:"rel(fk)"`
	Created    string    `orm:"index"`
	Updated    string    `orm:"index"`
	Views      int64
	ReplyTime  string `orm:"null"`
}

func RegisterDB() {
	if !com.IsExist(_DB_Name) {
		os.MkdirAll(path.Dir(_DB_Name), os.ModePerm)
		os.Create(_DB_Name)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_Name, 10)
}

func GetCategories(status int8) ([]*Category, error) {
	o := orm.NewOrm()
	cats := make([]*Category, 0)
	qs := o.QueryTable("category")
	var err error
	if status > -1 {
		_, err = qs.Filter("status", status).All(&cats)
	} else {
		_, err = qs.All(&cats)
	}
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func SwitchCategoryStatus(categoryId string, categoryStatus string) error {
	id, err := strconv.ParseInt(categoryId, 10, 64)
	if err != nil {
		return err
	}

	var status int8
	var tmp int64
	tmp, err = strconv.ParseInt(categoryStatus, 10, 8)
	if err != nil {
		return err
	}
	status = int8(tmp)

	if status == 0 || status == 1 {
		status = 1 - status
	}

	cat := Category{
		Id:     id,
		Status: status,
	}
	o := orm.NewOrm()
	resNum, updateErr := o.Update(&cat, "status")
	if updateErr != nil {
		return updateErr
	}
	if resNum == 0 {
		return errors.New("未找到该Category")
	}
	return nil
}

func AddTopic(title, content, cat_id string) error {
	catId, err := strconv.ParseInt(cat_id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	//确保Category ID 没超出范围
	cat := Category{}
	qs := o.QueryTable("category")
	err = qs.Filter("id", catId).One(&cat)
	if err != nil {
		return err
	}

	topic := Topic{
		Title:    title,
		Content:  content,
		Category: &cat,
	}
	_, err = o.Insert(&topic)
	return err
}

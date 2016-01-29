package models

import (
	//"fmt"
	"errors"
	//"fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
	Topic           []*Topic `orm:"reverse(many)"`
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Category        *Category `orm:"rel(fk)"`
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64
	Author          string
	ReplyTime       time.Time `orm:"null"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}

	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	cate := &Category{}
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return errors.New("Category名重复！")
	}
	cate = &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	_, err = o.Insert(cate)
	return err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	cate := &Category{Id: cid}
	o := orm.NewOrm()
	_, err = o.Delete(cate)
	return err
}

func AddTopic(title, content, category string) error {
	// cid, err := strconv.ParseInt(category, 10, 64)
	// if err != nil {
	// 	return err
	// }
	topic := &Topic{Title: title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
		//Category: cid,
	}
	o := orm.NewOrm()
	_, err := o.Insert(topic)
	return err
}

func GetAllTopics(isDesc bool) ([]*Topic, error) {
	topics := make([]*Topic, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").RelatedSel().All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func DeleteTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 32)
	if err != nil {
		return err
	}
	topic := Topic{Id: id}
	o := orm.NewOrm()
	_, err = o.Delete(&topic)
	return err
}

func ShowTopic(id string) (*Topic, error) {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	topic := &Topic{}
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).RelatedSel().One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

func ModifyTopic(id, title, content, category string) error {
	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	// cid, err := strconv.ParseInt(category, 10, 64)
	// if err != nil {
	// 	return err
	// }
	topic := &Topic{}
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	if err != nil {
		return err
	}
	topic.Title = title
	topic.Content = content
	topic.Updated = time.Now()
	//topic.Category = cid
	_, err = o.Update(topic)
	return err
}

package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20160111_222540 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20160111_222540{}
	m.Created = "20160111_222540"
	migration.Register("Users_20160111_222540", m)
}

// Run the migrations
func (m *Users_20160111_222540) Up() {
	// use m.Sql("CREATE TABLE ...") to make schema update
	m.Sql("CREATE TABLE users(`id` int(11) DEFAULT NULL,`username` varchar(128) NOT NULL,`age` int(11) DEFAULT NULL)")
}

// Reverse the migrations
func (m *Users_20160111_222540) Down() {
	// use m.Sql("DROP TABLE ...") to reverse schema update
	m.Sql("DROP TABLE `users`")
}

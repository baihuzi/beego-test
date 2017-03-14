package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

func GetAllUser() (user []*User, err error) {
	o := orm.NewOrm()

	user = make([]*User, 0)

	qt := o.QueryTable("user")
	_, err = qt.All(&user)

	return user, err
}


func AddUser(){user []*User,err error} {
	o := orm.NewOrm()
}
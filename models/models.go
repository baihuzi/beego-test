package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Profile), new(Post), new(Tag))
	//注册驱动
	orm.RegisterDriver(beego.AppConfig.String("dbDriver"), orm.DRMySQL)
	//mysql连接
	db_config := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlurls") + ":3306)/" + beego.AppConfig.String("mysqldb") + "?charset=utf8"
	//注册数据库
	orm.RegisterDataBase("default", beego.AppConfig.String("dbDriver"), db_config)
}

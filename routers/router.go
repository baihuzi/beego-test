package routers

import (
	"github.com/astaxie/beego"
	"myweb/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})
	beego.AutoRouter(&controllers.ListController{})
}

package controllers

import (
	"github.com/astaxie/beego"
	"myweb/models"
)

type ListController struct {
	beego.Controller
}

func (l *ListController) Get() {
	t := l.Input().Get("t")

	l.Ctx.WriteString(t)
}

func (l *ListController) Test() {
	l.TplName = "list.html"

	users, err := models.GetAllUser()
	if err != nil {
		beego.Error(err)
	}
	l.Data["Users"] = users
}

package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/logs"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	var user string = ""
	var ok bool
	user, ok = this.GetSession(SESSIONUSER).(string)
	if !ok {
		user = ""
	}
	this.TplName = "admin/login.html"
	this.Data["Failed"] = false
	this.Data["User"] = user
}

func (this *LoginController) Post() {
	user := this.GetString("useremail")
	psd := this.GetString("password")
	logout := this.GetString("logout")
	if logout == "true" {
		this.DelSession(SESSIONUSER)
		this.Data["Failed"] = false
		this.Data["User"] = ""
	} else {
		if validate(user, psd) {
			this.SetSession(SESSIONUSER, user)
			this.Data["Failed"] = false
			this.Data["User"] = user
		} else {
			this.Data["Failed"] = true
			this.Data["User"] = ""
		}
	}
	this.TplName = "admin/login.html"
}

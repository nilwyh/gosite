package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/logs"
)

const (
	HOUR = 3600
)

func init() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "SESSIONID"
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 24 * HOUR
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 1 * HOUR

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/blog", &controllers.BlogIndexController{})
	beego.Router("/blog/:id([0-9]+)", &controllers.BlogController{})
	beego.InsertFilter("*", beego.BeforeRouter, RedirectHttps)
}

// 重定向到https
var RedirectHttps = func(ctx *context.Context) {
	if beego.BConfig.Listen.EnableHTTPS && ctx.Input.Scheme() == "http" {
		ctx.Redirect(301, fmt.Sprintf("%s%s", controllers.Domain, ctx.Input.URL()))
	}
}
package controllers

import "github.com/astaxie/beego"

var Domain string

var SESSIONUSER = beego.AppConfig.String("sessionname")

func init() {
	if beego.BConfig.Listen.EnableHTTPS {
		Domain = "https://localhost:8181"
	} else {
		Domain = "http://localhost:8080"
	}
}

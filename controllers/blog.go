package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
	"io/ioutil"
	"html/template"
)


type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	content, err := ioutil.ReadFile("/home/yihan/golang/src/gosite/static/text/example.md")
	md := blackfriday.MarkdownCommon(content)
	html := bluemonday.UGCPolicy().SanitizeBytes(md)
	//id := c.Ctx.Input.Param(":id")
	id := c.GetString(":id")
	logs.Warn("!!!!!!!!!!!!!!! id : %s", id)

	c.Data["Website"] = "Yihan's Blog"
	c.Data["Email"] = "wangff9@gmail.com"
	c.Data["Title"] = id
	if err != nil {
		c.Data["Content"] = template.HTML("")
	} else {
		c.Data["Content"] = template.HTML(html[:])
	}
	c.LayoutSections = make(map[string]string)
	c.TplName = "blog/blog_content.html"
}

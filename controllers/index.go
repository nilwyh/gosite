package controllers

import (
	"github.com/astaxie/beego"
)

type BlogIndexController struct {
	beego.Controller
}

func (c *BlogIndexController) Get() {
	c.Data["Website"] = "Yihan's Blog"
	c.Data["Email"] = "wangff9@gmail.com"
	c.LayoutSections = make(map[string]string)
	c.TplName = "blog/blog_index.html"
}

func (c *BlogIndexController) Post() {
	key := c.GetString("search_key")
	c.Ctx.Redirect(302, "/blog/"+key)
}

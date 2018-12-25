package home

import "github.com/astaxie/beego"

type DetailController struct {
	beego.Controller
}

func (c *DetailController) Get() {

	c.TplName = "detail.html"
}

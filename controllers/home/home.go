package home

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Prepare() {

	//log.Info("  ------- init ----------- ")

	this.Data["xsrf_token"] = this.XSRFToken()

	this.Data["uid"] = this.GetSession("uid")
	this.Data["username"] = this.GetSession("username")

	this.Layout = "home/layout.html"
}


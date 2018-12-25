package center

import (
	"github.com/astaxie/beego"
	"github.com/labstack/gommon/log"
)

type CenterController struct {
	beego.Controller
}

func (this *CenterController) Prepare() {

	this.Data["xsrf_token"] = this.XSRFToken()

	//TODO:判断用户是否登录
	uid := this.GetSession("uid")
	if (uid == nil) {
		log.Info("您还没有登录。。。")
		this.Redirect("/", 302)
		return
	}

	//TODO：初始化模板布局
	this.Layout = "center/layout.html";
}

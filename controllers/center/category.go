package center

import "github.com/labstack/gommon/log"

type CategoryController struct {
	CenterController
}

func (this *CategoryController) Prepare() {
	this.Layout = "center/layout.html"
}

func (this *CategoryController) Get() {

	uid := this.GetSession("uid")
	username := this.GetSession("username")

	log.Info(uid)
	log.Info(username)

	this.TplName = "center/category.html";
}

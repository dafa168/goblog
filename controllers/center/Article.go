package center

import "github.com/labstack/gommon/log"

type ArticleController struct {
	CenterController
}

func (this *ArticleController) Prepare() {
	this.Layout = "center/layout.html"
}

func (this *ArticleController) Get() {

	uid := this.GetSession("uid")
	username := this.GetSession("username")

	log.Info(uid)
	log.Info(username)

	this.TplName = "center/article.html";
}

package center

import (
	"github.com/labstack/gommon/log"
)

type UserController struct {
	CenterController
}

func (this *UserController) Get() {

	uid := this.GetSession("uid")
	username := this.GetSession("username")

	log.Info(uid)
	log.Info(username)


	this.TplName = "center/user.html";
}

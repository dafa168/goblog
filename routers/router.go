package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers/home"
	"goblog/controllers/center"
)

func init() {

	beego.Router("/", &home.MainController{})

	beego.Router("/category", &home.CategoryController{})
	beego.Router("/article", &home.DetailController{})

	beego.Router("/login", &home.LoginController{})
	beego.Router("/reg", &home.RegisterController{})

	beego.Router("/center", &center.CenterController{})
	beego.Router("/center/user", &center.UserController{})
	beego.Router("/center/category", &center.CategoryController{})
	beego.Router("/center/article", &center.CategoryController{})
	beego.Router("/center/setting", &center.CategoryController{})

	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{})

	//beego.Router("/user/register", &controllers.UserController{}, "register")
	//beego.Router("/user/login", &controllers.UserController{})
	//beego.Router("/user/logout", &controllers.UserController{})
	//beego.Router("/user/find", &controllers.UserController{})
}

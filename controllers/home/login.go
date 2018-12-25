package home

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"goblog/models"
	"github.com/labstack/gommon/log"
	"goblog/util"
	"github.com/astaxie/beego/validation"
	"html/template"
)

type LoginController struct {
	beego.Controller
}

func init() {
	//log.Info("---------------login init----------------")
	//if (util.CheckLogin())
}

func (this *LoginController) Prepare() {
	this.XSRFExpire = 7200
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrf_token"] = this.XSRFToken()
}

func (this *LoginController) Get() {
	this.TplName = "home/login.html";
}

func (this *LoginController) Post() {

	user := new(models.User)
	user.Username = this.Input().Get("username")
	user.Password = this.Input().Get("password")

	//TODO:验证数据
	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Password, "password")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
		}

		this.Redirect("/login", 302)
		return
	}

	var users models.User
	q := orm.NewOrm()
	q.Using("default")

	err := q.QueryTable("user").Filter("username", user.Username).One(&users)

	if (err == orm.ErrNoRows) {
		log.Info("用户不存在")
		this.Redirect("/login", 302);
		return
	}

	//开始验证码密码
	if (!util.CheckPass(user.Password, users.Password)) {
		log.Info("密码不正确")
		this.Redirect("/login", 302)
		return
	}

	//更新登录时间
	//user.LoginAt = time.Now()
	//q.Update(user, "login_at");

	//写session
	this.SetSession("uid", users.Id);
	this.SetSession("username", users.Username)
	this.SetSession("login_time", users.LoginAt)
	this.SetSession("user_auth", users.Username+user.Password)

	//跳转
	this.Redirect("/user", 302)
	return
}

package home

import (
	"github.com/astaxie/beego/orm"
	"goblog/models"
	"github.com/labstack/gommon/log"
	"github.com/astaxie/beego/validation"
	"html/template"
	"goblog/util"
	"time"
)

type RegisterController struct {
	HomeController
}

func (this *RegisterController) Prepare() {
	this.XSRFExpire = 7200
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrf_token"] = this.XSRFToken()
}

func (this *RegisterController) Get() {
	this.TplName = "home/reg.html";
}

func (this *RegisterController) Post() {

	//TODO:数据验证
	user := new(models.User)
	user.Username = this.Input().Get("username")
	user.Password = this.Input().Get("password")
	user.Status = this.Input().Get("status")

	valid := validation.Validation{}
	valid.Required(user.Username, "username")
	valid.Required(user.Password, "password")
	valid.Required(user.Status, "status")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			log.Info(err.Key, err.Message)
		}

		this.Redirect("/reg", 302)
		return
	}

	q := orm.NewOrm()
	q.Using("default")
	flag := q.QueryTable("user").Filter("username", user.Username).Exist()

	if (flag == true) {
		log.Info("用户已经存在了。")
		this.Redirect("/login", 302)
		return
	}

	//i, _ := strconv.Atoi(status)
	//user.Status = int8(i)

	user.Password = util.EncodePass(user.Password)
	user.CreateAt = time.Now()
	user.LoginAt = time.Now()

	var err error
	_, err = q.Insert(user)

	if (err != nil) {
		log.Info("用户注册失败 " + err.Error())
		this.Redirect("/reg", 302)
		return
	}

	this.Redirect("/login", 302);
	return
}

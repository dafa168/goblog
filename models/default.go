package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Comments struct {
	Id       int64
	Aid      int64
	Info     string `orm:"size(2000)"`
	CreateAt int64
	UpdateAt int64
}

type Article struct {
	Id       int64
	Cid      int64
	Title    string
	Info     string
	Content  string `orm:"size(20000)"`
	Views    int64
	CreateAt time.Time
	UpdateAt time.Time
}

type Category struct {
	Id       int64
	Name     string
	CreateAt time.Time
}

type User struct {
	Id       int64
	Username string `orm:"size(50)"`
	Password string `orm:"size(128)"`
	Status   string
	CreateAt time.Time
	LoginAt  time.Time
}

func init() {

	orm.DefaultTimeLoc = time.Local

	orm.Debug = true

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego?charset=utf8&loc=Asia%2FShanghai", 30)

	orm.RegisterModel(new(User), new(Article), new(Category))

	orm.RunSyncdb("default", false, true)
}

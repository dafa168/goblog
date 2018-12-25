package home

type MainController struct {
	HomeController
}

func (this *MainController) Get() {


	this.TplName = "index.html"
}

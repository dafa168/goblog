package home

type CategoryController struct {
	HomeController
}

func (this *CategoryController) Get() {

	this.TplName = "category.html";
}

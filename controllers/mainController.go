package controllers

type MainController struct {
	BaseController
}

// @router / [get]
func (this *MainController) Get() {
	this.Redirect("/signin", 302)
}

// @router /intherow [get]
func (this *MainController) InTheRow() {
	this.TplName = "intherow.html"
}

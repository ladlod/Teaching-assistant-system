package controllers

type UsersController struct {
	BaseController
}

// @router / [get]
func (c *UsersController) Get() {
	c.TplName = "login.html"
}

package controllers

import (
	"Teching-assistant-system/models"

	"github.com/astaxie/beego"
)

type UsersController struct {
	BaseController
}

// @router /signin [get]
func (this *UsersController) GetSignin() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	} else if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}
	this.TplName = "signin.html"
}

// @router /signin [post]
func (this *UsersController) PostSignin() {
	user := models.User{}
	flash := beego.NewFlash()
	user.Name = this.GetString("username")
	user.Password = this.GetString("password")
	if this.GetString("usertype") == "Teacher" {
		user.Identity = 0
	} else {
		user.Identity = 1
	}

	if user.Signin() {
		this.Redirect("/signin", 302)
		return
	} else {
		flash.Error("用户不存在或密码错误")
		flash.Store(&this.Controller)
		this.Redirect("/signin", 302)
		return
	}
}

// @router /signup [get]
func (this *UsersController) GetSignup() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	} else if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}
	this.TplName = "signup.tpl"
}

// @router /signup [post]
func (this *UsersController) PostSignup() {
	user := models.User{}
	flash := beego.NewFlash()
	user.Name = this.GetString("username")
	user.Password = this.GetString("password")
	if this.GetString("usertype") == "Teacher" {
		user.Identity = 0
	} else {
		user.Identity = 1
	}

	if user.Signup() {
		flash.Notice("注册成功")
		flash.Store(&this.Controller)
		this.Redirect("/signin", 302)
	} else {
		flash.Error("用户名已存在")
		flash.Store(&this.Controller)
		this.Redirect("/signup", 302)
	}
}

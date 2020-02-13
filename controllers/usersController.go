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
	flash := beego.NewFlash()
	if this.GetString("usertype") == "on" {
		var teacher models.Teacher
		teacher.Account = this.GetString("account")
		teacher.Password = this.GetString("password")

		if teacher.Signin() {
			this.Redirect("/intherow", 302)
			return
		} else {
			flash.Error("用户不存在或密码错误")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
			return
		}
	} else {
		var student models.Teacher
		student.Account = this.GetString("account")
		student.Password = this.GetString("password")

		if student.Signin() {
			this.Redirect("/intherow", 302)
			return
		} else {
			flash.Error("用户不存在或密码错误")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
			return
		}
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
	this.TplName = "signup.html"
}

// @router /signup [post]
func (this *UsersController) PostSignup() {
	flash := beego.NewFlash()
	if this.GetString("usertype") == "on" {
		var teacher models.Teacher
		teacher.Account = this.GetString("account")
		teacher.Name = this.GetString("username")
		teacher.Password = this.GetString("password")

		if teacher.Signup() {
			flash.Notice("注册成功")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
		} else {
			flash.Error("账户已存在")
			flash.Store(&this.Controller)
			this.Redirect("/signup", 302)
		}
	} else {
		var student models.Student
		student.Account = this.GetString("account")
		student.Name = this.GetString("username")
		student.Password = this.GetString("password")

		if student.Signup() {
			flash.Notice("注册成功")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
		} else {
			flash.Error("账户已存在")
			flash.Store(&this.Controller)
			this.Redirect("/signup", 302)
		}
	}
}

package routers

import (
	"Teching-assistant-system/controllers"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

// 过滤器
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("identity").(string)
	if !ok && ctx.Request.RequestURI != "/signin" && ctx.Request.RequestURI != "/signup" {
		ctx.Redirect(302, "/signin")
	}
}
var FilterStudent = func(ctx *context.Context) {
	identity, _ := ctx.Input.Session("identity").(string)
	if identity != "student" {
		ctx.Redirect(302, "signin")
	}
}
var FilterTeacher = func(ctx *context.Context) {
	identity, _ := ctx.Input.Session("identity").(string)
	if identity != "teacher" {
		ctx.Redirect(302, "signin")
	}
}

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.CourseController{})
	beego.Include(&controllers.UsersController{})
}

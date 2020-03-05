package controllers

import (
	"Teaching-assistant-system/models"

	"github.com/astaxie/beego"
)

type ClockinController struct {
	BaseController
}

// @router /teacher/course/setclockin/:password [get]
func (this *ClockinController) SetClock() {
	flash := beego.NewFlash()

	password := this.Ctx.Input.Param(":password")
	course := this.GetSession("course").(models.Course)
	if course.AddClockin(password) {
		flash.Error("开启签到成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
		return
	} else {
		flash.Error("开启签到失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
		return
	}
}

// @router /student/course/clockin/:password [get]
func (this *ClockinController) Clockin() {
	flash := beego.NewFlash()
	password := this.Ctx.Input.Param(":password")
	student := this.GetSession("account").(models.Student)
	course := this.GetSession("course").(models.Course)

	if !course.JudgeClockin() {
		flash.Error("暂无开启的签到")
		flash.Store(&this.Controller)
		this.Redirect("/student/course", 302)
		return
	}

	if student.Clockin(password, &course) {
		flash.Error("签到成功")
		flash.Store(&this.Controller)
		this.Redirect("/student/course", 302)
		return
	} else {
		flash.Error("密码错误或签到时间已截止")
		flash.Store(&this.Controller)
		this.Redirect("/student/course", 302)
		return
	}
}

package controllers

import (
	"Teching-assistant-system/models"
	"strconv"

	"github.com/astaxie/beego"
)

type CourseController struct {
	BaseController
}

// @router /student/:cid [get]
func (this *CourseController) StudentSelectCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	course := models.SearchCourse(cid)
	this.SetSession("course", course)
	this.Redirect("/student/course", 302)
}

// @router /student/course [get]
func (this *CourseController) GetStudentCourse() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}
	course := this.GetSession("course").(*models.Course)
	students := course.QueryStudents()
	this.Data["course"] = course
	this.Data["students"] = students
	this.TplName = "studentcourse.html"
}

// @router /teacher/:cid [get]
func (this *CourseController) TeacherSelectCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	course := models.SearchCourse(cid)
	this.SetSession("course", course)
	this.Redirect("/teacher/course", 302)
}

// @router /teacher/course [get]
func (this *CourseController) GetTeacherCourse() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}
	course := this.GetSession("course").(*models.Course)
	students := course.QueryStudents()
	this.Data["course"] = course
	this.Data["students"] = students
	this.TplName = "studentcourse.html"
}

// @router /teacher/dropcourse/:cid [get]
func (this *CourseController) DeleteCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	teacher := this.GetSession("account").(models.Teacher)
	flash := beego.NewFlash()
	if teacher.DeleteCourse(cid) {
		flash.Notice("删除成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	} else {
		flash.Notice("删除失败，课堂不存在")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	}
}

// @router /student/joincourse/:cid [get]
func (this *CourseController) JionCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	student := this.GetSession("account").(models.Student)
	flash := beego.NewFlash()
	if student.JionCourse(cid) {
		flash.Error("加入成功")
		flash.Store(&this.Controller)
		this.Redirect("/student", 302)
	} else {
		flash.Error("加入失败，您已经在课堂内请重试")
		flash.Store(&this.Controller)
		this.Redirect("/student", 302)
	}
}

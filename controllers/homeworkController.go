package controllers

import (
	"Teaching-assistant-system/models"

	"github.com/astaxie/beego"
)

type HomeworkController struct {
	BaseController
}

// @router /teacher/course/homework	[get]
func (this *HomeworkController) GetAddHomework() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}

	this.Data["course"] = this.GetSession("course").(models.Course)
	this.TplName = "teacherAddHomework.html"
}

// @router /teacher/course/homework [post]
func (this *HomeworkController) PostAddHomework() {
	flash := beego.NewFlash()
	course := this.GetSession("course").(models.Course)
	content := this.GetString("content")
	ddl := this.GetString("ddl")
	if ddl == "" {
		ddl = "0000-00-00"
	}

	teacher := this.GetSession("account").(models.Teacher)
	homework := models.Homework{Content: content, Ddl: ddl, Course: &course}
	if homework.AddHomework() {
		students := course.QueryStudents()
		for i := range students {
			models.NoticeTBuild(&teacher, &course, students[i], 1)
		}
		flash.Error("布置作业成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course/homework", 302)
		return
	} else {
		flash.Error("布置作业失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course/homework", 302)
		return
	}
}

// @router /teacher/course/homework/:hid [get]
func (this *HomeworkController) GetTeacherHomework() {
	this.Redirect("/intherow", 302)
}

// @router /student/course/homework/:hid [get]
func (this *HomeworkController) GetStudentHomework() {
	this.Redirect("/intherow", 302)
}

// @router /student/course/homework/:hid [post]
func (this *HomeworkController) PostStudentHomework() {
	this.Redirect("/intherow", 302)
}

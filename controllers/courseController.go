package controllers

import (
	"Teaching-assistant-system/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type CourseController struct {
	BaseController
}

// @router /student/:cid [get]
func (this *CourseController) StudentSelectCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	course := models.SearchCourse(cid)
	course.Student = course.QueryStudents()
	this.SetSession("course", *course)
	this.Redirect("/student/course", 302)
}

// @router /student/course [get]
func (this *CourseController) GetStudentCourse() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	}

	student := this.GetSession("account").(models.Student)
	this.Data["student"] = student

	course := this.GetSession("course").(models.Course)
	this.Data["course"] = course
	this.Data["students"] = course.Student

	if fileInfo, err := course.QueryFiles(); err == nil {
		this.Data["fileInfo"] = fileInfo
	} else {
		this.Data["notice"] = err
	}

	this.TplName = "studentcourse.html"
}

// @router /teacher/:cid [get]
func (this *CourseController) TeacherSelectCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	course := models.SearchCourse(cid)
	course.Student = course.QueryStudents()
	this.SetSession("course", course)
	this.Redirect("/teacher/course", 302)
}

// @router /teacher/course [get]
func (this *CourseController) GetTeacherCourse() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	}
	course := this.GetSession("course").(*models.Course)

	this.Data["course"] = course
	this.Data["students"] = course.Student

	if fileInfo, err := course.QueryFiles(); err == nil {
		this.Data["fileInfo"] = fileInfo
	} else {
		this.Data["notice"] = err
	}

	this.TplName = "teachercourse.html"
}

// @router /teacher/dropcourse/:cid [get]
func (this *CourseController) DeleteCourse() {
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	teacher := this.GetSession("account").(models.Teacher)
	flash := beego.NewFlash()
	if teacher.DeleteCourse(cid) {
		flash.Error("删除成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	} else {
		flash.Error("删除失败，课堂不存在")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	}
}

// @router /teacher/addstudent/:scid [get]
func (this *CourseController) AddStudent() {
	s := this.Ctx.Input.Param(":scid")
	ids := strings.Split(s, "&&")

	sid, _ := strconv.Atoi(ids[0])
	cid, _ := strconv.Atoi(ids[1])
	nid, _ := strconv.Atoi(ids[2])

	student := &models.Student{Id: sid}
	course := &models.Course{Id: cid}
	notice := &models.NoticeS{Id: nid}

	teacher := this.GetSession("account").(models.Teacher)
	flash := beego.NewFlash()
	if teacher.AddStudent(student, course, notice) {
		flash.Error("添加学生成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	} else {
		flash.Error("添加学生失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	}

}

// @router /teacher/refusestudent/:scid [get]
func (this *CourseController) RefuseStudent() {
	s := this.Ctx.Input.Param(":scid")
	ids := strings.Split(s, "&&")

	nid, _ := strconv.Atoi(ids[2])

	notice := &models.NoticeS{Id: nid}

	teacher := this.GetSession("account").(models.Teacher)
	flash := beego.NewFlash()
	if teacher.RefuseStudent(notice) {
		flash.Error("已拒绝该学生")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	} else {
		flash.Error("操作失败，请重试")
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
		flash.Error("请求已发送，请等待老师同意")
		flash.Store(&this.Controller)
		this.Redirect("/student", 302)
	} else {
		flash.Error("加入失败，您已经在课堂内，请重试")
		flash.Store(&this.Controller)
		this.Redirect("/student", 302)
	}
}

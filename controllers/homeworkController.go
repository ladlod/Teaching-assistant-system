package controllers

import (
	"Teaching-assistant-system/models"
	"os"
	"path"
	"strconv"

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
	this.TplName = "teacher/teacherAddHomework.html"
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
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}

	hid, _ := strconv.Atoi(this.Ctx.Input.Param(":hid"))
	homework := models.SearchHomeWork(hid)

	this.Data["homework"] = homework
	this.Data["course"] = this.GetSession("course").(models.Course)
	this.Data["student"] = this.GetSession("account").(models.Student)

	this.TplName = "student/studenthomework.html"
}

// @router /student/course/homework/:hid [post]
func (this *HomeworkController) PostStudentHomework() {
	hid_s := this.Ctx.Input.Param(":hid")
	hid, _ := strconv.Atoi(hid_s)

	flash := beego.NewFlash()
	f, h, err := this.GetFile("uploadfile")
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/student/course/homework/"+hid_s, 302)
		return
	}
	fileName := h.Filename
	f.Close()

	student := this.GetSession("account").(models.Student)

	if _, filestat := os.Stat("files/homework/" + hid_s + "/" + fileName); filestat == nil {
		flash.Error("文件已存在或文件名重复")
		flash.Store(&this.Controller)
		this.Redirect("/student/course/homework/"+hid_s, 302)
		return
	}

	if err := this.SaveToFile("uploadfile", path.Join("files/homework/"+hid_s, fileName)); err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/student/course/homework/"+hid_s, 302)
		return
	} else {
		student.SubmitHomework(hid)
		flash.Error("上传成功!")
		flash.Store(&this.Controller)
		this.Redirect("/student/course/homework/"+hid_s, 302)
		return
	}
}

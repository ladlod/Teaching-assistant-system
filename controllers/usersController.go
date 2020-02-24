package controllers

import (
	"Teaching-assistant-system/models"
	"strconv"

	"github.com/astaxie/beego"
)

type UsersController struct {
	BaseController
}

// @router /signin [get]
func (this *UsersController) GetSignin() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}
	this.TplName = "signin.html"
}

// @router /signin [post]
func (this *UsersController) PostSignin() {
	flash := beego.NewFlash()
	if this.GetString("usertype") == "teacher" {
		var teacher models.Teacher
		teacher.Account = this.GetString("account")
		teacher.Password = this.GetString("password")

		if teacher.Signin() {
			this.SetSession("account", teacher)
			this.SetSession("identity", "teacher")
			this.Redirect("/teacher", 302)
			return
		} else {
			flash.Notice("用户不存在或密码错误")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
			return
		}
	} else if this.GetString("usertype") == "student" {
		var student models.Student
		student.Account = this.GetString("account")
		student.Password = this.GetString("password")

		if student.Signin() {
			this.SetSession("account", student)
			this.SetSession("identity", "student")
			this.Redirect("/student", 302)
			return
		} else {
			flash.Notice("用户不存在或密码错误")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
			return
		}
	}
	return
}

// @router /signup [get]
func (this *UsersController) GetSignup() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = err
	}
	this.TplName = "signup.html"
}

// @router /signup [post]
func (this *UsersController) PostSignup() {
	flash := beego.NewFlash()
	if this.GetString("usertype") == "teacher" {
		var teacher models.Teacher
		teacher.Account = this.GetString("account")
		teacher.Name = this.GetString("username")
		teacher.Password = this.GetString("password")

		if teacher.Signup() {
			flash.Notice("注册成功")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
		} else {
			flash.Notice("账户已存在")
			flash.Store(&this.Controller)
			this.Redirect("/signup", 302)
		}
	} else if this.GetString("usertype") == "student" {
		var student models.Student
		student.Account = this.GetString("account")
		student.Name = this.GetString("username")
		student.Password = this.GetString("password")

		if student.Signup() {
			flash.Notice("注册成功")
			flash.Store(&this.Controller)
			this.Redirect("/signin", 302)
		} else {
			flash.Notice("账户已存在")
			flash.Store(&this.Controller)
			this.Redirect("/signup", 302)
		}
	}
}

// @router /student [get]
func (this *UsersController) GetStudent() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}
	if courseid, ok := flash.Data["notice"]; ok {
		cid, _ := strconv.Atoi(courseid)
		search_course := models.SearchCourse(cid)
		if search_course == nil {
			this.Data["notice"] = "课堂不存在"
		} else {
			this.Data["search_course"] = search_course
		}
	}
	var student = this.GetSession("account").(models.Student)

	courses := student.QueryCourse()
	notices := student.QueryNotice()

	this.Data["courses"] = courses
	this.Data["notices"] = notices
	this.Data["username"] = student.Name
	this.TplName = "student.html"
}

// @router /student [post]
func (this *UsersController) PostSearchCourse() {
	flash := beego.NewFlash()

	cid := this.GetString("course")
	if _, err := strconv.Atoi(cid); err == nil {
		flash.Notice(cid)
		flash.Store(&this.Controller)
		this.Redirect("/student", 302)
	} else {
		flash.Error("请输入课堂编号，格式为4位以内数字")
		flash.Store(&this.Controller)
		this.Redirect("/student", 302)
	}
}

// @router /teacher [get]
func (this *UsersController) GetTeacher() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}
	var teacher = this.GetSession("account").(models.Teacher)

	courses := teacher.QueryCourse()
	notices := teacher.QueryNotice()

	this.Data["courses"] = courses
	this.Data["notices"] = notices
	this.Data["username"] = teacher.Name
	this.TplName = "teacher.html"
}

// @router /teacher [post]
func (this *UsersController) PostCreateCourse() {
	flash := beego.NewFlash()
	var course models.Course
	course.Name = this.GetString("course")
	teacher := this.GetSession("account").(models.Teacher)
	if id, b := teacher.MakeCourse(&course); b {
		flash.Error("创建成功, 课堂ID为" + strconv.Itoa(id))
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	} else {
		flash.Error("创建失败，请稍后再试")
		flash.Store(&this.Controller)
		this.Redirect("/teacher", 302)
	}
}

// @router /student/setting [get]
func (this *UsersController) GetSetStudent() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	} else if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}

	var student = this.GetSession("account").(models.Student)

	this.Data["username"] = student.Name
	this.TplName = "setstudent.html"
}

// @router /student/setting [post]
func (this *UsersController) PostSetStudent() {
	flash := beego.NewFlash()
	var student = this.GetSession("account").(models.Student)
	student.Name = this.GetString("username")
	student.Password = this.GetString("password")

	if student.Change() {
		this.SetSession("account", student)
		this.SetSession("identity", "student")
		flash.Notice("修改成功")
		flash.Store(&this.Controller)
		this.Redirect("/student/setting", 302)
	} else {
		flash.Error("修改失败")
		flash.Store(&this.Controller)
		this.Redirect("/student/setting", 302)
	}
}

// @router /teacher/setting [get]
func (this *UsersController) GetSetTeacher() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	} else if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}

	var teacher = this.GetSession("account").(models.Teacher)

	this.Data["username"] = teacher.Name
	this.TplName = "setteacher.html"
}

// @router /teacher/setting [post]
func (this *UsersController) PostSetTeacher() {
	flash := beego.NewFlash()
	var teacher = this.GetSession("account").(models.Teacher)
	teacher.Name = this.GetString("username")
	teacher.Password = this.GetString("password")

	if teacher.Change() {
		this.SetSession("account", teacher)
		this.SetSession("identity", "teacher")
		flash.Notice("修改成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/setting", 302)
	} else {
		flash.Error("修改失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/setting", 302)
	}
}

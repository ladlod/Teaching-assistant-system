package controllers

import (
	"Teaching-assistant-system/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type ExamController struct {
	BaseController
}

// @router /teacher/course/addexam [get]
func (this *ExamController) GetAddExam() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}

	this.Data["course"] = this.GetSession("course").(models.Course)
	this.TplName = "teacher/teacherAddExam.html"
}

// @router /teacher/course/addexam [post]
func (this *ExamController) PostAddExam() {
	flash := beego.NewFlash()
	course := this.GetSession("course").(models.Course)
	chooseNum, _ := strconv.Atoi(this.GetString("chooseNum"))
	chooseScore, _ := strconv.ParseFloat(this.GetString("chooseScore"), 64)
	blankNum, _ := strconv.Atoi(this.GetString("blankNum"))
	blankScore, _ := strconv.ParseFloat(this.GetString("blankScore"), 64)
	answerNum, _ := strconv.Atoi(this.GetString("answerNum"))
	answerScore, _ := strconv.ParseFloat(this.GetString("answerScore"), 64)

	C1, _ := strconv.ParseFloat(this.GetString("C1"), 64)
	C2, _ := strconv.ParseFloat(this.GetString("C2"), 64)
	C3, _ := strconv.ParseFloat(this.GetString("C3"), 64)
	C4, _ := strconv.ParseFloat(this.GetString("C4"), 64)
	C5, _ := strconv.ParseFloat(this.GetString("C5"), 64)
	C6, _ := strconv.ParseFloat(this.GetString("C6"), 64)
	C7, _ := strconv.ParseFloat(this.GetString("C7"), 64)
	C8, _ := strconv.ParseFloat(this.GetString("C8"), 64)
	C9, _ := strconv.ParseFloat(this.GetString("C9"), 64)
	C10, _ := strconv.ParseFloat(this.GetString("C10"), 64)
	C11, _ := strconv.ParseFloat(this.GetString("C11"), 64)
	C12, _ := strconv.ParseFloat(this.GetString("C12"), 64)

	starttime := this.GetString("starttime")
	tmp := strings.Split(starttime, "T")
	starttime = tmp[0] + " " + tmp[1]

	time, _ := strconv.Atoi(this.GetString("time"))
	exam := models.Exam{
		ChooseNum:   chooseNum,
		ChooseScore: chooseScore,
		BlankNum:    blankNum,
		BlankScore:  blankScore,
		AnswerNum:   answerNum,
		AnswerScore: answerScore,
		C1:          C1,
		C2:          C2,
		C3:          C3,
		C4:          C4,
		C5:          C5,
		C6:          C6,
		C7:          C7,
		C8:          C8,
		C9:          C9,
		C10:         C10,
		C11:         C11,
		C12:         C12,
		StartTime:   starttime,
		Time:        time,
		Course:      &course,
	}

	if exam.AddExam() {
		flash.Error("发布考试成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course/addexam", 302)
		return
	} else {
		flash.Error("发布考试失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course/addexam", 302)
		return
	}
}

// @router /student/course/exam/:eid [get]
func (this *ExamController) GetStudentExam() {
	flash := beego.NewFlash()
	eid, _ := strconv.Atoi(this.Ctx.Input.Param(":eid"))
	exam := models.SearchExam(eid)
	if !exam.JudgeStartData() {
		flash.Error("考试未开始")
		flash.Store(&this.Controller)
		this.Redirect("/student/course", 302)
		return
	} else if exam.JudgeOutData() {
		this.Data["course"] = this.GetSession("course").(models.Course)
		this.Redirect("/intherow", 302)
		return
	} else {
		this.Data["course"] = this.GetSession("course").(models.Course)
		this.Data["exam"] = exam
		this.TplName = "student/joinexam.html"
		return
	}
}

// @router /student/joinexam/:eid [get]
func (this *ExamController) StudentJoinExam() {
	this.Redirect("/intherow", 302)
}

// @router /teacher/course/exam/:eid [get]
func (this *ExamController) GetTeacherExam() {
	this.Redirect("/intherow", 302)
}

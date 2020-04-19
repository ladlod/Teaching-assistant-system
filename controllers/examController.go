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
		student := this.GetSession("account").(models.Student)
		this.Data["score"] = exam.QueryScore(&student)
		this.TplName = "student/queryscore.html"
		return
	} else {
		this.Data["course"] = this.GetSession("course").(models.Course)
		this.Data["exam"] = exam
		this.TplName = "student/joinexam.html"
		return
	}
}

// @router /student/examing/:eid [get]
func (this *ExamController) StudentJoinExam() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}

	eid, _ := strconv.Atoi(this.Ctx.Input.Param(":eid"))
	exam := models.SearchExam(eid)
	student := this.GetSession("account").(models.Student)
	this.Data["student"] = student
	this.Data["problems"] = student.JoinExam(exam)
	this.TplName = "student/exam.html"
}

// @router /student/examing/:eid [post]
func (this *ExamController) StudentPostExam() {
	flash := beego.NewFlash()

	eid, _ := strconv.Atoi(this.Ctx.Input.Param(":eid"))
	exam := models.SearchExam(eid)
	student := this.GetSession("account").(models.Student)
	problems := student.JoinExam(exam)
	var answers []string

	for i := range problems {
		tmp := this.GetString(strconv.Itoa(problems[i].Id))
		answers = append(answers, tmp)
	}

	student.SubmitPaper(answers, exam)
	flash.Error("交卷成功，请等待教师批阅")
	flash.Store(&this.Controller)
	this.Redirect("/student/course", 302)
	return
}

// @router /teacher/course/exam/:eid [get]
func (this *ExamController) GetTeacherExam() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}

	this.Data["teacher"] = this.GetSession("account").(models.Teacher)
	eid, _ := strconv.Atoi(this.Ctx.Input.Param(":eid"))
	exam := models.SearchExam(eid)
	this.Data["students"] = exam.QueryAllStat()
	this.TplName = "teacher/exam.html"
}

// @router /teacher/course/reviewpaper/:esid [get]
func (this *ExamController) GetReviewPaper() {
	flash := beego.ReadFromRequest(&this.Controller)
	if err, ok := flash.Data["error"]; ok {
		this.Data["notice"] = err
	}

	teacher := this.GetSession("account").(models.Teacher)
	this.Data["teacher"] = teacher
	esid := this.Ctx.Input.Param(":esid")
	ids := strings.Split(esid, "&&")
	eid, _ := strconv.Atoi(ids[0])
	sid, _ := strconv.Atoi(ids[1])
	exam := models.SearchExam(eid)
	this.Data["exam"] = exam
	this.Data["students"] = exam.QueryAllStat()

	this.Data["problems"] = teacher.ReviewPaper(eid, sid)

	this.TplName = "teacher/exam.html"
}

// @router /teacher/course/reviewpaper/:esid [post]
func (this *ExamController) PostScore() {
	flash := beego.NewFlash()
	teacher := this.GetSession("account").(models.Teacher)
	esid := this.Ctx.Input.Param(":esid")
	ids := strings.Split(esid, "&&")
	eid, _ := strconv.Atoi(ids[0])
	sid, _ := strconv.Atoi(ids[1])
	problems := teacher.ReviewPaper(eid, sid)
	var score float64 = 0
	for i := range problems {
		if tmps := this.GetString(strconv.Itoa(problems[i].Id)); tmps != "" {
			tmp, _ := strconv.ParseFloat(tmps, 64)
			score += tmp
		}
	}
	if teacher.UpdateScore(eid, sid, score) {
		flash.Error("阅卷完毕，请批阅下一份试卷")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course/reviewpaper/"+esid, 302)
		return
	} else {
		flash.Error("阅卷失败，请重新批阅试卷")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course/reviewpaper/"+esid, 302)
		return
	}
}

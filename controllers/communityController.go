package controllers

import (
	"Teaching-assistant-system/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type CommunityController struct {
	BaseController
}

// @router /community [get]
func (this *CommunityController) Get() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	}

	if this.GetSession("identity").(string) == "student" {
		this.Data["identity"] = "student"
		student := this.GetSession("account").(models.Student)
		this.Data["user"] = student
		this.Data["notices"] = student.QueryAnswerNotice()
	} else if this.GetSession("identity").(string) == "teacher" {
		this.Data["identity"] = "teacher"
		teacher := this.GetSession("account").(models.Teacher)
		this.Data["user"] = teacher
		this.Data["notices"] = teacher.QueryAnswerNotice()
	}
	this.Data["queByTime"] = models.QueryQuestionsByTime()
	this.Data["queBySupport"] = models.QueryQuestionsBySupport()
	this.TplName = "community/community.html"
}

// @router /community/postquestion [get]
func (this *CommunityController) GetNewQuestion() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	}

	if this.GetSession("identity").(string) == "student" {
		this.Data["identity"] = "student"
		this.Data["user"] = this.GetSession("account").(models.Student)
	} else if this.GetSession("identity").(string) == "teacher" {
		this.Data["identity"] = "teacher"
		this.Data["user"] = this.GetSession("account").(models.Teacher)
	}
	this.TplName = "community/postnew.html"
}

// @router /community/postquestion [post]
func (this *CommunityController) PostNewQuestion() {
	flash := beego.NewFlash()

	var student *models.Student
	var teacher *models.Teacher
	if this.GetSession("identity").(string) == "student" {
		tmp := this.GetSession("account").(models.Student)
		student = &tmp
		teacher = nil
	} else if this.GetSession("identity").(string) == "teacher" {
		tmp := this.GetSession("account").(models.Teacher)
		student = nil
		teacher = &tmp
	}

	topic := this.GetString("topic")
	content := this.GetString("content")

	if models.AddQuestion(topic, content, student, teacher) {
		flash.Error("发布成功")
		flash.Store(&this.Controller)
		this.Redirect("/community/postquestion", 302)
		return
	} else {
		flash.Error("发布失败")
		flash.Store(&this.Controller)
		this.Redirect("/community/postquestion", 302)
		return
	}
}

// @router /community/:qid [get]
func (this *CommunityController) GetQuestion() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	}

	if this.GetSession("identity").(string) == "student" {
		this.Data["identity"] = "student"
		this.Data["user"] = this.GetSession("account").(models.Student)
	} else if this.GetSession("identity").(string) == "teacher" {
		this.Data["identity"] = "teacher"
		this.Data["user"] = this.GetSession("account").(models.Teacher)
	}

	qid, _ := strconv.Atoi(this.Ctx.Input.Param(":qid"))
	question := models.GetQuestionById(qid)
	this.Data["question"] = question

	this.Data["ansByTime"] = question.QueryAnswerByTime()
	this.Data["ansBySupport"] = question.QueryAnswerBySupport()

	this.TplName = "community/question.html"
}

// @router /community/:qid [post]
func (this *CommunityController) PostNewAnswer() {
	flash := beego.NewFlash()

	var student *models.Student
	var teacher *models.Teacher
	if this.GetSession("identity").(string) == "student" {
		tmp := this.GetSession("account").(models.Student)
		student = &tmp
		teacher = nil
	} else if this.GetSession("identity").(string) == "teacher" {
		tmp := this.GetSession("account").(models.Teacher)
		student = nil
		teacher = &tmp
	}

	content := this.GetString("content")
	qids := this.Ctx.Input.Param(":qid")
	qid, _ := strconv.Atoi(qids)
	question := models.GetQuestionById(qid)

	if question.AnswerQuestion(content, teacher, student) {
		flash.Error("评论成功")
		flash.Store(&this.Controller)
		this.Redirect("/community/"+qids, 302)
		return
	} else {
		flash.Error("评论失败")
		flash.Store(&this.Controller)
		this.Redirect("/community/"+qids, 302)
		return
	}
}

// @router /community/support/:qid [get]
func (this *CommunityController) SupportQuestion() {
	flash := beego.NewFlash()
	var student *models.Student
	var teacher *models.Teacher
	if this.GetSession("identity").(string) == "student" {
		tmp := this.GetSession("account").(models.Student)
		student = &tmp
		teacher = nil
	} else if this.GetSession("identity").(string) == "teacher" {
		tmp := this.GetSession("account").(models.Teacher)
		student = nil
		teacher = &tmp
	}

	qids := this.Ctx.Input.Param(":qid")
	qid, _ := strconv.Atoi(qids)
	question := models.GetQuestionById(qid)
	if !question.SupportQuestion(student, teacher) {
		flash.Error("您已经推荐过了")
		flash.Store(&this.Controller)
	}
	this.Redirect("/community/"+qids, 302)
	return
}

// @router /community/answer/support/:aid [get]
func (this *CommunityController) SupportAnswer() {
	flash := beego.NewFlash()
	var student *models.Student
	var teacher *models.Teacher
	if this.GetSession("identity").(string) == "student" {
		tmp := this.GetSession("account").(models.Student)
		student = &tmp
		teacher = nil
	} else if this.GetSession("identity").(string) == "teacher" {
		tmp := this.GetSession("account").(models.Teacher)
		student = nil
		teacher = &tmp
	}

	aid, _ := strconv.Atoi(this.Ctx.Input.Param(":aid"))
	answer := models.GetAnswerById(aid)
	if !answer.SupportAnswer(student, teacher) {
		flash.Error("您已经点亮过了")
		flash.Store(&this.Controller)
	}
	this.Redirect("/community/"+strconv.Itoa(answer.Question.Id), 302)
	return
}

// @router /community/myquestions [get]
func (this *CommunityController) GetMyQuestion() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	}

	if this.GetSession("identity").(string) == "student" {
		student := this.GetSession("account").(models.Student)
		this.Data["user"] = student
		this.Data["questions"] = student.QueryMyQuestion()
	} else if this.GetSession("identity").(string) == "teacher" {
		teacher := this.GetSession("account").(models.Teacher)
		this.Data["user"] = teacher
		this.Data["questions"] = teacher.QueryMyQuestion()
	}

	this.TplName = "community/myquestion.html"
}

// @router /community/delete/:qid [get]
func (this *CommunityController) DeleteQuestion() {
	flash := beego.NewFlash()
	qid, _ := strconv.Atoi(this.Ctx.Input.Param(":qid"))
	question := models.GetQuestionById(qid)

	if question.Delete() {
		flash.Error("删除成功")
		flash.Store(&this.Controller)
	} else {
		flash.Error("删除失败")
		flash.Store(&this.Controller)
	}

	this.Redirect("/community/myquestions", 302)
}

// @router /community/notice/:nqid [get]
func (this *CommunityController) DealNotice() {
	flash := beego.NewFlash()
	nqid := strings.Split(this.Ctx.Input.Param(":nqid"), "&&")
	nid, _ := strconv.Atoi(nqid[0])
	notice := &models.AnswerNotice{Id: nid}
	if notice.Delete() {
		this.Redirect("/community/"+nqid[1], 302)
		return
	} else {
		flash.Error("未知错误")
		flash.Store(&this.Controller)
		this.Redirect("/community", 302)
		return
	}
}

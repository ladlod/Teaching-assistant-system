package controllers

import (
	"Teaching-assistant-system/models"

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
		this.Data["user"] = this.GetSession("account").(models.Student)
	} else if this.GetSession("identity").(string) == "teacher" {
		this.Data["user"] = this.GetSession("account").(models.Teacher)
	}
	this.Data["queByTime"] = models.QueryQuestionsByTime()
	this.Data["queBySupport"] = models.QueryQuestionsBySupport()
	this.TplName = "intherow.html"
}

// @router /community/postquestion [get]
func (this *CommunityController) GetNewQuestion() {
	this.Redirect("/intherow", 302)
}

// @router /community/postquestion [post]
func (this *CommunityController) PostNewQuestion() {
	this.Redirect("/intherow", 302)
}

// @router /community/:qid [get]
func (this *CommunityController) GetQuestion() {
	this.Redirect("/intherow", 302)
}

// @router /community/:qid [post]
func (this *CommunityController) PostNewAnswer() {
	this.Redirect("/intherow", 302)
}

// @router /community/support/:qid [get]
func (this *CommunityController) SupportQuestion() {
	this.Redirect("/intherow", 302)
}

// @router /community/delete/:qid [get]
func (this *CommunityController) DeleteQuestion() {
	this.Redirect("/intherow", 302)
}

// @router /community/answer/support/:aid [get]
func (this *CommunityController) SupportAnswer() {
	this.Redirect("/intherow", 302)
}

// @router /community/answer/delete/:aid [get]
func (this *CommunityController) DeleteAnswer() {
	this.Redirect("/intherow", 302)
}

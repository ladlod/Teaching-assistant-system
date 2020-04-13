package controllers

import (
	"Teaching-assistant-system/models"
	"strconv"

	"github.com/astaxie/beego"
)

type ProblemsController struct {
	BaseController
}

// @router /teacher/problems [get]
func (this *ProblemsController) GetProblems() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	} else if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}

	this.Data["user"] = this.GetSession("account").(models.Teacher)
	this.TplName = "teacher/questionbank.html"
}

// @router /teacher/problems/:cid [get]
func (this *ProblemsController) GetProblemsByCapter() {
	flash := beego.ReadFromRequest(&this.Controller)
	if not, ok := flash.Data["error"]; ok {
		this.Data["notice"] = not
	} else if not, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = not
	}

	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	this.Data["user"] = this.GetSession("account").(models.Teacher)
	this.Data["problems"] = models.QueryProblems(cid)
	this.Data["cid"] = cid
	this.Data["chapter"] = models.Chapter[cid-1]

	this.TplName = "teacher/questionchapter.html"
}

// @router /teacher/problems/:cid [post]
func (this *ProblemsController) PostProblemsByCapter() {
	flash := beego.NewFlash()

	var problem models.Problem
	cid, _ := strconv.Atoi(this.Ctx.Input.Param(":cid"))
	typ, _ := strconv.Atoi(this.GetString("typ"))
	question := this.GetString("question")
	if typ == 1 {
		A := this.GetString("A")
		B := this.GetString("B")
		C := this.GetString("C")
		D := this.GetString("D")
		answer := this.GetString("chooseanswer")
		problem = models.Problem{Type: typ, Question: question, A: A, B: B, C: C, D: D, Answer: answer}
	} else {
		answer := this.GetString("blankanswer")
		problem = models.Problem{Type: typ, Question: question, Answer: answer}
	}
	if models.AddProblem(problem, cid) {
		flash.Error("添加题目成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/problems/"+strconv.Itoa(cid), 302)
	} else {
		flash.Error("添加题目失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/problems/"+strconv.Itoa(cid), 302)
		return
	}
}

// @router /teacher/problems/delete/:pid [get]
func (this *ProblemsController) DeleteProblem() {
	flash := beego.NewFlash()

	pid, _ := strconv.Atoi(this.Ctx.Input.Param(":pcid"))
	probelm := models.Problem{Id: pid}
	if cid, err := probelm.Delete(); err == true {
		flash.Error("删除题目成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/problems/"+strconv.Itoa(cid), 302)
		return
	} else {
		flash.Error("删除题目失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/problems/"+strconv.Itoa(cid), 302)
		return
	}
}

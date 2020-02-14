package controllers

import (
	"github.com/astaxie/beego/session"
)

type MainController struct {
	BaseController
}

var GlobalSessions *session.Manager //全局Session

// SessionInit 初始化Session
func Init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "tasid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          true,
		ProviderConfig:  "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()
}

// @router / [get]
func (this *MainController) Get() {
	this.Redirect("/signin", 302)
}

// @router /intherow [get]
func (this *MainController) InTheRow() {
	this.TplName = "intherow.html"
}

// @router /help [get]
func (this *MainController) GetHelp() {
	this.Redirect("/intherow", 302)
}

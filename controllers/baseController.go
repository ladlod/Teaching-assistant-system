package controllers

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["StartTime"] = time.Now().Format("2006-01-02 15:04:05")
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

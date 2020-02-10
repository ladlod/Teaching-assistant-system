package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	this.Data["StartTime"] = time.Now()
}

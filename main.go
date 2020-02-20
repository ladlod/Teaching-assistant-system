package main

import (
	"Teaching-assistant-system/models"
	"Teaching-assistant-system/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
	beego.InsertFilter("/*", beego.BeforeRouter, routers.FilterUser)
	beego.InsertFilter("/student/*", beego.BeforeRouter, routers.FilterStudent)
	beego.InsertFilter("/teacher/*", beego.BeforeRouter, routers.FilterTeacher)
}

func main() {
	beego.Run()
}

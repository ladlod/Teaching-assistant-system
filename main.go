package main

import (
	"Teching-assistant-system/models"
	_ "Teching-assistant-system/routers"

	"github.com/astaxie/beego"
)

func init() {
	models.Init()
}

func main() {
	beego.Run()
}

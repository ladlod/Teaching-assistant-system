package routers

import (
	"Teching-assistant-system/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.ClassController{})
	beego.Include(&controllers.UsersController{})
}

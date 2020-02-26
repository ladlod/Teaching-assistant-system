package controllers

import (
	"Teaching-assistant-system/models"
	"path"
	"strconv"

	"github.com/astaxie/beego"
)

type FileController struct {
	BaseController
}

// @router /teacher/course [post]
func (this *FileController) PostUploadFile() {
	flash := beego.NewFlash()
	f, h, err := this.GetFile("uploadfile")
	if err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/teacher/upload", 302)
		return
	}
	fileName := h.Filename
	f.Close()
	course := this.GetSession("course").(*models.Course)
	if err := this.SaveToFile("uploadfile", path.Join("files/"+strconv.Itoa(course.Id), fileName)); err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
		return
	} else {
		flash.Error("上传成功!")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
		return
	}
}

// @router /student/download/:filename [get]
func (this *FileController) GetDownloadFile() {

}

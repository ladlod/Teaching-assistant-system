package controllers

import (
	"Teaching-assistant-system/models"
	"os"
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
	course := this.GetSession("course").(models.Course)

	if _, filestat := os.Stat("files/courseware/" + strconv.Itoa(course.Id) + "/" + fileName); filestat == nil {
		flash.Error("文件已存在或文件名重复")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
		return
	}

	if err := this.SaveToFile("uploadfile", path.Join("files/courseware/"+strconv.Itoa(course.Id), fileName)); err != nil {
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

// @router /download/courseware/:filename [get]
func (this *FileController) GetDownloadFile() {
	filename := this.Ctx.Input.Param(":filename")
	course := this.GetSession("course").(models.Course)
	d_url := "files/courseware/" + strconv.Itoa(course.Id) + "/" + filename
	this.Ctx.Output.Download(d_url)
}

// @router /teacher/delete/:filename [get]
func (this *FileController) GetDeleteFile() {
	flash := beego.NewFlash()
	filename := this.Ctx.Input.Param(":filename")
	course := this.GetSession("course").(models.Course)
	d_url := "files/courseware/" + strconv.Itoa(course.Id) + "/" + filename
	if err := os.RemoveAll(d_url); err == nil {
		flash.Error("删除成功")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
	} else {
		flash.Error("删除失败")
		flash.Store(&this.Controller)
		this.Redirect("/teacher/course", 302)
	}
}

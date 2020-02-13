package models

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

var T *Teacher = nil //存储当前用户信息
var S *Student = nil
var identity int = 0 //存储当前用户状态
var C *Course = nil  //存储当前课程信息

// Init 初始化
func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:112358@/test?charset=utf8")
	orm.RegisterModel(new(Teacher), new(Student), new(Course))
}

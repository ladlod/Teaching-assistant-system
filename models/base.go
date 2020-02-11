package models

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

// Init 初始化
func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("users", "mysql", "root:112358@/teching_assistant?charset=utf8")
	orm.RegisterModel(new(Student), new(Teacher), new(Course))
}

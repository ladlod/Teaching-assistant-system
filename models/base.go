package models

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
)

// Init 初始化
func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:112358@/test?charset=utf8")
	orm.RegisterModel(new(Teacher), new(Student), new(Course))
}

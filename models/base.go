package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

// Init 初始化
func Init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:112358@/test?charset=utf8")
	orm.RegisterModel(new(Teacher), new(Student), new(Course), new(NoticeS), new(NoticeT), new(Homework))
	orm.RunSyncdb("default", false, true) //自动建表
	O = orm.NewOrm()
	O.Using("default")
}

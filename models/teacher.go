package models

import "github.com/astaxie/beego/orm"

type Teacher struct {
	Id       int
	Account  string
	Name     string
	Password string
	//Courses  []*Course `orm:"reverse(many)"`
}

// Signup 用户注册
func (teacher *Teacher) Signup() bool {
	orm := orm.NewOrm()
	orm.Using("default")

	if err := orm.Read(teacher, "Name"); err == nil {
		return false
	}
	if _, err := orm.Insert(teacher); err == nil {
		return true
	} else {
		return false
	}
}

// Signin 用户登录
func (teacher *Teacher) Signin() bool {
	orm := orm.NewOrm()
	orm.Using("default")

	if err := orm.Read(teacher, "Account", "Password"); err == nil {
		identity = 1
		T = teacher
		return true
	}
	return false
}

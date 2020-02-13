package models

import "github.com/astaxie/beego/orm"

type Student struct {
	Id       int
	Account  string
	Name     string
	Password string
	//Courses  []*Course `orm:"reverse(many)"`
}

// Signup 用户注册
func (student *Student) Signup() bool {
	orm := orm.NewOrm()
	orm.Using("default")

	if err := orm.Read(student, "Name"); err == nil {
		return false
	}
	if _, err := orm.Insert(student); err == nil {
		return true
	} else {
		return false
	}
}

// Signin 用户登录
func (student *Student) Signin() bool {
	orm := orm.NewOrm()
	orm.Using("default")

	if err := orm.Read(student, "Account", "Password"); err == nil {
		identity = 2
		S = student
		return true
	}
	return false
}

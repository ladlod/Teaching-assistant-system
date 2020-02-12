package models

import "github.com/astaxie/beego/orm"

// User 用户
type User struct {
	Id       int
	Name     string
	Password string
	Identity int
	Courses  []*Course
}

// Signup 用户注册
func (user *User) Signup() bool {
	orm := orm.NewOrm()
	orm.Using("default")

	if err := orm.Read(user, "Name"); err == nil {
		return false
	}
	if _, err := orm.Insert(user); err == nil {
		return true
	} else {
		return false
	}
}

// Signin 用户登录
func (user *User) Signin() bool {
	orm := orm.NewOrm()
	orm.Using("default")

	if err := orm.Read(user, "Name", "Password"); err == nil {
		return true
	}
	return false
}

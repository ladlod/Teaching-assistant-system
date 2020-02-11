package models

import "github.com/astaxie/beego/orm"

// User 所有用户
type User struct {
	Id       int
	Name     string
	Password string
}

// Signup 用户注册
func (user *User) Signup() (bool, error) {
	orm := orm.NewOrm()
	orm.Using("users")

	if err := orm.Read(user, "Name"); err == nil {
		return false, nil
	}
	if _, err := orm.Insert(user); err == nil {
		return true, nil
	} else {
		return false, nil
	}
}

// Signin 用户登录
func (user *User) Signin() bool {
	orm := orm.NewOrm()
	orm.Using("users")

	if err := orm.Read(user, "Name", "Password"); err == nil {
		return true
	}
	return false
}

// Student 学生，继承自用户
type Student struct {
	User,
	Courses []Course
}

// Teacher 老师，继承自用户
type Teacher struct {
	User,
	Courses []Course
}

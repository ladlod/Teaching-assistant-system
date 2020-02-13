package models

import "github.com/astaxie/beego/orm"

/*Student 学生
属性说明：
	Id 顺序产生的编号
	Account 账户名，不允许重复
	Name 学生姓名
	PassWord 登录密码
方法说明：
	Signup 注册
	Signin 登录
	QueryCourse 查询我选择的课堂
	JoinCourse 加入课堂
*/
type Student struct {
	Id       int
	Account  string
	Name     string
	Password string
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

// QueryCourse 查询我加入的课堂
func (student *Student) QueryCourse() {
}

// JionCourse 加入课堂
func (student *Student) JionCourse(course *Course) bool {
	return course.Addstudent(student)
}

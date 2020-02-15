package models

import "github.com/astaxie/beego/orm"

/*Teacher 老师
属性说明：
	Id 顺序产生的编号
	Account 账户名，不允许重复
	Name 学生姓名
	PassWord 登录密码
方法说明：
	Signup 注册
	Signin 登录
	QueryCourse 查询我创建的课堂
	MakeCourse 创建课堂
	DeleteCourse 删除课堂
*/
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
		return true
	}
	return false
}

// Change 修改账户信息
func (teacher *Teacher) Change() bool {
	return true
}

// QueryCourse 查询我创建的课堂
func (teacher *Teacher) QueryCourse() {
}

// MakeCourse 创建课堂
func (teacher *Teacher) MakeCourse(course *Course) bool {
	return course.MakeCourse()
}

// DeleteCourse 删除课堂
func (teacher *Teacher) DeleteCourse(course *Course) bool {
	return course.DeleteCourse()
}

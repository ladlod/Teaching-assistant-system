package models

/*Student 学生
属性说明：
	Id 顺序产生的编号
	Account 账户名，不允许重复
	Name 学生姓名
	PassWord 登录密码
方法说明：
	Signup 注册
	Signin 登录
	Change 修改用户信息
	QueryCourse 查询我选择的课堂
	JoinCourse 加入课堂
*/
type Student struct {
	Id       int `orm:"column(id);auto"`
	Account  string
	Name     string
	Password string
	Course   []*Course `orm:"rel(m2m)"`
}

// Signup 用户注册
func (student *Student) Signup() bool {
	if err := O.Read(student, "Name"); err == nil {
		return false
	}
	if _, err := O.Insert(student); err == nil {
		return true
	} else {
		return false
	}
}

// Signin 用户登录
func (student *Student) Signin() bool {
	if err := O.Read(student, "Account", "Password"); err == nil {
		return true
	}
	return false
}

// Change 修改用户信息
func (student *Student) Change() bool {
	if _, err := O.Update(student, "Name", "Password"); err == nil {
		return true
	}
	return false
}

// JionCourse 加入课堂
func (student *Student) JionCourse(cid int) bool {
	course := Course{Id: cid}
	return course.Addstudent(student)
}

// QueryCourse 查询我选择的课堂
func (student *Student) QueryCourse() []*Course {
	var courses []*Course
	O.QueryTable("course").Filter("Student__Student__id", student.Id).All(&courses)
	for i, _ := range courses {
		O.Read(courses[i].Teacher)
	}
	return courses
}

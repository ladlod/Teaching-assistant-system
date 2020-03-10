package models

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
	QueryNotice 查询我的通知
	MakeCourse 创建课堂
	DeleteCourse 删除课堂
	AddStudent 添加学生
	RefuseStudent 拒绝学生
	QueryMyQuestion 查询我发的帖子
	QueryMyAnswer 查询我发的回帖
*/
type Teacher struct {
	Id       int    `orm:"column(id);auto"`
	Account  string `orm:"unique"`
	Name     string
	Password string
	Courses  []*Course `orm:"reverse(many)"`
}

func (teacher *Teacher) TableUnique() [][]string {
	return [][]string{
		[]string{"account"},
	}
}

// Signup 用户注册
func (teacher *Teacher) Signup() bool {
	if _, err := O.Insert(teacher); err == nil {
		return true
	} else {
		return false
	}
}

// Signin 用户登录
func (teacher *Teacher) Signin() bool {
	if err := O.Read(teacher, "Account", "Password"); err == nil {
		return true
	}
	return false
}

// Change 修改账户信息
func (teacher *Teacher) Change() bool {
	if _, err := O.Update(teacher, "Name", "Password"); err == nil {
		return true
	}
	return false
}

// QueryCourse 查询我创建的课堂
func (teacher *Teacher) QueryCourse() []*Course {
	var courses []*Course
	_, err := O.QueryTable("course").Filter("teacher_id", teacher.Id).All(&courses)
	if err == nil {
		return courses
	}
	return nil
}

// MakeCourse 创建课堂
func (teacher *Teacher) MakeCourse(course *Course) (int, bool) {
	course.Teacher = teacher
	return course.MakeCourse()
}

// DeleteCourse 删除课堂
func (teacher *Teacher) DeleteCourse(cid int) bool {
	course := Course{Id: cid}
	return course.DeleteCourse()
}

// AddStudent 添加学生
func (teacher *Teacher) AddStudent(student *Student, course *Course, notice *NoticeS) bool {
	return course.Addstudent(student) && notice.DeleteNotice()
}

// RefuseStudent 拒绝学生
func (teacher *Teacher) RefuseStudent(notice *NoticeS) bool {
	return notice.DeleteNotice()
}

// QueryNotice 查询我的通知
func (teacher *Teacher) QueryNotice() []*NoticeS {
	var notices []*NoticeS
	O.Raw("select * from notice_s where course_id in (select id from course where teacher_id = ?)", teacher.Id).QueryRows(&notices)

	return notices
}

// QueryMyQuestion 查询我发的帖子
func (teacher *Teacher) QueryMyQuestion() []*Question {
	var questions []*Question
	O.QueryTable("question").Filter("teacher_id", teacher.Id).OrderBy("-Id").All(&questions)
	return questions
}

// QueryMyAnswer 查询我发的回帖
func (teacher *Teacher) QueryMyAnswer() []*Answer {
	var answers []*Answer
	O.QueryTable("answer").Filter("teacher_id", teacher.Id).OrderBy("-Id").All(&answers)
	return answers
}

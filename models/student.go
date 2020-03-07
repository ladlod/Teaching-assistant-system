package models

import "time"

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
	QueryNotice 查询我的通知
	JoinCourse 加入课堂
	SubmitHomework 提交作业
	ClockIn 签到
*/
type Student struct {
	Id        int    `orm:"column(id);auto"`
	Account   string `orm:"unique"`
	Name      string
	Password  string
	Course    []*Course `orm:"rel(m2m);rel_through(Teaching-assistant-system/models.StudentCourses)"`
	ClockStat string    `orm:"-"`
}

func (student *Student) TableUnique() [][]string {
	return [][]string{
		[]string{"account"},
	}
}

// Signup 用户注册
func (student *Student) Signup() bool {
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
	n, _ := O.QueryTable("student_courses").Filter("course_id", cid).Filter("student_id", student.Id).Count()
	if n != 0 {
		return false
	}
	course := Course{Id: cid}

	return NoticeSBuild(student, &course)
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

//QueryNotice 查询我的通知
func (student *Student) QueryNotice() []*NoticeT {
	var notices []*NoticeT
	O.QueryTable("notice_t").Filter("student_id", student.Id).All(&notices)
	//O.Raw("select * from notice_t where course_id in (select course_id from student_courses where student_id = ?)", student.Id).QueryRows(&notices)

	return notices
}

// SubmitHomework 提交作业
func (student *Student) SubmitHomework(hid int, filename string) bool {
	var s_h StudentHomework
	O.QueryTable("student_homework").Filter("student_id", student.Id).Filter("homework_id", hid).One(&s_h)
	s_h.Stat = "已提交"
	s_h.Filename = filename
	s_h.Time = time.Now().Format("2006-01-02 15:04:05")
	if _, err := O.Update(&s_h, "Stat", "Filename", "Time"); err == nil {
		return true
	}

	return false
}

// ClockIn 签到
func (student *Student) Clockin(password string, course *Course) bool {
	timeTmp := "2006-01-02 15:04:05"
	ddl, _ := time.ParseInLocation(timeTmp, course.Ddl, time.Local)
	if time.Now().After(ddl) || password != course.Password {
		return false
	}

	studentcourse := StudentCourses{Student: student, Course: course}
	O.Read(&studentcourse, "student_id", "course_id")
	studentcourse.Stat = "已签到"
	if _, err := O.Update(&studentcourse, "Stat"); err != nil {
		return false
	}

	return true
}

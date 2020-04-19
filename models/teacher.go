package models

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

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
	QueryAnswerNotice 查询我的回帖通知
	ReviewPaper 阅卷
	UpdateScore 给学生记录成绩
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

// QueryAnswerNotice 查询我的回帖通知
func (teacher *Teacher) QueryAnswerNotice() []*AnswerNotice {
	var notices []*AnswerNotice
	O.QueryTable("answer_notice").Filter("Question__Teacher__Id", teacher.Id).OrderBy("-Id").All(&notices)
	return notices
}

// ReviewPaper 阅卷
func (teacher *Teacher) ReviewPaper(eid int, sid int) []*Problem {
	var problems []*Problem
	paper, _ := os.Open("files/exam/" + strconv.Itoa(eid) + "/paper" + strconv.Itoa(sid))
	defer paper.Close()
	submitpaper, _ := os.Open("files/exam/" + strconv.Itoa(eid) + "/submit" + strconv.Itoa(sid))
	defer submitpaper.Close()
	r1 := bufio.NewReader(paper)
	r2 := bufio.NewReader(submitpaper)
	for {
		ids, _, err := r1.ReadLine()
		if err == io.EOF {
			break
		}
		id, _ := strconv.Atoi(string(ids))
		answer, _, _ := r2.ReadLine()
		problem := &Problem{Id: id}
		O.Read(problem)

		problem.Answer_student = string(answer)
		problems = append(problems, problem)
	}

	return problems
}

// UpdateScore 给学生记录成绩
func (teacher *Teacher) UpdateScore(eid int, sid int, score float64) bool {
	var se StudentExam
	O.QueryTable("student_exam").Filter("exam_id", eid).Filter("student_id", sid).One(&se)
	if se.Stat == "已评分" {
		return false
	}
	se.Score += score
	se.Stat = "已评分"
	if _, err := O.Update(&se, "score", "stat"); err == nil {
		return true
	}
	return false
}

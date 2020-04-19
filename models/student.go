package models

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"time"
)

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
	QuitCourse 退出课堂
	SubmitHomework 提交作业
	ClockIn 签到
	QueryMyQuestion 查询我发的帖子
	QUeryFNotice 查询我的回帖通知
	JoinExam 参加考试读取试题
	SubmitPaper 交卷
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

// QuitCourse 退出课堂
func (student *Student) QuitCourse(cid int) bool {
	course := &Course{Id: cid}
	return course.DeleteStudent(student)
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

// QueryMyQuestion 查询我发的帖子
func (student *Student) QueryMyQuestion() []*Question {
	var questions []*Question
	O.QueryTable("question").Filter("student_id", student.Id).OrderBy("-Id").All(&questions)
	return questions
}

// QueryAnswerNotice 查询我的回帖通知
func (student *Student) QueryAnswerNotice() []*AnswerNotice {
	var notices []*AnswerNotice
	O.QueryTable("answer_notice").Filter("Question__Student__Id", student.Id).OrderBy("-Id").All(&notices)
	return notices
}

// JoinExam 参加考试读取试题
func (student *Student) JoinExam(exam *Exam) []*Problem {
	var problems []*Problem
	paper, _ := os.Open("files/exam/" + strconv.Itoa(exam.Id) + "/paper" + strconv.Itoa(student.Id))
	defer paper.Close()
	r := bufio.NewReader(paper)
	for {
		ids, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		id, _ := strconv.Atoi(string(ids))
		problem := &Problem{Id: id}
		O.Read(problem)
		problems = append(problems, problem)
	}

	return problems
}

// SubmitPaper 交卷
func (student *Student) SubmitPaper(answers []string, exam *Exam) {
	paper, _ := os.Open("files/exam/" + strconv.Itoa(exam.Id) + "/paper" + strconv.Itoa(student.Id))
	submitpaper, _ := os.Create("files/exam/" + strconv.Itoa(exam.Id) + "/submit" + strconv.Itoa(student.Id))
	defer submitpaper.Close()
	r := bufio.NewReader(paper)
	var score float64 = 0
	for i := range answers {
		ids, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		id, _ := strconv.Atoi(string(ids))
		problem := &Problem{Id: id}
		O.Read(problem)
		if problem.Type == 1 && problem.Answer == answers[i] {
			score += exam.ChooseScore
		}
		submitpaper.WriteString(answers[i])
		if i != len(answers)-1 {
			submitpaper.WriteString("\n")
		}
	}
	var se StudentExam
	O.QueryTable("student_exam").Filter("student_id", student.Id).Filter("exam_id", exam.Id).One(&se)
	se.Score = score
	se.Stat = "已交卷未评分"
	O.Update(&se, "stat", "score")
	defer paper.Close()
}

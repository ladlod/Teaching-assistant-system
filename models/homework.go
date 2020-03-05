package models

import (
	"os"
	"strconv"
	"time"
)

// 查找作业
func SearchHomeWork(hid int) *Homework {
	var homework Homework
	err := O.QueryTable("homework").Filter("id", hid).One(&homework)
	if err == nil {
		return &homework
	}
	return nil
}

/* Homework 作业
属性说明：
	Id 顺序产生的编号
	Content 作业内容
	Ddl 作业截止日期
	Course  发布作业的课堂
	Stat 学生的完成状态，临时存储数据，不在数据库中存储
方法说明：
	AddHomework 布置作业
	QueryStat 查询某个学生的完成状态
	QueryStudentsStat 查询全部学生的完成状态
	QueryFiles 查询作业文件
	JudgeOutdata 判断作业是否截止
*/
type Homework struct {
	Id      int    `orm:"column(id);auto"`
	Content string `orm:"size(1024)"`
	Ddl     string
	Course  *Course `orm:"rel(fk)"`
	Stat    string  `orm:"-"`
}

// AddHomework 布置作业
func (h *Homework) AddHomework() bool {
	if _, err := O.Insert(h); err == nil {
		path := "files/homework/" + strconv.Itoa(h.Id)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return false
		}
		students := h.Course.QueryStudents()
		for i := range students {
			s_h := &StudentHomework{Stat: "未提交", Homework: h, Student: students[i]}
			if _, err := O.Insert(s_h); err != nil {
				return false
			}
		}
		return true
	}
	return false
}

// QueryStat 查询某个学生的完成状态
func (h *Homework) QueryStat(student *Student) {
	var sh StudentHomework
	O.QueryTable("student_homework").Filter("homework_id", h.Id).Filter("student_id", student.Id).One(&sh)
	h.Stat = sh.Stat
}

// QueryStudentsStat 查询全部学生的完成状态
func (h *Homework) QueryStudentsStat() []*StudentHomework {
	var students []*StudentHomework
	O.QueryTable("student_homework").Filter("homework_id", h.Id).All(&students)
	for i := range students {
		O.Read(students[i].Student)
	}

	return students
}

// QueryFiles 查询作业文件
func (h *Homework) QueryFiles() ([]os.FileInfo, error) {
	fd, err := os.Open("files/homework/" + strconv.Itoa(h.Id))
	if err != nil {
		return nil, err
	}
	return fd.Readdir(0)
}

// JudgeOutdata 判断作业是否截止
func (h *Homework) JudgeOutdata() bool {
	if h.Ddl == "0000-00-00" {
		return true
	}

	timeTmp := "2006-01-02"
	ddl, _ := time.ParseInLocation(timeTmp, h.Ddl, time.Local)

	if time.Now().Before(ddl) {
		return true
	}

	return false
}

type StudentHomework struct {
	Id       int `orm:"column(id);auto"`
	Stat     string
	Filename string    `orm:"null"`
	Time     string    `orm:"null"`
	Homework *Homework `orm:"rel(fk)"`
	Student  *Student  `orm:"rel(fk)"`
}

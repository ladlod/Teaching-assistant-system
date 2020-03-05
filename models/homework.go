package models

import (
	"os"
	"strconv"
)

func SearchHomeWork(hid int) *Homework {
	var homework Homework
	err := O.QueryTable("homework").Filter("id", hid).One(&homework)
	if err == nil {
		return &homework
	}
	return nil
}

type Homework struct {
	Id      int `orm:"column(id);auto"`
	Content string
	Ddl     string
	Course  *Course `orm:"rel(fk)"`
	Stat    string  `orm:"-"`
}

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

func (h *Homework) QueryStat(student *Student) {
	var sh StudentHomework
	O.QueryTable("student_homework").Filter("homework_id", h.Id).Filter("student_id", student.Id).One(&sh)
	h.Stat = sh.Stat
}

type StudentHomework struct {
	Id       int `orm:"column(id);auto"`
	Stat     string
	Homework *Homework `orm:"rel(fk)"`
	Student  *Student  `orm:"rel(fk)"`
}

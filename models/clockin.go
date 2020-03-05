package models

import "time"

type Clockin struct {
	Id       int `orm:"auto"`
	Password string
	Ddl      string
	Course   *Course `orm:"rel(fk)"`
}

func NewClockin(password string, cid int) bool {
	start := time.Now()
	end := start.Add(time.Minute * 5)

	ddl := end.Format("2006-01-02 15:04:05")
	course := SearchCourse(cid)

	clockin := &Clockin{Password: password, Ddl: ddl, Course: course}
	O.QueryTable("clockin").Filter("course_id", cid).Delete()
	if _, err := O.Insert(clockin); err == nil {
		students := course.QueryStudents()
		for i := range students {
			O.QueryTable("student_clockin").Filter("student_id", students[i].Id).Filter("clockin_id", clockin.Id).Delete()
			studentclockin := &StudentClockin{Stat: "未签到", Student: students[i], Clockin: clockin}
			O.Insert(studentclockin)
		}
		return true
	}
	return false
}

func SearchClockin(cid int) bool {
	if O.QueryTable("clockin").Filter("course_id", cid).Exist() {
		return true
	}
	return false
}

type StudentClockin struct {
	Id      int      `orm:"auto"`
	Stat    string   `orm:"null"`
	Student *Student `orm:"rel(fk)"`
	Clockin *Clockin `orm:"rel(fk)"`
}

package models

import (
	"os"
	"strconv"
)

// SearchCourse 查找课堂
func SearchCourse(cid int) *Course {
	var course Course
	err := O.QueryTable("course").Filter("id", cid).One(&course)
	if err == nil {
		O.Read(course.Teacher)
		return &course
	}
	return nil
}

/* Course 课堂
属性说明：
	Id 顺序产生的编号
	Name 课堂名
方法说明：
	MakeCourse 创建课堂
	DeleteCourse 删除课堂
	QueryStudents 查询选择这门课的学生
	Addstudent 添加学生
	QueryFiles 查询课件
	QueryHomework 查询作业
*/
type Course struct {
	Id      int `orm:"column(id);auto"`
	Name    string
	Teacher *Teacher   `orm:"rel(fk)"`
	Student []*Student `orm:"reverse(many)"`
}

// MakeCourse 创建课堂
func (course *Course) MakeCourse() (int, bool) {
	if _, err := O.Insert(course); err == nil {
		path := "files/courseware/" + strconv.Itoa(course.Id)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return 0, false
		}
		return course.Id, true
	}
	return 0, false
}

// DeleteCourse 删除课堂
func (course *Course) DeleteCourse() bool {
	path := "files/courseware/" + strconv.Itoa(course.Id)
	if err := os.RemoveAll(path); err != nil {
		return false
	}
	if _, err := O.Delete(course); err == nil {
		return true
	}
	return false
}

// QueryStudents 查询选择这门课的学生
func (course *Course) QueryStudents() []*Student {
	var students []*Student
	//O.Raw("select * from student where id in (select student_id from student_courses where course_id = ?)", course.Id).QueryRows(&students)
	O.QueryTable("student").Filter("Course__Course__id", course.Id).All(&students)

	for i := range students {
		clock := &Clockin{Course: course}
		O.Read(clock, "course_id")
		sclock := &StudentClockin{Student: students[i], Clockin: clock}
		O.Read(sclock, "student_id", "clockin_id")
		students[i].ClockStat = sclock.Stat
	}

	return students
}

// Addstudent 添加学生
func (course *Course) Addstudent(student *Student) bool {
	ids := []int{student.Id, course.Id}

	if _, err := O.Raw("Insert into student_courses values(null, ?, ?)", ids).Exec(); err == nil {
		return true
	} else {
		return false
	}
}

// QueryFiles 查询课件
func (course *Course) QueryFiles() ([]os.FileInfo, error) {
	fd, err := os.Open("files/courseware/" + strconv.Itoa(course.Id))
	if err != nil {
		return nil, err
	}
	return fd.Readdir(0)
}

// QueryHomework 查询作业
func (course *Course) QueryHomework() []*Homework {
	var homeworks []*Homework
	_, err := O.QueryTable("homework").Filter("course_id", course.Id).All(&homeworks)
	if err == nil {
		return homeworks
	}

	return nil
}

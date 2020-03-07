package models

import (
	"fmt"
	"os"
	"strconv"
	"time"
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
	Teacher 课堂教师
	Teacher 课堂内学生
	Password 签到密码
	Ddl 签到截止时间
方法说明：
	MakeCourse 创建课堂
	DeleteCourse 删除课堂
	QueryStudents 查询选择这门课的学生
	Addstudent 添加学生
	DeleteStudent 删除学生
	QueryFiles 查询课件
	QueryHomework 查询作业
	AddClockin 发起签到
	JudgeClockin 判断是否有签到发起
*/
type Course struct {
	Id       int `orm:"column(id);auto"`
	Name     string
	Teacher  *Teacher   `orm:"rel(fk)"`
	Student  []*Student `orm:"reverse(many)"`
	Password string     `orm:"null"`
	Ddl      string     `orm:"null"`
}

func (course *Course) TableIndex() [][]string {
	return [][]string{
		[]string{"teacher"},
	}
}

type StudentCourses struct {
	Id      int      `orm:"auto"`
	Stat    string   `orm:"null"`
	Student *Student `orm:"rel(fk)"`
	Course  *Course  `orm:"rel(fk)"`
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
		studentcourse := &StudentCourses{Student: students[i], Course: course}
		O.Read(studentcourse, "student_id", "course_id")
		students[i].ClockStat = studentcourse.Stat
	}

	return students
}

// Addstudent 添加学生
func (course *Course) Addstudent(student *Student) bool {
	ids := []int{student.Id, course.Id}

	if _, err := O.Raw("Insert into student_courses values(null, ?, ?)", ids).Exec(); err == nil {
		return true
	}
	return false
}

// DeleteStudent 删除学生
func (course *Course) DeleteStudent(student *Student) bool {
	if _, err := O.QueryTable("student_courses").Filter("student_id", student.Id).Filter("course_id", course.Id).Delete(); err == nil {
		return true
	}
	return false
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
	SortHomework(homeworks)
	if err == nil {
		return homeworks
	}

	return nil
}

// SortHomework 对作业进行排序
func SortHomework(homeworks []*Homework) {
	hmap := make(map[*Homework]float64)
	timeTmp := "2006-01-02 15:04:05"
	for i := range homeworks {
		now := time.Now()
		ddl, _ := time.ParseInLocation(timeTmp, homeworks[i].Ddl, time.Local)
		if now.Before(ddl) {
			hmap[homeworks[i]] = ddl.Sub(now).Seconds()
		} else {
			hmap[homeworks[i]] = -1
		}
	}
	fmt.Println(hmap)

	for i := range homeworks {
		for j := 0; j < len(homeworks)-i-1; j++ {
			if hmap[homeworks[j]] == -1 ||
				(hmap[homeworks[j+1]] != -1 && hmap[homeworks[j]] > hmap[homeworks[j+1]]) {
				tmp := homeworks[j]
				homeworks[j] = homeworks[j+1]
				homeworks[j+1] = tmp
			}
		}
	}
}

// AddClockin 发起签到
func (course *Course) AddClockin(password string) bool {
	start := time.Now()
	end := start.Add(time.Minute * 5)

	ddl := end.Format("2006-01-02 15:04:05")
	course.Password = password
	course.Ddl = ddl

	if _, err := O.Update(course, "Password", "Ddl"); err == nil {
		students := course.QueryStudents()
		for i := range students {
			studentcourse := StudentCourses{Student: students[i], Course: course}
			O.Read(&studentcourse, "student_id", "course_id")
			studentcourse.Stat = "未签到"
			O.Update(&studentcourse, "Stat")
		}
		return true
	}
	return false
}

// JudgeClockin 判断是否有签到发起
func (course *Course) JudgeClockin() bool {
	O.Read(course)
	if course.Password == "" && course.Ddl == "" {
		return false
	}
	return true
}

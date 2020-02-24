package models

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
	QueryFiles 查询课件 //未完成
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
		return course.Id, true
	}
	return 0, false
}

// DeleteCourse 删除课堂
func (course *Course) DeleteCourse() bool {
	if _, err := O.Delete(course); err == nil {
		return true
	}
	return false
}

// QueryStudents 查询选择这门课的学生
func (course *Course) QueryStudents() []*Student {
	var students []*Student
	O.QueryTable("student").Filter("Course__Course__id", course.Id).All(&students)
	return students
}

//Addstudent 添加学生
func (course *Course) Addstudent(student *Student) bool {
	ids := []int{student.Id, course.Id}

	if _, err := O.Raw("Insert into student_courses values(null, ?, ?)", ids).Exec(); err == nil {
		return true
	} else {
		return false
	}
}

//QueryFiles 查询课件
func (course *Course) QueryFiles() {
}

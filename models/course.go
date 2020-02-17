package models

// SearchCourse 查找课堂
func SearchCourse(cid int) []*Course {
	var courses []*Course
	err := O.QueryTable("course").Filter("id", cid).One(&courses)
	if err == nil {
		for i, _ := range courses {
			O.Read(courses[i].Teacher)
		}
		return courses
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
	Addstudent 添加学生
	QueryFiles 查询课件
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

//Addstudent 添加学生
func (course *Course) Addstudent(student *Student) bool {
	n, _ := O.QueryTable("course").Filter("id", course.Id).Filter("Student__Id", student.Id).Count()

	ids := []int{student.Id, course.Id}
	if n != 0 {
		return false
	}
	if _, err := O.Raw("Insert into student_courses values(null, ?, ?)", ids).Exec(); err == nil {
		return true
	} else {
		return false
	}
}

//QueryFiles 查询课件
func (course *Course) QueryFiles() {
}

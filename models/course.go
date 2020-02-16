package models

import "github.com/astaxie/beego/orm"

// SearchCourse 查找课堂
func SearchCourse(coursename string) Course {
	res := Course{}
	return res
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
	Id   int
	Name string
	Tid  int
}

// MakeCourse 创建课堂
func (course *Course) MakeCourse() (int, bool) {
	orm := orm.NewOrm()
	orm.Using("default")

	if _, err := orm.Insert(course); err == nil {
		return course.Id, true
	}
	return 0, false
}

// DeleteCourse 删除课堂
func (course *Course) DeleteCourse() bool {
	return true
}

//Addstudent 添加学生
func (course *Course) Addstudent(student *Student) bool {
	return true
}

//QueryFiles 查询课件
func (course *Course) QueryFiles() {

}

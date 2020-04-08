package models

/* Exam 考试
属性说明：
方法说明：
*/
type Exam struct {
	Id     int `orm:"column(id);auto"`
	Course *Course
}

package models

import (
	"strconv"
	"time"
)

//生成通知
func NoticeSBuild(student *Student, course *Course) bool {
	content := student.Name + "请求加入您的课堂，课堂编号为" + strconv.Itoa(course.Id) + "。"
	t := time.Now().Format("2006-01-02 15:04:05")
	notice := &NoticeS{Content: content, Time: t, Student: student, Course: course}

	if _, err := O.Insert(notice); err == nil {
		return true
	}
	return false
}

func NoticeTBuild(teacher *Teacher, course *Course, typ int) bool {
	var content string
	if typ == 1 {
		content = strconv.Itoa(course.Id) + "布置了新作业"
	} else {
		content = strconv.Itoa(course.Id) + "发布了新考试"
	}
	t := time.Now().Format("2006-01-02 15:04:05")
	notice := &NoticeT{Type: typ, Content: content, Time: t, Teacher: teacher, Course: course}

	if _, err := O.Insert(notice); err == nil {
		return true
	}
	return false
}

/* NoticeT 老师对课堂内学生的通知
属性说明：
	Id 顺序产生的编号
	Type 通知类别是1代表考试通知或是1代表作业通知
	Content 通知内容
	Teacher 发起通知的教师
	Course 通知的对象班级
方法说明：
	DeleteNotice 删除通知
*/
type NoticeT struct {
	Id      int `orm:"column(id);auto"`
	Type    int
	Content string
	Time    string
	Teacher *Teacher `orm:"rel(fk)"`
	Course  *Course  `orm:"rel(fk)"`
}

// DeleteNotice 删除通知
func (notice *NoticeT) DeleteNotice() bool {
	if _, err := O.Delete(notice); err == nil {
		return true
	}
	return false
}

/* NoticeS 学生对教师的通知
属性说明：
	Id 顺序产生的编号
	Content 通知内容
	Teacher 发起通知的学生
	Course 通知的对象课堂
方法说明：
	DeleteNotice 删除通知
*/
type NoticeS struct {
	Id      int `orm:"column(id);auto"`
	Content string
	Time    string
	Student *Student `orm:"rel(fk)"`
	Course  *Course  `orm:"rel(fk)"`
}

// DeleteNotice 删除通知
func (notice *NoticeS) DeleteNotice() bool {
	if _, err := O.Delete(notice); err == nil {
		return true
	}
	return false
}

package models

import (
	"os"
	"strconv"
	"time"
)

// 查找考试
func SearchExam(eid int) *Exam {
	var exam Exam
	err := O.QueryTable("exam").Filter("id", eid).One(&exam)
	if err == nil {
		return &exam
	}
	return nil
}

/* Exam 考试
属性说明：
	Id 顺序产生的编号
	ChooseNum 选择题题目数
	ChooseScore 选择题分值
	BlankNum 填空题题目数
	BlankScore 填空题分值
	AnswerNum 简答题题目数
	AnswerScore 简答题分值
	C1~C12 每个章节所占的比重
	StartTime 开始时间
	Time 考试时长
	Course 考试所属班级
方法说明：
	AddExam 发布考试
	JudgeStartData 判断考试是否开始
	JudgeOutData 判断考试是否结束
	QueryStat 查询某个学生的完成状态
*/
type Exam struct {
	Id          int `orm:"column(id);auto"`
	ChooseNum   int
	ChooseScore float64
	BlankNum    int
	BlankScore  float64
	AnswerNum   int
	AnswerScore float64
	C1          float64
	C2          float64
	C3          float64
	C4          float64
	C5          float64
	C6          float64
	C7          float64
	C8          float64
	C9          float64
	C10         float64
	C11         float64
	C12         float64
	StartTime   string
	Time        int
	Course      *Course `orm:"rel(fk)"`
	Stat        string  `orm:"-"`
}

// AddExam 发布考试
func (exam *Exam) AddExam() bool {
	if _, err := O.Insert(exam); err == nil {
		path := "files/exam/" + strconv.Itoa(exam.Id)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return false
		}
		students := exam.Course.QueryStudents()
		for i := range students {
			s_e := &StudentExam{Stat: "未参加", Score: 0, Student: students[i], Exam: exam}
			if _, err := O.Insert(s_e); err != nil {
				return false
			}
			if !BuildPaper(exam, students[i]) {
				return false
			}
			NoticeTBuild(exam.Course.Teacher, exam.Course, students[i], 2)
		}
		return true
	}
	return false
}

// JudgeStartData 判断考试是否开始
func (exam *Exam) JudgeStartData() bool {
	timeTmp := "2006-01-02 15:04"
	starttime, _ := time.ParseInLocation(timeTmp, exam.StartTime, time.Local)

	if time.Now().Before(starttime) {
		return false
	}
	return true
}

// JudgeOutData 判断考试是否结束
func (exam *Exam) JudgeOutData() bool {
	timeTmp := "2006-01-02 15:04"
	starttime, _ := time.ParseInLocation(timeTmp, exam.StartTime, time.Local)

	addtime, _ := time.ParseDuration(strconv.Itoa(exam.Time) + "m")
	endtime := starttime.Add(addtime)

	if time.Now().Before(endtime) {
		return false
	}
	return true
}

// QueryStat 查询某个学生的完成状态
func (exam *Exam) QueryStat(student *Student) {
	var se StudentExam
	O.QueryTable("student_exam").Filter("exam_id", exam.Id).Filter("student_id", student.Id).One(&se)
	exam.Stat = se.Stat
}

// QueryAllScore 查询全部学生的成绩
func (exam *Exam) QueryAllScore() []*StudentExam {
	var ses []*StudentExam
	O.QueryTable("student_exam").Filter("exam_id", exam.Id).OrderBy("-score").All(&ses)
	return ses
}

// QueryScore 查询某个学生的成绩
func (exam *Exam) QueryScore(student *Student) float64 {
	var se StudentExam
	O.QueryTable("student_exam").Filter("exam_id", exam.Id).Filter("student_id", student.Id).One(&se)
	return se.Score
}

/* StudentExam 记录每个学生的考试信息
	属性说明：
	Id 顺序产生的编号
	Stat 是否已完成考试
	Score 学生的分数，默认为0
	Student 学生
	Exam 考试
方法说明：
*/
type StudentExam struct {
	Id      int `orm:"column(id);auto"`
	Stat    string
	Score   float64
	Student *Student `orm:"rel(fk)"`
	Exam    *Exam    `orm:"rel(fk)"`
}

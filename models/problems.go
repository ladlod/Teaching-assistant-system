package models

import (
	"os"
	"strconv"
)

var Chapter []string = []string{
	"1.绪论",
	"2.一个简单的语法制导翻译器",
	"3.词法分析",
	"4.语法分析",
	"5.语法制导的翻译",
	"6.中间代码生成",
	"7.运行时刻环境",
	"8.代码生成",
	"9.机器无关优化",
	"10.指令级并行",
	"11.并行性和局部性优化",
	"12.过程间分析",
}

// BuildPaper 组卷
func BuildPaper(exam *Exam, student *Student) bool {
	var problemNum [12][3]int
	var C [12]float64
	C[0] = exam.C1 / 100
	C[1] = exam.C2 / 100
	C[2] = exam.C3 / 100
	C[3] = exam.C4 / 100
	C[4] = exam.C5 / 100
	C[5] = exam.C6 / 100
	C[6] = exam.C7 / 100
	C[7] = exam.C8 / 100
	C[8] = exam.C9 / 100
	C[9] = exam.C10 / 100
	C[10] = exam.C11 / 100
	C[11] = exam.C12 / 100
	flag := -1
	for i := range C {
		if C[i] != 0 && flag == -1 {
			flag = i
			problemNum[i][0] = exam.ChooseNum
			problemNum[i][1] = exam.BlankNum
			problemNum[i][2] = exam.AnswerNum
		} else {
			problemNum[i][0] = int(float64(exam.ChooseNum) * C[i])
			problemNum[i][1] = int(float64(exam.BlankNum) * C[i])
			problemNum[i][2] = int(float64(exam.AnswerNum) * C[i])
		}
	}
	for i := flag + 1; i < 12; i++ {
		problemNum[flag][0] -= problemNum[i][0]
		problemNum[flag][1] -= problemNum[i][1]
		problemNum[flag][2] -= problemNum[i][2]
	}

	var problems []Problem
	for i := 0; i < 3; i++ {
		for j := 0; j < 12; j++ {
			var tmp []Problem
			O.Raw("select * from problem where chapter = ? and type = ? order by rand() limit ?", j+1, i+1, problemNum[j][i]).QueryRows(&tmp)
			for k := range tmp {
				problems = append(problems, tmp[k])
			}
		}
	}

	if paper, err := os.Create("files/exam/" + strconv.Itoa(exam.Id) + "/" + strconv.Itoa(student.Id)); err == nil {
		defer paper.Close()
		for i := range problems {
			paper.WriteString(strconv.Itoa(problems[i].Id) + "\n")
		}
		return true
	}
	return false
}

// GetProblemById 获得题目
func GetProblemById(pid int) *Problem {
	var problem Problem
	O.QueryTable("problem").Filter("id", pid).One(&problem)
	return &problem
}

// QueryProblems 查询某章节的题目
func QueryProblems(cid int) []*Problem {
	var problems []*Problem
	O.QueryTable("problem").Filter("chapter", cid).All(&problems)

	return problems
}

// AddProblem 添加题目
func AddProblem(probelm Problem, cid int) bool {
	probelm.Chapter = cid
	if _, err := O.Insert(&probelm); err == nil {
		return true
	}
	return false
}

/* Problem 题目
属性说明：
	Id 顺序产生的编号
	Chapter 题目属于的章节
	Type 题目的类型 1-选择题 2-填空题 3-简答题
	Question 题目内容
	Answer 题目参考答案
方法说明：
	Delete 删除题目
*/
type Problem struct {
	Id       int `orm:"column(id);auto"`
	Chapter  int
	Type     int
	A        string `orm:"null"`
	B        string `orm:"null"`
	C        string `orm:"null"`
	D        string `orm:"null"`
	Question string `orm:"size(1024)"`
	Answer   string `orm:"size(1024)"`
}

func (problem *Problem) TableIndex() [][]string {
	return [][]string{
		[]string{"chapter", "type"},
	}
}

// Delete 删除题目
func (problem *Problem) Delete() (int, bool) {
	O.Read(problem)
	if _, err := O.Delete(problem); err == nil {
		return problem.Chapter, true
	}
	return 0, false
}

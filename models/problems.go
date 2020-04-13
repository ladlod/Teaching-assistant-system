package models

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
		[]string{"chapter"},
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

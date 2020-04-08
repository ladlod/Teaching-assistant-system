package models

/* Chapter 章节
属性说明：
	Id 顺序产生的编号
	Tpoic 章节名称
方法说明：
	QueryProblems 查询该章节的题目
	AddProblem 添加题目
*/
type Chapter struct {
	Id    int `orm:"column(id);auto"`
	Topic string
}

// NewChapter 创建新章节
func NewChapter(topic string) int64 {
	Chapter := Chapter{Topic: topic}
	id, _ := O.Insert(Chapter)
	return id
}

// QueryAllChapter 查询全部章节
func QUeryAllChapter() []*Chapter {
	var chapter []*Chapter
	O.QueryTable("chapter").OrderBy("id").All(&chapter)

	return chapter
}

// QueryProblems 查询该章节的题目
func (chapter *Chapter) QueryProblems() []*Problem {
	var problems []*Problem
	O.QueryTable("problem").Filter("chapter_id", chapter.Id).All(&problems)

	return problems
}

// AddProblem 添加题目
func (chapter *Chapter) AddProblem(probelm Problem) bool {
	probelm.Chapter = chapter
	if _, err := O.Insert(probelm); err == nil {
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
	Chapter  *Chapter
	Type     int
	Question string `orm:"size(1024)"`
	Answer   string `orm:"size(1024)"`
}

func (problem *Problem) TableIndex() [][]string {
	return [][]string{
		[]string{"chapter_id"},
	}
}

// Delete 删除题目
func (problem *Problem) Delete() bool {
	if _, err := O.Delete(problem); err == nil {
		return true
	}
	return false
}

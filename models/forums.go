package models

import "time"

// GetQuestionAuthorName 获得发帖人姓名
func GetQuestionAuthorName(questions []*Question) {
	for i := range questions {
		if questions[i].Teacher == nil {
			O.Read(questions[i].Student)
			questions[i].AuthorName = questions[i].Student.Name
		} else {
			O.Read(questions[i].Teacher)
			questions[i].AuthorName = questions[i].Teacher.Name
		}
	}
}

// GetAnswerAuthorName 获得回帖人姓名
func GetAnswerAuthorName(answers []*Answer) {
	for i := range answers {
		if answers[i].Teacher == nil {
			O.Read(answers[i].Student)
			answers[i].AuthorName = answers[i].Student.Name
		} else {
			O.Read(answers[i].Teacher)
			answers[i].AuthorName = answers[i].Teacher.Name
		}
	}
}

// QueryQuestionsBySupport 根据推荐数查询全部帖子
func QueryQuestionsBySupport() []*Question {
	var questions []*Question
	O.QueryTable("question").OrderBy("-support").All(&questions)
	GetQuestionAuthorName(questions)

	return questions
}

// QueryQuestionsByTime 根据发帖时间查询全部帖子
func QueryQuestionsByTime() []*Question {
	var questions []*Question
	O.QueryTable("question").OrderBy("-id").All(&questions)
	GetQuestionAuthorName(questions)

	return questions
}

// AddQuestion 发帖
func AddQuestion(topic string, content string, student *Student, teacher *Teacher) bool {
	time := time.Now().Format("2006-01-02 15:04:05")

	var q *Question
	if teacher == nil {
		q = &Question{Topic: topic, Content: content, Time: time, Student: student}
	} else {
		q = &Question{Topic: topic, Content: content, Time: time, Teacher: teacher}
	}

	if _, err := O.Insert(q); err == nil {
		return true
	}
	return false
}

// GetQuestionById 根据帖子Id获得帖子
func GetQuestionById(id int) *Question {
	var question Question
	err := O.QueryTable("question").Filter("id", id).One(&question)
	if err == nil {
		if question.Teacher == nil {
			O.Read(question.Student)
			question.AuthorName = question.Student.Name
		} else {
			O.Read(question.Teacher)
			question.AuthorName = question.Teacher.Name
		}
		return &question
	}
	return nil
}

// GetAnswerById 根据回子Id获得回子
func GetAnswerById(id int) *Answer {
	var answer Answer
	err := O.QueryTable("answer").Filter("id", id).One(&answer)
	if err == nil {
		return &answer
	}
	return nil
}

/* Question 发帖
属性说明：
	Id 顺序产生的编号
	Topic 帖子标题
	Content 帖子内容
	Time 发帖时间
	Support 帖子推荐数
	Teacher 发帖教师，当发帖人为学生时为null
	Student 发帖学生，当发帖人为教师时为null
	Answers 帖子的回帖
方法说明：
	Delete 删帖
	QueryAnswerByTime 根据回帖时间查询全部回帖
	QueryAnswerBySupport 根据回帖点亮数查询全部回帖
	SupportQuestion 推荐帖子
	AnswerQuestion 回帖
*/
type Question struct {
	Id         int `orm:"column(id);auto"`
	Topic      string
	Content    string `orm:"size(1024)"`
	Time       string
	Support    int       `orm:"default(0)"`
	Student    *Student  `orm:"rel(fk);null"`
	Teacher    *Teacher  `orm:"rel(fk);null"`
	Answers    []*Answer `orm:"reverse(many)"`
	AuthorName string    `orm:"-"`
}

// Delete 删帖
func (question *Question) Delete() bool {
	if _, err := O.Delete(question); err == nil {
		return true
	}
	return false
}

// QueryAnswerByTime 根据回帖时间查询全部回帖
func (question *Question) QueryAnswerByTime() []*Answer {
	var answers []*Answer
	O.QueryTable("answer").Filter("question_id", question.Id).All(&answers)
	GetAnswerAuthorName(answers)

	return answers
}

// QueryAnswerBySupport 根据回帖点亮数查询全部回帖
func (question *Question) QueryAnswerBySupport() []*Answer {
	var answers []*Answer
	O.QueryTable("answer").Filter("question_id", question.Id).OrderBy("-support").All(&answers)
	GetAnswerAuthorName(answers)

	return answers
}

// SupportQuestion 推荐帖子
func (question *Question) SupportQuestion(student *Student, teacher *Teacher) bool {
	if student == nil {
		if O.QueryTable("support_question").Filter("teacher_id", teacher.Id).Filter("question_id", question.Id).Exist() {
			return false
		} else {
			sq := &SupportQuestion{Teacher: teacher, Question: question}
			O.Insert(sq)
		}
	}
	if teacher == nil {
		if O.QueryTable("support_question").Filter("student_id", student.Id).Filter("question_id", question.Id).Exist() {
			return false
		} else {
			sq := &SupportQuestion{Student: student, Question: question}
			O.Insert(sq)
		}
	}
	question.Support++
	if _, err := O.Update(question); err == nil {
		return true
	}
	return false
}

// AnswerQuestion 回帖
func (question *Question) AnswerQuestion(content string, teacher *Teacher, student *Student) bool {
	time := time.Now().Format("2006-01-02 15:04:05")
	var an *Answer
	if teacher == nil {
		an = &Answer{Content: content, Time: time, Student: student, Question: question}
	} else {
		an = &Answer{Content: content, Time: time, Teacher: teacher, Question: question}
	}

	if _, err := O.Insert(an); err == nil {
		AddAnswerNotice(student, teacher, question)
		return true
	}
	return false
}

/* Answer 回帖
属性说明：
	Id 顺序产生的编号
	Content 回帖内容
	Time 回帖时间
	Support 回帖点赞数
	Teacher 回帖教师，当回帖人为学生时为null
	Student 回帖学生，当回帖人为教师时为null
	Question 回复的主贴
方法说明：
	SupportAnswer点亮回帖
*/
type Answer struct {
	Id         int    `orm:"column(id);auto"`
	Content    string `orm:"size(512)"`
	Time       string
	Support    int       `orm:"default(0)"`
	Student    *Student  `orm:"rel(fk);null"`
	Teacher    *Teacher  `orm:"rel(fk);null"`
	Question   *Question `orm:"rel(fk)"`
	AuthorName string    `orm:"-"`
}

func (answer *Answer) TableIndex() [][]string {
	return [][]string{
		[]string{"question_id"},
	}
}

// SupportAnswer 点亮回帖
func (answer *Answer) SupportAnswer(student *Student, teacher *Teacher) bool {
	if student == nil {
		if O.QueryTable("support_answer").Filter("teacher_id", teacher.Id).Filter("answer_id", answer.Id).Exist() {
			return false
		} else {
			sa := &SupportAnswer{Teacher: teacher, Answer: answer}
			O.Insert(sa)
		}
	}
	if teacher == nil {
		if O.QueryTable("support_answer").Filter("student_id", student.Id).Filter("answer_id", answer.Id).Exist() {
			return false
		} else {
			sa := &SupportAnswer{Student: student, Answer: answer}
			O.Insert(sa)
		}
	}
	answer.Support++
	if _, err := O.Update(answer); err == nil {
		return true
	}
	return false
}

type SupportQuestion struct {
	Id       int       `orm:"auto"`
	Student  *Student  `orm:"rel(fk);null"`
	Teacher  *Teacher  `orm:"rel(fk);null"`
	Question *Question `orm:"rel(fk)"`
}

func (sq *SupportQuestion) TableIndex() [][]string {
	return [][]string{
		[]string{"student_id", "teacher_id", "question_id"},
	}
}

type SupportAnswer struct {
	Id      int      `orm:"auto"`
	Student *Student `orm:"rel(fk);null"`
	Teacher *Teacher `orm:"rel(fk);null"`
	Answer  *Answer  `orm:"rel(fk)"`
}

func (sa *SupportAnswer) TableIndex() [][]string {
	return [][]string{
		[]string{"student_id", "teacher_id", "answer_id"},
	}
}

type AnswerNotice struct {
	Id            int `orm:"auto"`
	Content       string
	SourceStudent *Student  `orm:"rel(fk);null"`
	SourceTeacher *Teacher  `orm:"rel(fk);null"`
	Question      *Question `orm:"rel(fk)"`
}

// 构造回帖时的通知
func AddAnswerNotice(s *Student, t *Teacher, q *Question) bool {
	var answernotice *AnswerNotice
	if s == nil {
		content := t.Name + "回复了您的帖子。"
		answernotice = &AnswerNotice{Content: content, SourceTeacher: t, Question: q}
	} else {
		content := s.Name + "回复了您的帖子。"
		answernotice = &AnswerNotice{Content: content, SourceStudent: s, Question: q}
	}
	if _, err := O.Insert(answernotice); err == nil {
		return true
	}
	return false
}

// DeleteNotice 删除通知
func (notice *AnswerNotice) Delete() bool {
	if _, err := O.Delete(notice); err == nil {
		return true
	}
	return false
}

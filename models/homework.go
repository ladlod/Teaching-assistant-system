package models

import (
	"os"
	"strconv"
)

type FileHomework struct {
	Id      int `orm:"column(id);auto"`
	Content string
	Ddl     string
	Course  *Course `orm:"rel(fk)"`
}

func (h *FileHomework) AddHomework(content string, ddl string, course *Course) bool {
	homework := FileHomework{Content: content, Ddl: ddl, Course: course}
	if _, err := O.Insert(homework); err == nil {
		path := "files/homework/" + strconv.Itoa(homework.Id)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return false
		}
		return true
	}
	return false
}

type QuestionHomework struct {
	Id      int `orm:"column(id);auto"`
	Chapter int
	ChoiceQ int8
	BlankQ  int8
	AnswerQ int8
	Course  *Course `orm:"rel(fk)"`
}

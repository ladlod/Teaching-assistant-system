package models

import (
	"os"
	"strconv"
)

type Homework struct {
	Id      int `orm:"column(id);auto"`
	Content string
	Ddl     string
	Course  *Course `orm:"rel(fk)"`
}

func (h *Homework) AddHomework() bool {
	if _, err := O.Insert(h); err == nil {
		path := "files/homework/" + strconv.Itoa(h.Id)
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return false
		}
		return true
	}
	return false
}

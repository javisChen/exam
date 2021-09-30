package db

import (
	"exam/models"
	"fmt"
	"testing"
)

func TestSelectOne(t *testing.T) {
	user := models.User{}
	_ = SelectOne("select * from user where id = ?", &user, 1)
	fmt.Println(user)
}

func TestSelectList(t *testing.T) {
	var users []models.User
	_, _ = SelectList("select * from user", &users)
	fmt.Println(users)
}

func TestExec(t *testing.T) {
	question := models.Question{
		PaperId: 1,
		Title:   "1",
		Type:    0,
		Score:   0,
		Answer:  "1",
	}
	rows, _ := Exec("insert into question(paper_id, title, type, score, answer) values(?, ?, ?, ?, ?)", question.PaperId, question.Title, question.Type, question.Score, question.Answer)
	fmt.Println(rows)
}

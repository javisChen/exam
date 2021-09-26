package question

import (
	"exam/core/db"
	"exam/models"
	"github.com/beego/beego/v2/client/orm"
)

func SelectListByPaperId(paperId int64) ([]models.Question, error) {
	sql := "select * from question where paper_id = ?"
	var questions []models.Question
	_, err := db.SelectList(sql, &questions, paperId)
	return questions, err
}

func Insert(question models.Question, txOrm orm.TxOrmer) (int64, error) {
	sql := "insert into question(paper_id, title, type , score, answer) values(?, ?, ?, ?, ?)"
	result, err := db.ExecTx(txOrm, sql, question.PaperId, question.Title, question.Type, question.Score, question.Answer)
	if err != nil {
		return 0, err
	}
	question.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

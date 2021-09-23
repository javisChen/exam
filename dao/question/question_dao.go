package question

import (
	"exam/core"
	"exam/models"
	"github.com/beego/beego/v2/client/orm"
)

func SelectListByPaperId(paperId int64) []models.Question {
	var ormer = core.GetOrm()
	var questions []models.Question
	_, _ = ormer.Raw("select * from question where paper_id = ?", paperId).QueryRows(&questions)
	return questions
}

func Insert(question models.Question, txOrm orm.TxOrmer) (int64, error) {
	sql := "insert into question(paper_id, title, type , score, answer) values(?, ?, ?, ?, ?)"
	result, err := txOrm.Raw(sql, question.PaperId, question.Title, question.Type, question.Score, question.Answer).Exec()
	if err != nil {
		return 0, err
	}
	question.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

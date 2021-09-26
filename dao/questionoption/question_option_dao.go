package questionoption

import (
	"exam/core/db"
	"exam/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

func SelectListByQuestionIds(questionIds []int64) ([]models.QuestionOption, error) {
	sql := "select * from question_option where question_id in (%s)"
	var idsStr string
	for range questionIds {
		idsStr += "?,"
	}
	idsStr = strings.TrimRight(idsStr, ",")
	var questionOptions []models.QuestionOption
	_, err := db.SelectList(fmt.Sprintf(sql, idsStr), &questionOptions, questionIds)
	return questionOptions, err
}

func Insert(option *models.QuestionOption, txOrm orm.TxOrmer) (int64, error) {
	sql := "insert into question_option(question_id, title, seq) values(?, ?, ?)"
	result, err := db.ExecTx(txOrm, sql, option.QuestionId, option.Title, option.Seq)
	if err != nil {
		return 0, err
	}
	option.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

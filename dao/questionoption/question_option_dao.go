package questionoption

import (
	"exam/core"
	"exam/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

func SelectListByQuestionIds(questionIds []int64) []models.QuestionOption {
	var ormer = core.GetOrm()
	var idsStr string
	for range questionIds {
		idsStr += "?,"
	}
	idsStr = strings.TrimRight(idsStr, ",")
	var questionOptions []models.QuestionOption
	_, _ = ormer.Raw(fmt.Sprintf("select * from question_option where question_id in (%s)", idsStr), questionIds).QueryRows(&questionOptions)
	return questionOptions
}

func Insert(option *models.QuestionOption, txOrm orm.TxOrmer) (int64, error) {
	sql := "insert into question_option(question_id, title, seq) values(?, ?, ?)"
	result, err := txOrm.Raw(sql, option.QuestionId, option.Title, option.Seq).Exec()
	if err != nil {
		return 0, err
	}
	option.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

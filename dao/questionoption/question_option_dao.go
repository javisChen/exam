package questionoption

import (
	"exam/core"
	"exam/models"
	"fmt"
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

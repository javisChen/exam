package question

import (
	"exam/core"
	"exam/models"
)

func SelectListByPaperId(paperId int64) []models.Question {
	var ormer = core.GetOrm()
	var questions []models.Question
	_, _ = ormer.Raw("select * from question where paper_id = ?", paperId).QueryRows(&questions)
	return questions
}

package paper

import (
	"exam/core"
	"exam/models"
)

func SelectById(id int64) models.Paper {
	var ormer = core.GetOrm()
	paper := models.Paper{}
	_ = ormer.Raw("select * from paper where id = ?", id).QueryRow(&paper)
	return paper
}

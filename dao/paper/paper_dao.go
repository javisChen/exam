package paper

import (
	"exam/core"
	"exam/models"
	"github.com/beego/beego/v2/client/orm"
)

func SelectById(id int64) models.Paper {
	var ormer = core.GetOrm()
	paper := models.Paper{}
	_ = ormer.Raw("select * from paper where id = ?", id).QueryRow(&paper)
	return paper
}

func SelectList() []models.Paper {
	var papers []models.Paper
	core.GetOrm().Raw("select * from paper order by gmt_created desc").QueryRows(&papers)
	return papers
}

func Insert(paper *models.Paper, txOrm orm.TxOrmer) (int64, error) {
	sql := "insert into paper(title) values(?)"
	result, err := txOrm.Raw(sql, paper.Title).Exec()
	if err != nil {
		return 0, err
	}
	paper.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

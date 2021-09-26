package paper

import (
	"exam/core/db"
	"exam/models"
	"github.com/beego/beego/v2/client/orm"
)

func SelectById(id int64) (models.Paper, error) {
	sql := "select * from paper where id = ?"
	paper := models.Paper{}
	err := db.SelectOne(sql, &paper, id)
	return paper, err
}

func SelectList() ([]models.Paper, error) {
	sql := "select * from paper order by gmt_created desc"
	var papers []models.Paper
	_, err := db.SelectList(sql, &papers)
	return papers, err
}

func Insert(paper *models.Paper, txOrm orm.TxOrmer) (int64, error) {
	sql := "insert into paper(title) values(?)"
	result, err := db.ExecTx(txOrm, sql, paper.Title)
	if err != nil {
		return 0, err
	}
	paper.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

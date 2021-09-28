package userpaperrecord

import (
	"exam/core/db"
	"exam/models"
)

func SelectByUserIdAndPaperId(paperId int64, userId int64) (models.UserPagerRecord, error) {
	sql := "select * from user_pager_record where paper_id = ? and user_id = ?"
	record := models.UserPagerRecord{}
	err := db.SelectOne(sql, &record, paperId, userId)
	return record, err
}

func Insert(record models.UserPagerRecord) (int64, error) {
	sql := "insert into user_pager_record(paper_id, user_id) values(?, ?)"
	result, err := db.Exec(sql, record.PaperId, record.UserId)
	if err != nil {
		return 0, err
	}
	record.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

func UpdateStatus(id int64, status int) (int64, error) {
	sql := "update user_pager_record set status = ? where id = ?"
	result, err := db.Exec(sql, status, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

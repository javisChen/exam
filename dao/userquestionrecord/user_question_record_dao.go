package userquestionrecord

import (
	"exam/core/db"
	"exam/models"
)

func SelectByQuestionIdAndUserId(record models.UserQuestionRecord) (models.UserQuestionRecord, error) {
	sql := "select * from user_question_record where question_id = ? and user_id = ?"
	var records models.UserQuestionRecord
	err := db.SelectOne(sql, &records, record.QuestionId, record.UserId)
	if err != nil {
		return records, err
	}
	return records, nil
}

func Insert(record models.UserQuestionRecord) (int64, error) {
	sql := "insert into user_question_record(user_id, paper_id, question_id, user_answer) values(?, ?, ?, ?)"
	result, err := db.Exec(sql, record.UserId, record.PaperId, record.QuestionId, record.UserAnswer)
	if err != nil {
		return 0, err
	}
	record.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

func UpdateUserAnswer(record models.UserQuestionRecord) (int64, error) {
	sql := "update user_question_record set user_answer = ? where question_id = ?"
	result, err := db.Exec(sql, record.UserAnswer, record.QuestionId)
	if err != nil {
		return 0, err
	}
	record.Id, _ = result.LastInsertId()
	return result.RowsAffected()
}

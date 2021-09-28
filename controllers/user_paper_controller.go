package controllers

import (
	"encoding/json"
	"exam/core"
	userPaperRecordDao "exam/dao/userpaperrecord"
	userQuestionRecordDao "exam/dao/userquestionrecord"
	"exam/models"
)

type UserPaperController struct {
	core.BaseController
}

func (c UserPaperController) Create() {
	jsonParam := c.GetJsonParam()
	loginUser := c.GetLoginUser()

	val, _ := jsonParam["paperId"].(json.Number).Int64()
	result, _ := userPaperRecordDao.SelectByUserIdAndPaperId(val, loginUser.Id)
	if result == (models.UserPagerRecord{}) {
		_, _ = userPaperRecordDao.Insert(models.UserPagerRecord{
			BaseModel: models.BaseModel{},
			UserId:    loginUser.Id,
			PaperId:   val,
		})
	}
	c.Success()
}

func (c UserPaperController) Finish() {
	jsonParam := c.GetJsonParam()
	loginUser := c.GetLoginUser()

	val, _ := jsonParam["paperId"].(json.Number).Int64()
	result, _ := userPaperRecordDao.SelectByUserIdAndPaperId(val, loginUser.Id)
	_, _ = userPaperRecordDao.UpdateStatus(result.Id, 1)
	c.Success()
}

func (c UserPaperController) Answer() {
	jsonParam := c.GetJsonParam()
	loginUser := c.GetLoginUser()

	questionId, _ := jsonParam["questionId"].(json.Number).Int64()
	paperId, _ := jsonParam["paperId"].(json.Number).Int64()
	userAnswer, _ := jsonParam["userAnswer"].(string)
	userId := loginUser.Id
	result, _ := userQuestionRecordDao.SelectByQuestionIdAndUserId(models.UserQuestionRecord{
		QuestionId: questionId,
		UserId:     userId,
	})
	if (result == models.UserQuestionRecord{}) {
		userQuestionRecordDao.Insert(models.UserQuestionRecord{
			QuestionId: questionId,
			UserId:     userId,
			UserAnswer: userAnswer,
			PaperId:    paperId,
		})
	} else {
		userQuestionRecordDao.UpdateUserAnswer(models.UserQuestionRecord{
			BaseModel:  models.BaseModel{Id: result.Id},
			QuestionId: questionId,
			UserId:     userId,
			UserAnswer: userAnswer,
		})
	}
	c.Success()
}

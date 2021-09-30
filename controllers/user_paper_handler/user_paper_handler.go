package user_paper_handler

import (
	"exam/core/web"
	userPaperRecordDao "exam/dao/userpaperrecord"
	userQuestionRecordDao "exam/dao/userquestionrecord"
	"exam/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	jsonParam, err := web.GetJsonParam(c)
	if err != nil {
		web.Error(c)
	}
	loginUser := web.GetLoginUser(c)

	val := int64((*jsonParam)["paperId"].(float64))
	result, _ := userPaperRecordDao.SelectByUserIdAndPaperId(val, loginUser.Id)
	if result == (models.UserPagerRecord{}) {
		_, _ = userPaperRecordDao.Insert(models.UserPagerRecord{
			BaseModel: models.BaseModel{},
			UserId:    loginUser.Id,
			PaperId:   val,
		})
	}
	web.Ok(c)
}

func Finish(c *gin.Context) {
	jsonParam, err := web.GetJsonParam(c)
	if err != nil {
		web.Error(c)
	}
	loginUser := web.GetLoginUser(c)

	val := int64((*jsonParam)["paperId"].(float64))
	result, _ := userPaperRecordDao.SelectByUserIdAndPaperId(val, loginUser.Id)
	_, _ = userPaperRecordDao.UpdateStatus(result.Id, 1)
	web.Ok(c)
}

func Answer(c *gin.Context) {
	jsonParam, err := web.GetJsonParam(c)
	if err != nil {
		web.Error(c)
	}
	loginUser := web.GetLoginUser(c)
	questionId := int64((*jsonParam)["questionId"].(float64))
	paperId := int64((*jsonParam)["paperId"].(float64))
	userAnswer, _ := (*jsonParam)["userAnswer"].(string)
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
	web.Ok(c)
}

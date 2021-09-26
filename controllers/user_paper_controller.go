package controllers

import (
	"encoding/json"
	"exam/core"
	userPaperRecordDao "exam/dao/userpaperrecord"
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

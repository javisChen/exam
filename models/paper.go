package models

import "github.com/beego/beego/v2/client/orm"

type Paper struct {
	BaseModel
	Title string `json:"title" gorm:"column:title"` // 试卷标题
}

func (m *Paper) TableName() string {
	return "paper"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Paper))
}

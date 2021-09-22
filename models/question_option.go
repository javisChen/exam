package models

import "github.com/beego/beego/v2/client/orm"

type QuestionOption struct {
	BaseModel
	QuestionId uint64 `json:"question_id" orm:"column(question_id)"` // 试题id,关联question.id
	Title      string `json:"title" orm:"column(title)"`             // 选项标题
	Seq        int    `json:"seq" orm:"column(seq)"`                 // 序列
}

func (m *QuestionOption) TableName() string {
	return "question_option"
}

func init() {
	orm.RegisterModel(new(QuestionOption))
}

package models

type QuestionOption struct {
	BaseModel
	QuestionID uint64 `json:"question_id" gorm:"column:question_id"` // 试题id,关联question.id
	Title      string `json:"title" gorm:"column:title"`             // 选项标题
	Code       int    `json:"code" gorm:"column:code"`               // 选项编码
}

func (m *QuestionOption) TableName() string {
	return "question_option"
}
